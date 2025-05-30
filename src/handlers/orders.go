package handlers

import (
	"encoding/json"
	"level-scale/logger"
	"net/http"

	"level-scale/db"
	"level-scale/middleware"
	"level-scale/models"
)

type CheckoutRequest struct {
	ShippingAddress string `json:"ShippingAddress"`
}

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var cart models.Cart
	if err := db.Db.First(&cart, "user_id = ?", userID).Error; err != nil {
		http.Error(w, "no active cart", http.StatusBadRequest)
		return
	}

	var cartItems []models.CartItem
	if err := db.Db.Where("cart_id = ?", cart.ID).Find(&cartItems).Error; err != nil {
		http.Error(w, "could not load cart items", http.StatusInternalServerError)
		return
	}
	if len(cartItems) == 0 {
		http.Error(w, "cart is empty", http.StatusBadRequest)
		return
	}

	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	order := models.Order{
		UserID:          userID,
		CartID:          cart.ID,
		Status:          "pending",
		ShippingAddress: req.ShippingAddress,
	}

	if err := db.Db.Create(&order).Error; err != nil {
		http.Error(w, "could not create order", http.StatusInternalServerError)
		return
	}

	var total float32
	for _, item := range cartItems {
		var product models.Product
		if err := db.Db.First(&product, item.ProductID).Error; err != nil {
			http.Error(w, "product not found", http.StatusBadRequest)
			return
		}

		if uint16(item.Quantity) > product.Stock {
			http.Error(w, "insufficient stock", http.StatusBadRequest)
			return
		}

		product.Stock -= uint16(item.Quantity)
		db.Db.Save(&product)

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
		db.Db.Create(&orderItem)

		total += float32(item.Quantity) * product.Price
	}

	invoice := models.Invoice{
		OrderID:   order.ID,
		Amount:    total,
		CreatedAt: order.CreatedAt,
	}
	db.Db.Create(&invoice)

	delivery := models.Delivery{
		OrderID:    order.ID,
		ExpectedAt: order.CreatedAt.AddDate(0, 0, 3),
		Status:     "scheduled",
	}

	db.Db.Create(&delivery)
	db.Db.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"orderId": order.ID})
	if err != nil {
		logger.Log.Errorw("Response encoding failed", "err", err)
		http.Error(w, "Order placed created but response failed", http.StatusInternalServerError)
		return
	}
}
