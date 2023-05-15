package mlHandler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	che "github.com/tp-study-ai/backend/internal/app/ml"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
	"net/http"
	"strconv"
)

type HandlerML struct {
	UseCase che.UseCase
}

func NewHandlerML(useCase che.UseCase) *HandlerML {
	return &HandlerML{
		UseCase: useCase,
	}
}

func (h HandlerML) GetSimilar(ctx echo.Context) error {
	var che models.SimilarRequest
	if err := ctx.Bind(&che); err != nil {
		return tools.CustomError(ctx, err, 0, "ошибка формирования запроса")
	}

	tasks, err := h.UseCase.GetSimilar(che)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка работы рекомендательной системы")
	}

	result, err := json.Marshal(tasks)
	if err != nil {
		return tools.CustomError(ctx, err, 3, "ошибка формирования рекомендаций")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerML) Recommendations(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	response, err := h.UseCase.Recommendations(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка получения рекомендаций")
	}

	result, err := json.Marshal(response)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerML) ColdStart(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	response, err := h.UseCase.ColdStart(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка получения задачи холодного старта")
	}
	if response == nil && err == nil {
		result, err1 := json.Marshal(models.Message{Message: "Холодный старт успешно пройден"})
		if err1 != nil {
			return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
		}

		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusOK, result)
	}

	result, err := json.Marshal(response)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
