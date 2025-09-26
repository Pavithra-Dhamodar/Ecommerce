package dao

import (
	"context"
	"ecommercefinal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao struct {
	Collection *mongo.Collection
}

func NewUserDao(db *mongo.Database) *UserDao{
	return &UserDao{Collection: db.Collection("users")}

}

func(u *UserDao) CreateUser(ctx context.Context, user *models.User) error{
   _, err:=u.Collection.InsertOne(ctx, user)
   if(err!=nil){
	return err
   }
   return nil
}

