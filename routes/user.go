package routes

import (
	"kenangan/handler"
	"kenangan/pkg/mysql"
	"kenangan/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerUser(userRepository)

	e.POST("/user", h.CreateUser)
	e.GET("/user/:id", h.GetUserByID)
}
