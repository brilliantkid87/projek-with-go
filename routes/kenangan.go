package routes

import (
	"kenangan/handler"
	"kenangan/pkg/mysql"
	"kenangan/repositories"

	"github.com/labstack/echo/v4"
)

func KenanganRoutes(e *echo.Group) {
	kenanganRepository := repositories.RepositoryKenangan(mysql.DB)
	h := handler.HandlerKenangan(kenanganRepository)

	e.POST("/kenangan", h.CreateKenangan)
	e.GET("/kenangan/:id", h.GetKenanganByID)
}
