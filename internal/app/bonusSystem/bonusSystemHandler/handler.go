package bonusSystemHandler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/bonusSystem"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/tools"
	"net/http"
	"strconv"
)

type HandlerBonusSystem struct {
	UseCase bonusSystem.UseCase
}

func NewHandlerBonusSystem(useCase bonusSystem.UseCase) *HandlerBonusSystem {
	return &HandlerBonusSystem{
		UseCase: useCase,
	}
}

func (h HandlerBonusSystem) GetCountTaskOfDate(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	days, err := h.UseCase.GetCountTaskOfDate(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 0, "ошибка получения задач")
	}

	result, err := json.Marshal(days)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerBonusSystem) GetChockMode(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	shockMode, err := h.UseCase.GetShockMode(int(user.Id))
	if err != nil {
		return tools.CustomError(ctx, err, 0, "ошибка получения ударного режима")
	}

	result, err := json.Marshal(shockMode)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
