package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"level-scale/dbmanager"
	"level-scale/middleware"
	"level-scale/models"
)

type AddToCartRequest struct {
	ProductID uint64 `json:"productId"`
	Quantity  uint8  `json:"quantity"`
}

func UpsertCart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var cart models.Cart
	if err := dbmanager.Db.FirstOrCreate(&cart, models.Cart{UserID: userID}).Error; err != nil {
		http.Error(w, "could not get or create cart", http.StatusInternalServerError)
		return
	}

	var items []AddToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	var errors []error

	for _, req := range items {
		var item models.CartItem
		err := dbmanager.Db.
			Where("cart_id = ? AND product_id = ?", cart.ID, req.ProductID).
			First(&item).Error

		if err == nil {
			item.Quantity += req.Quantity
			if err := dbmanager.Db.Save(&item).Error; err != nil {
				errors = append(errors, fmt.Errorf("could not add to cart: %s", err.Error()))
			}
		} else {
			newItem := models.CartItem{
				CartID:    cart.ID,
				ProductID: req.ProductID,
				Quantity:  req.Quantity,
			}
			if err := dbmanager.Db.Create(&newItem).Error; err != nil {
				errors = append(errors, fmt.Errorf("failed to update cart item: %w", err))
			}
		}
	}
	if len(errors) > 0 {
		http.Error(w, errors[0].Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
