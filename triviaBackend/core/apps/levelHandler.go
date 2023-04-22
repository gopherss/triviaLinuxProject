package apps

import (
	"context"
	"strconv"
	level "trivia-backend/core/controllers"
	"trivia-backend/core/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LevelHTTPService struct {
	gtw level.LevelGateway
}

func (lh *LevelHTTPService) ListLevelHandler(c echo.Context) error {
	status, response := lh.gtw.ListLevel()

	return c.JSON(status, response)
}

func (lh *LevelHTTPService) CreateQuestionHandler(c echo.Context) error {
	le := new(models.Level)

	c.Bind(le)

	status, response := lh.gtw.CreateLevel(le)

	return c.JSON(status, response)
}

func (lh *LevelHTTPService) EditLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	le := new(models.Level)

	c.Bind(le)

	status, response := lh.gtw.EditLevel(le, int64(id))

	return c.JSON(status, response)
}

func (lh *LevelHTTPService) RemoveLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	status, response := lh.gtw.RemoveLevel(int64(id))

	return c.JSON(status, response)
}

func NewLevelHTTPService(ctx context.Context, db *gorm.DB) *LevelHTTPService {

	return &LevelHTTPService{level.NewLevelGateway(ctx, db)}

}
