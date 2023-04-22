package controllers

import (
	"context"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type UserStorage interface {
	insertUserInDb(us *models.User)
	listUserInDb() []models.User
}

type UserService struct {
	db  *gorm.DB
	ctx context.Context
}

func (u *UserService) listUserInDb() []models.User {

	var users = make([]models.User, 0, 4)

	tx := u.db.WithContext(u.ctx)

	tx.Find(&users)

	return users
}

func (u *UserService) insertUserInDb(us *models.User) {
	tx := u.db.WithContext(u.ctx)

	tx.Create(us)
}

func NewUserStorage(ctx context.Context, db *gorm.DB) UserStorage {
	return &UserService{db: db, ctx: ctx}
}
