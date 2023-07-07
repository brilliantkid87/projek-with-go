package handler

import (
	kenangandto "kenangan/dto/kenangan"
	dto "kenangan/dto/result"
	"kenangan/model"
	"kenangan/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KenanganHandler struct {
	KenanganRepository repositories.KenanganRepository
}

func HandlerKenangan(KenanganRepository repositories.KenanganRepository) *KenanganHandler {
	return &KenanganHandler{KenanganRepository}
}

func (h *KenanganHandler) CreateKenangan(c echo.Context) error {
	request := new(kenangandto.CreateKenanganRequested)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	kenangan := model.Kenangan{
		ID:          request.ID,
		Title:       request.Title,
		Image:       request.Image,
		Description: request.Description,
		UserID:      request.UserID,
	}
	data, err := h.KenanganRepository.CreateKenangan(kenangan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: data})
}

func (h *KenanganHandler) GetKenanganByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	kenangan, err := h.KenanganRepository.GetKenanganByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: kenangan})
}
