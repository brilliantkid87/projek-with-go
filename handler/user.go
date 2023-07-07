package handler

import (
	dto "kenangan/dto/result"
	userdto "kenangan/dto/user"
	"kenangan/model"
	"kenangan/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepository repositories.UserRespository
}

func HandlerUser(UserRepository repositories.UserRespository) *UserHandler {
	return &UserHandler{UserRepository}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags User
// @Accept json
// @Produce json
// @Param user body userdto.CreateUserRequest true "User object to be created"
// @Success 200 {object} dto.SuccessResult
// @Failure 400 {object} dto.ErrorResult
// @Failure 500 {object} dto.ErrorResult
// @Router /user [post]

func (h *UserHandler) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	user := model.User{
		ID:       request.ID,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Kenangan: []model.Kenangan{},
	}
	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: data})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user information by the provided ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.SuccessResult
// @Failure 400 {object} dto.ErrorResult
// @Router /user/{id} [get]

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: user})
}
