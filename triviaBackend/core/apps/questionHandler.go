package apps

import (
	"context"
	question "trivia-backend/core/controllers"
	"trivia-backend/core/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type QuestionHTTPService struct {
	gtw question.QuestionGateway
}

func (q *QuestionHTTPService) ListQuestionHandler(c echo.Context) error {

	status, response := q.gtw.ListQuestion()

	return c.JSON(status, response)
}

func (q *QuestionHTTPService) CreateQuestionHandler(c echo.Context) error {

	qu := new(models.Question)

	c.Bind(qu)

	//qu.LevelID = helpers.UUID()

	status, response := q.gtw.CreateQuestion(qu)
	return c.JSON(status, response)
}

func NewQuestionHTTPService(ctx context.Context, db *gorm.DB) *QuestionHTTPService {
	return &QuestionHTTPService{question.NewQuestionGateway(ctx, db)}
}
