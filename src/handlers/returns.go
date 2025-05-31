package handlers

import (
	"encoding/json"
	"level-scale/logger"
	"net/http"
	"time"

	"level-scale/dbmanager"
	"level-scale/middleware"
	"level-scale/models"
)

type ReturnRequest struct {
	OrderItemID  uint   `json:"orderItemId"`
	ReturnReason string `json:"reason"`
}

func ReturnItemsHandler(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	var reqs []ReturnRequest
	if err := json.NewDecoder(r.Body).Decode(&reqs); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	for _, req := range reqs {
		var item models.OrderItem
		if err := dbmanager.Db.First(&item, "id = ?", req.OrderItemID).Error; err != nil {
			http.Error(w, "order item not found", http.StatusBadRequest)
			return
		}

		var order models.Order
		if err := dbmanager.Db.First(&order, "id = ?", item.OrderID).Error; err != nil || order.UserID != userID {
			http.Error(w, "unauthorized return attempt", http.StatusUnauthorized)
			return
		}

		ret := models.Return{
			OrderItemID:  req.OrderItemID,
			ReturnReason: req.ReturnReason,
			Status:       "requested",
			RejectReason: "",
			CreatedAt:    time.Now(), // optional since GORM can auto-create
		}

		if err := dbmanager.Db.Create(&ret).Error; err != nil {
			http.Error(w, "failed to create return", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]string{"status": "return requested"})
	if err != nil {
		logger.Log.Errorw("Response encoding failed", "err", err)
		http.Error(w, "Return successful but response failed", http.StatusInternalServerError)
		return
	}
}
