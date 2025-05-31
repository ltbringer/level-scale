package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"level-scale/dbmanager"
	"level-scale/logger"
	"level-scale/middleware"
	"level-scale/models"
)

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       uint16  `json:"stock"`
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var shop models.Shop
	if err := dbmanager.Db.First(&shop, "seller_id = ?", userID).Error; err != nil {
		http.Error(w, "no shop found for seller", http.StatusBadRequest)
		return
	}

	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	product := models.Product{
		ShopID:      shop.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := dbmanager.Db.Create(&product).Error; err != nil {
		logger.Log.Errorw("product creation failed", "err", err)
		http.Error(w, "failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"productId": product.ID})
	if err != nil {
		logger.Log.Errorw("product creation failed", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := dbmanager.Db.Model(&models.Product{})

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 || limit > 100 {
		limit = 10
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	category := r.URL.Query().Get("category")
	if category != "" {
		query = query.Where("category = ?", category)
	}

	subCategory := r.URL.Query().Get("subCategory")
	if subCategory != "" {
		query = query.Where("sub_category = ?", subCategory)
	}

	var products []models.Product
	if err := query.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		http.Error(w, "failed to fetch products", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		logger.Log.Errorw("failed to encode product list", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
