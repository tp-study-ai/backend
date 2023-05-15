package chatGPTHandler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	che "github.com/tp-study-ai/backend/internal/app/chatGPT"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
	"net/http"
	"strconv"
)

type HandlerChatGPT struct {
	UseCase che.UseCase
}

func NewHandlerChatGPT(useCase che.UseCase) *HandlerChatGPT {
	return &HandlerChatGPT{
		UseCase: useCase,
	}
}

func (h HandlerChatGPT) ChatGPT(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	var ChatGPTRequest models.ChatGPT
	if err := ctx.Bind(&ChatGPTRequest); err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка формирования запроса")
	}

	response, err := h.UseCase.Chat(ChatGPTRequest)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "ошибка получения ответа")
	}

	result, err := json.Marshal(response)
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка формирования ответа")
	}

	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
