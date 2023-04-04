package tools

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/tp-study-ai/backend/internal/app/models"
	"net/http"
	"strconv"
)

func CustomError(ctx echo.Context, err error, number int, comment string) error {
	che := models.CustomError{
		Number:  number,
		Comment: comment,
		Error:   err.Error(),
	}
	result, _ := json.Marshal(che)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusInternalServerError, result)
}
