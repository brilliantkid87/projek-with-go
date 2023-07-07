package routes

import (
	"kenangan/handler"
	"kenangan/pkg/middleware"
	"kenangan/pkg/mysql"
	"kenangan/repositories"

	"github.com/labstack/echo/v4"
)

func Authroutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handler.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
