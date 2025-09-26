package services

import (
	"context"
	"ecommercefinal/dao"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartService struct {
	cartDAO    *dao.CartDAO
	productDAO *dao.ProductDao
}

func NewCartService(cartDAO *dao.CartDAO, productDAO *dao.ProductDao) *CartService {
	return &CartService{cartDAO: cartDAO, productDAO: productDAO}
}

func (s *CartService) AddToCart(ctx context.Context, userID, productID primitive.ObjectID) error {
	// Validate product exists
	products, _ := s.productDAO.GetAll(ctx)
	exists := false
	for _, p := range products {
		if p.ProductId == productID {
			exists = true
			break
		}
	}
	if !exists {
		return nil // silently ignore if product not found
	}

	// Just delegate to DAO
	return s.cartDAO.AddToCart(ctx, userID, productID)
}

func (s *CartService) GetCart(ctx context.Context, userID primitive.ObjectID) (*[]primitive.ObjectID, error) {
	cart, err := s.cartDAO.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	if cart == nil {
		empty := []primitive.ObjectID{}
		return &empty, nil
	}
	return &cart.Products, nil
}

func (s *CartService) ClearCart(ctx context.Context, userID primitive.ObjectID) error {
	return s.cartDAO.Clear(ctx, userID)
}
