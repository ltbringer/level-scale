package handlers

import (
	"encoding/json"
	"level-scale/logger"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "healthy"})
	if err != nil {
		logger.Log.Errorw("Response encoding failed", "err", err)
		http.Error(w, "Healthy service but response failed", http.StatusInternalServerError)
		return
	}
}
