package handler

import (
	"context"
	"ecommercefinal/models"
	"ecommercefinal/services"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	dao *services.ProductService
}

func NewProductHandler(dao *services.ProductService) *ProductHandler {
	return &ProductHandler{dao: dao}

}

func (pro *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)
	pro.dao.CreateProduct(context.Background(), &p)
	json.NewEncoder(w).Encode(p)
}

func (pro *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, _ := pro.dao.GetAll(context.Background())
	json.NewEncoder(w).Encode(products)
}
