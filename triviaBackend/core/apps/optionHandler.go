package apps

import (
	"context"
	option "trivia-backend/core/controllers"
	"trivia-backend/core/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OptionHTTPService struct {
	gtw option.OptionGateway
}

func (o *OptionHTTPService) ListOptionHandler(c echo.Context) error {
	status, response := o.gtw.ListOption()

	return c.JSON(status, response)
}

func (o *OptionHTTPService) CreateOptionHandler(c echo.Context) error {
	op := new(models.Option)

	c.Bind(op)

	status, response := o.gtw.CreateOption(op)

	return c.JSON(status, response)
}

func NewOptionHTTPService(ctx context.Context, db *gorm.DB) *OptionHTTPService {
	return &OptionHTTPService{option.NewOptionGateway(ctx, db)}
}
