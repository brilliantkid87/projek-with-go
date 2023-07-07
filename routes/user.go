package routes

import (
	"kenangan/handler"
	"kenangan/pkg/mysql"
	"kenangan/repositories"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs" // Import the generated Swagger docs
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerUser(userRepository)

	// Routes with Swagger annotations
	/**
	* @swagger
	* /user:
	*   post:
	*     summary: Create a user
	*     tags:
	*       - User
	*     parameters:
	*       - in: body
	*         name: user
	*         required: true
	*         schema:
	*           $ref: '#/definitions/User'
	*     responses:
	*       '200':
	*         description: User created successfully
	*       '400':
	*         description: Invalid request data
	 */
	e.POST("/user", h.CreateUser)

	/**
	* @swagger
	* /user/{id}:
	*   get:
	*     summary: Get a user by ID
	*     tags:
	*       - User
	*     parameters:
	*       - in: path
	*         name: id
	*         required: true
	*         description: User ID
	*         schema:
	*           type: integer
	*     responses:
	*       '200':
	*         description: User found
	*       '404':
	*         description: User not found
	 */
	e.GET("/user/:id", h.GetUserByID)

	// Serve Swagger documentation
	e.GET("/docs/*", echoSwagger.WrapHandler)
}
