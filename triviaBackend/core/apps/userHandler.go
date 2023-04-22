package apps

import (
	"context"
	user "trivia-backend/core/controllers"
	"trivia-backend/core/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHTTPService struct {
	gtw user.UserGateway
}

func (u *UserHTTPService) ListUserHandler(c echo.Context) error {

	status, response := u.gtw.ListUser()

	return c.JSON(status, response)
}

func (u *UserHTTPService) CreateUserHandler(c echo.Context) error {

	myUser := new(models.User)

	c.Bind(myUser)

	myUser.Password, _ = myUser.HashPassword()

	status, response := u.gtw.CreateUser(myUser)

	return c.JSON(status, response)
}

func NewUserHTTPService(ctx context.Context, db *gorm.DB) *UserHTTPService {

	return &UserHTTPService{user.NewUserGateway(ctx, db)}
}
