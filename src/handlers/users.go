package handlers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"level-scale/dbmanager"
	"level-scale/logger"
	"level-scale/models"
	"level-scale/settings"
	"net/http"
	"strings"
	"time"
)

type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsSeller bool   `json:"isSeller"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "password hash failed", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email:        strings.ToLower(req.Email),
		PasswordHash: string(hashed),
		IsSeller:     req.IsSeller,
	}

	if err := dbmanager.Db.Create(&user).Error; err != nil {
		logger.Log.Warnw("register failed", "err", err)
		http.Error(w, "email already exists or DB error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.Log.Debugw("register success", user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		logger.Log.Errorw("Response encoding failed", "err", err)
		http.Error(w, "User registered but response failed.", http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
	}
	var user models.User
	if err := dbmanager.Db.First(&user, "email = ?", strings.ToLower(req.Email)).Error; err != nil {
		logger.Log.Errorw("Invalid User", "email", req.Email, "err", err)
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	signed, err := token.SignedString(settings.JWTSecret)
	if err != nil {
		logger.Log.Errorw("JWT signing failed", "err", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"token": signed})
	if err != nil {
		logger.Log.Errorw("Response encoding failed", "err", err)
		http.Error(w, "Token created but response failed", http.StatusInternalServerError)
		return
	}
}
