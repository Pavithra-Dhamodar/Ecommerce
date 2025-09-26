package services

import (
	"context"
	"ecommercefinal/dao"
	"ecommercefinal/models"
)

type UserService struct {
	User *dao.UserDao
}

func NewUserService(dao *dao.UserDao) *UserService {
	return &UserService{User: dao}
}

func (u *UserService) CreateUser(ctx context.Context, user *models.User) error {
	err := u.User.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
