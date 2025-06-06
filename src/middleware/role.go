package middleware

import (
	"net/http"

	"level-scale/dbmanager"
	"level-scale/models"
)

func RequireSeller(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := GetUserID(r)

		var user models.User
		if err := dbmanager.Db.First(&user, userID).Error; err != nil || !user.IsSeller {
			http.Error(w, "seller access only", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
