package testisHandler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/models"
	che "github.com/tp-study-ai/backend/internal/app/testis"
	"github.com/tp-study-ai/backend/tools"
	"net/http"
	"strconv"
)

type HandlerTestis struct {
	UseCase che.UseCase
}

func NewHandlerTestis(useCase che.UseCase) *HandlerTestis {
	return &HandlerTestis{
		UseCase: useCase,
	}
}

func (h HandlerTestis) CheckSolution(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	var solution models.CheckSolutionRequest
	if err := ctx.Bind(&solution); err != nil {
		return tools.CustomError(ctx, err, 1, "неверное формирования запроса")
	}

	test := &models.CheckSolutionUseCaseResponse{}
	testisResponse, err := h.UseCase.CheckSolution(solution, int(user.Id))
	if err != nil || testisResponse == test {
		return tools.CustomError(ctx, err, 2, "ошибка тестирования задачи")
	}

	result, err := json.Marshal(testisResponse)
	if err != nil {
		return tools.CustomError(ctx, err, 3, "ошибка формирования ответа тестирования")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
