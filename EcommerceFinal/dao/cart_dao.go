package dao

import (
	"context"
	"ecommercefinal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartDAO struct {
	Collection *mongo.Collection
}

func NewCartDAO(db *mongo.Database) *CartDAO {
	return &CartDAO{Collection: db.Collection("carts")}
}

func (c *CartDAO) AddToCart(ctx context.Context, userID, productID primitive.ObjectID) error {
	_, err := c.Collection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{"$push": bson.M{"products": productID}},
		options.Update().SetUpsert(true), // ✅ correct
	)
	return err
}

func (c *CartDAO) GetCart(ctx context.Context, userID primitive.ObjectID) (*models.Cart, error) {
	var cart models.Cart
	err := c.Collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &cart, nil
}

func (c *CartDAO) Clear(ctx context.Context, userID primitive.ObjectID) error {
	_, err := c.Collection.DeleteOne(ctx, bson.M{"user_id": userID})
	return err
}
