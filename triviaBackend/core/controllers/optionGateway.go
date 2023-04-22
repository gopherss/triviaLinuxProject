package controllers

import (
	"context"
	"net/http"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type OptionGateway interface {
	CreateOption(o *models.Option) (int, models.Response)
	ListOption() (int, []models.Option)
}

type OptionLogic struct {
	So OptionStorage
}

func (o *OptionLogic) ListOption() (int, []models.Option) {

	return http.StatusOK, o.So.listOptionInDb()
}

func (o *OptionLogic) CreateOption(op *models.Option) (int, models.Response) {

	go o.So.insertOptionInDb(op)

	return http.StatusAccepted, models.Response{
		Message: "Option Successfully addedd.",
		Success: true,
	}
}

func NewOptionGateway(ctx context.Context, db *gorm.DB) OptionGateway {

	return &OptionLogic{NewOptionStorage(ctx, db)}
}
