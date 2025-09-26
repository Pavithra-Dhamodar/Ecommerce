package services

import (
	"context"
	"ecommercefinal/dao"
	"ecommercefinal/models"
)

type ProductService struct {
	S *dao.ProductDao
}

func NewProductService(pro *dao.ProductDao) *ProductService {
	return &ProductService{S: pro}
}

func (p *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	err := p.S.Create(ctx, product)
	return err
}

func (p *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	return p.S.GetAll(ctx)
}
