package dao

import (
	"context"
	"ecommercefinal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductDao struct {
	Collection *mongo.Collection
}

func NewProductDao(db *mongo.Database) *ProductDao {
	return &ProductDao{Collection: db.Collection("products")}
}

func (dao *ProductDao) Create(ctx context.Context, product *models.Product) error {
	_, err := dao.Collection.InsertOne(ctx, product)
	return err
}

func (dao *ProductDao) GetAll(ctx context.Context) ([]models.Product, error) {
	cursor, err := dao.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, err
}
