package controllers

import (
	"context"
	"net/http"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type LevelGateway interface {
	RemoveLevel(id int64) (int, models.Response)
	CreateLevel(lg *models.Level) (int, models.Response)
	ListLevel() (int, []models.Level)
	EditLevel(lg *models.Level, id int64) (int, models.Response)
}

type LevelLogic struct {
	Sl LevelStorage
}

func (lg *LevelLogic) ListLevel() (int, []models.Level) {

	return http.StatusOK, lg.Sl.listLevelInDb()

}

func (lg *LevelLogic) CreateLevel(le *models.Level) (int, models.Response) {

	for _, value := range lg.Sl.listLevelInDb() {
		if le.LevelName == value.LevelName || le.LevelName == 0 || le.LevelName < 0 {
			return http.StatusConflict, models.Response{
				Message: "The level is invalid",
				Success: false,
			}
		}
	}

	go lg.Sl.insertLevelInDb(le)

	return http.StatusAccepted, models.Response{
		Message: "Level Successfully Added.",
		Success: true,
	}
}

func (lg *LevelLogic) EditLevel(le *models.Level, id int64) (int, models.Response) {

	for _, value := range lg.Sl.listLevelInDb() {
		if le.LevelName == value.LevelName || le.LevelName == 0 || le.LevelName < 0 {
			return http.StatusConflict, models.Response{
				Message: "The level data is invalid",
				Success: false,
			}
		}
	}

	go lg.Sl.editLevelInDb(le, id)

	return http.StatusOK, models.Response{
		Message: "Level Successfully Edited.",
		Success: true,
	}
}

func (lg *LevelLogic) RemoveLevel(id int64) (int, models.Response) {
	for _, value := range lg.Sl.listLevelInDb() {
		if value.LevelID != id {
			return http.StatusConflict, models.Response{
				Message: "The Id level is invalid",
				Success: false,
			}
		}
	}
	go lg.Sl.removeLevelInDb(id)

	return http.StatusOK, models.Response{
		Message: "Level Successfully Remove",
		Success: true,
	}
}

func NewLevelGateway(ctx context.Context, db *gorm.DB) LevelGateway {
	return &LevelLogic{NewLevelStorage(ctx, db)}
}
