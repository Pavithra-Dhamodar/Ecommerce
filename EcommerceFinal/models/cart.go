package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID   `bson:"user_id" json:"user_id"`
	Products []primitive.ObjectID `bson:"products" json:"products"`
}
