package auth

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
	"github.com/tp-study-ai/backend/tools/authManager"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	tokenCookieKey = "token"
)

type HandlerAuth struct {
	UseCase     UseCase
	AuthManager authManager.AuthManager
}

func NewHandlerAuth(usecase UseCase, authManager authManager.AuthManager) *HandlerAuth {
	return &HandlerAuth{
		UseCase:     usecase,
		AuthManager: authManager,
	}
}

func createTokenCookie(token string, domen string, exp time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     tokenCookieKey,
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(exp),
		Domain:   domen,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
	}
}

type OK struct {
	Message string `json:"message"`
}

func (h HandlerAuth) Register(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь находится в системе"), 0, "пользователь уже зарегестрирован")
	}

	var reg models.UserJson
	err := ctx.Bind(&reg)
	if err != nil || len(reg.Username) == 0 || len(reg.Password) == 0 {
		return tools.CustomError(ctx, err, 1, "битый json на авторизацию")
	}
	fmt.Println(reg)

	UserId, err := h.UseCase.Register(&models.UserJson{Username: reg.Username, Password: reg.Password})
	if err != nil || UserId == 0 {
		return tools.CustomError(ctx, err, 2, "ошибка при регистрации")
	}

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(1)) // подставить id пользователя полученного из usecase
	if err != nil {
		return tools.CustomError(ctx, err, 2, "отрыгнул jsw token или что то связанное с ним")
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)

	var che = OK{
		Message: "успешно зарегестрировались",
	}

	result, _ := json.Marshal(che)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) Login(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь уже вошел в систему"), 0, "пользователь уже зарегестрирован")
	}

	var reg models.UserJson
	err := ctx.Bind(&reg)
	if err != nil || len(reg.Username) == 0 || len(reg.Password) == 0 {
		return tools.CustomError(ctx, err, 1, "битый json на логин")
	}

	b, err := h.UseCase.Login(&models.UserJson{Username: reg.Username, Password: reg.Password})
	if err != nil || b == false {
		return tools.CustomError(ctx, err, 2, "ошибка при логине")
	}

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(1)) // подставить id пользователя полученного из usecase
	if err != nil {
		return tools.CustomError(ctx, err, 2, "отрыгнул jsw token или что то связанное с ним")
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	var che = OK{
		Message: "успешно вошли в систему",
	}

	ctx.SetCookie(tokenCookie)

	result, _ := json.Marshal(che)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) Logout(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 2, "ошибка при логине")
	}
	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	resetTokenCookie := createTokenCookie("", host, -time.Hour)

	ctx.SetCookie(resetTokenCookie)

	var che = OK{
		Message: "успешно вышли из системы",
	}

	result, _ := json.Marshal(che)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) GetUserById(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	user1, err := h.UseCase.GetUserById(int(user.Id))
	if user == nil {
		return tools.CustomError(ctx, err, 0, "ошибка при получении пользователя")
	}

	result, _ := json.Marshal(user1)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
