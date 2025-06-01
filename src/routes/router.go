package routes

import (
	"net/http"

	"level-scale/handlers"
	"level-scale/middleware"

	"github.com/go-chi/chi/v5"
)

func Init() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.MetricsMiddleware)
	r.Get("/health", handlers.HealthCheck)
	r.Post("/register", handlers.RegisterUser)
	r.Post("/login", handlers.AuthenticateUser)
	r.Get("/products", handlers.GetProducts)

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Group(func(seller chi.Router) {
			seller.Use(middleware.RequireSeller)
			seller.Post("/products", handlers.CreateProducts)
		})
		protected.Post("/cart/items", handlers.UpsertCart)
		protected.Post("/checkout", handlers.CreateOrder)
		protected.Post("/returns", handlers.UndoOrder)
	})

	return r
}
