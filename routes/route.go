package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	KenanganRoutes(e)
	UserRoutes(e)
	Authroutes(e)
}
