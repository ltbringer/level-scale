package routes

import (
	"net/http"

	"level-scale/handlers"
	"level-scale/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Get("/products", handlers.GetAllProductsHandler)

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Group(func(seller chi.Router) {
			seller.Use(middleware.RequireSeller)
			seller.Post("/products", handlers.CreateProductHandler)
		})
		protected.Post("/cart/items", handlers.AddToCartHandler)
		protected.Post("/checkout", handlers.CheckoutHandler)
		protected.Post("/returns", handlers.ReturnItemsHandler)
	})

	return r
}
