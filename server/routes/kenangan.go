package routes

import (
	"kenangan/handler"
	"kenangan/pkg/mysql"
	"kenangan/repositories"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func KenanganRoutes(e *echo.Group) {
	kenanganRepository := repositories.RepositoryKenangan(mysql.DB)
	h := handler.HandlerKenangan(kenanganRepository)

	e.POST("/kenangan", h.CreateKenangan)
	e.GET("/kenangan/:id", h.GetKenanganByID)
	e.GET("/swagger/*any", echoSwagger.WrapHandler)
}
