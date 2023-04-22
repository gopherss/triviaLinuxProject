package controllers

import (
	"context"
	"net/http"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type UserGateway interface {
	CreateUser(ul *models.User) (int, models.Response)
	ListUser() (int, []models.User)
}

type UserLogic struct {
	Su UserStorage
}

func (ul *UserLogic) ListUser() (int, []models.User) {

	return http.StatusOK, ul.Su.listUserInDb()
}

func (ul *UserLogic) CreateUser(us *models.User) (int, models.Response) {

	length := int(0)

	if len(us.Name) < length || len(us.LastName) < length ||
		len(us.Email) < length || len(us.Password) < length {
		return http.StatusOK, models.Response{
			Message: "User is invalid",
			Success: false,
		}
	}

	go ul.Su.insertUserInDb(us)

	return http.StatusAccepted, models.Response{
		Message: "User Successfully added.",
		Success: true,
	}
}

func NewUserGateway(ctx context.Context, db *gorm.DB) UserGateway {
	return &UserLogic{NewUserStorage(ctx, db)}
}
