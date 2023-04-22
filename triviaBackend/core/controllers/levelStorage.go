package controllers

import (
	"context"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type LevelStorage interface {
	removeLevelInDb(id int64)
	editLevelInDb(lv *models.Level, id int64)
	insertLevelInDb(lv *models.Level)
	listLevelInDb() []models.Level
}

type LevelService struct {
	db  *gorm.DB
	ctx context.Context
}

func (ls *LevelService) listLevelInDb() []models.Level {

	var levels = make([]models.Level, 0, 4)

	tx := ls.db.WithContext(ls.ctx)

	tx.Find(&levels)

	return levels
}

func (ls *LevelService) insertLevelInDb(lv *models.Level) {
	tx := ls.db.WithContext(ls.ctx)

	tx.Create(lv)
}

func (ls *LevelService) editLevelInDb(lv *models.Level, id int64) {

	tx := ls.db.WithContext(ls.ctx)

	tx.Model(&lv).Where("level_id = ?", id).Updates(models.Level{LevelName: lv.LevelName})

}

func (ls *LevelService) removeLevelInDb(id int64) {
	tx := ls.db.WithContext(ls.ctx)

	tx.Delete(&models.Level{}, id)
}

func NewLevelStorage(ctx context.Context, db *gorm.DB) LevelStorage {

	return &LevelService{db: db, ctx: ctx}
}
