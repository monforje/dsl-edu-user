package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days left: %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
