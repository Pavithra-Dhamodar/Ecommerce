package main

import (
	"ecommercefinal/dao"
	"ecommercefinal/db"
	"ecommercefinal/handler"
	"ecommercefinal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	client := db.Connect()
	database := client.Database("ecommerce")

	// DAOs
	cartDao := dao.NewCartDAO(database)
	productDao := dao.NewProductDao(database)
	userDao := dao.NewUserDao(database) // ✅ Correct

	// Services
	cartService := services.NewCartService(cartDao, productDao) // if your CartService needs all 3
	productService := services.NewProductService(productDao)
	userService := services.NewUserService(userDao)

	// Handlers
	cartHandler := handler.NewCartHandler(cartService)
	productHandler := handler.NewProductHandler(productService)
	userHandler := handler.NewUserHandler(userService)

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/cart/add", cartHandler.AddToCart).Methods("POST")
	r.HandleFunc("/cart", cartHandler.GetCart).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
