package controllers

import (
	"context"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type OptionStorage interface {
	insertOptionInDb(op *models.Option)
	listOptionInDb() []models.Option
}

type OptionService struct {
	db  *gorm.DB
	ctx context.Context
}

func (o *OptionService) listOptionInDb() []models.Option {
	var options = make([]models.Option, 0, 4)

	tx := o.db.WithContext(o.ctx)

	tx.Find(&options)

	return options
}

func (o *OptionService) insertOptionInDb(op *models.Option) {
	tx := o.db.WithContext(o.ctx)

	tx.Create(op)

}

func NewOptionStorage(ctx context.Context, db *gorm.DB) OptionStorage {
	return &OptionService{db: db, ctx: ctx}
}
