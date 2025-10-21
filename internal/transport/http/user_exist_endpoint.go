package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExistRequest struct {
	TelegramID int64 `json:"telegram_id"`
}

type ExistResponse struct {
	Exist bool `json:"exist"`
}

func (e *Endpoint) Exist(ctx echo.Context) error {
	var req ExistRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid JSON",
		})
	}

	exist, err := e.s.MongoService().IsExist(req.TelegramID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, ExistResponse{Exist: exist})
}
