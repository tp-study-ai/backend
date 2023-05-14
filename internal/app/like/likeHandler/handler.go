package likeHandler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	che "github.com/tp-study-ai/backend/internal/app/like"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
	"net/http"
	"strconv"
)

type HandlerLike struct {
	UseCase che.UseCase
}

func NewHandlerLike(useCase che.UseCase) *HandlerLike {
	return &HandlerLike{
		UseCase: useCase,
	}
}

func (h HandlerLike) LikeTask(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	fmt.Println("Id users:", user.Id)

	var like models.LikeJson
	if err := ctx.Bind(&like); err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирования запроса")
	}

	fmt.Println("like", like)

	err := h.UseCase.LikeTask(models.LikeJson{UserId: user.Id, TaskId: like.TaskId})
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка оценки задачи")
	}
	result, err := json.Marshal(models.Message{
		Message: "лайк поставлен",
	})
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerLike) DeleteLike(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	fmt.Println("Id users:", user.Id)

	var like models.LikeJson
	if err := ctx.Bind(&like); err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирования запроса")
	}

	fmt.Println("like", like)

	err := h.UseCase.DeleteLike(models.LikeJson{UserId: user.Id, TaskId: like.TaskId})
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка оценки задачи")
	}
	result, err := json.Marshal(models.Message{
		Message: "лайк удален",
	})
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerLike) GetLikeTasks(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	fmt.Println(user.Id)

	tasks, err := h.UseCase.GetLikeTask(user.Id)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка полчения оценных задач")
	}

	fmt.Println(tasks)

	result, err := json.Marshal(tasks)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования оцененых задач")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
