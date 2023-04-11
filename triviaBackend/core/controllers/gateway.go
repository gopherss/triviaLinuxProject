package controllers

import (
	"context"
	"net/http"
	"trivia-backend/core/models"

	"gorm.io/gorm"
)

type QuestionGateway interface {
	CreateQuestion(q *models.Question) (int, models.Response)
	ListQuestion() (int, []models.Question)
}

type QuestionLogic struct {
	Sq QuestionStorage
}

func (q *QuestionLogic) ListQuestion() (int, []models.Question) {

	return http.StatusOK, q.Sq.listQuestionInDb()
}

func (q *QuestionLogic) CreateQuestion(qu *models.Question) (int, models.Response) {

	go q.Sq.insertQuestionInDb(qu)

	return http.StatusAccepted, models.Response{
		Message: "Question Successfully added.",
		Success: true,
	}

}

func NewQuestionGateway(ctx context.Context, db *gorm.DB) QuestionGateway {
	return &QuestionLogic{NewQuestionStorage(ctx, db)}
}
