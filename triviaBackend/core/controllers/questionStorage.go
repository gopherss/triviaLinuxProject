package controllers

import (
	"context"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type QuestionStorage interface {
	insertQuestionInDb(qu *models.Question)
	listQuestionInDb() []models.Question
}

type QuestionService struct {
	db  *gorm.DB
	ctx context.Context
}

func (q *QuestionService) listQuestionInDb() []models.Question {
	var questions = make([]models.Question, 0, 4)
	tx := q.db.WithContext(q.ctx)

	tx.Find(&questions)

	return questions
}

func (q *QuestionService) insertQuestionInDb(qu *models.Question) {

	tx := q.db.WithContext(q.ctx)

	tx.Create(qu)

}

func NewQuestionStorage(ctx context.Context, db *gorm.DB) QuestionStorage {
	return &QuestionService{db: db, ctx: ctx}
}
