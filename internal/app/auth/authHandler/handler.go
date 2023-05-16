package authHandler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/auth"
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
	UseCase     auth.UseCase
	AuthManager authManager.AuthManager
}

func NewHandlerAuth(usecase auth.UseCase, authManager authManager.AuthManager) *HandlerAuth {
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
		Secure:   true,
	}
}

type OK struct {
	Message string `json:"message"`
}

func (h HandlerAuth) Register(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь находится в системе"), 0, "пользователь уже зарегестрирован")
	}

	var UserRequest models.UserJson
	err := ctx.Bind(&UserRequest)
	if err != nil || len(UserRequest.Username) == 0 || len(UserRequest.Password) == 0 {
		return tools.CustomError(ctx, err, 1, "битый json на авторизацию")
	}

	User, err := h.UseCase.Register(&models.UserJson{Username: UserRequest.Username, Password: UserRequest.Password})
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка при регистрации")
	}

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(User.Id)) // подставить id пользователя полученного из usecase
	if err != nil {
		return tools.CustomError(ctx, err, 2, "отрыгнул jsw token или что то связанное с ним")
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)

	result, _ := json.Marshal(User)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) Login(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь уже вошел в систему"), 0, "пользователь уже зарегестрирован")
	}

	var UserRequest models.UserJson
	err := ctx.Bind(&UserRequest)
	if err != nil || len(UserRequest.Username) == 0 || len(UserRequest.Password) == 0 {
		return tools.CustomError(ctx, err, 1, "битый json на логин")
	}

	User, err := h.UseCase.Login(&models.UserJson{Username: UserRequest.Username, Password: UserRequest.Password})
	if err != nil {
		return tools.CustomError(ctx, err, 2, "ошибка при логине")
	}

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(User.Id)) // подставить id пользователя полученного из usecase
	if err != nil {
		return tools.CustomError(ctx, err, 2, "отрыгнул jsw token или что то связанное с ним")
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)

	result, _ := json.Marshal(User)
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

	user1, err := h.UseCase.GetUserById(user.Id)
	if user == nil {
		return tools.CustomError(ctx, err, 0, "ошибка при получении пользователя")
	}

	result, _ := json.Marshal(user1)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h HandlerAuth) Update(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
	}

	fmt.Println(user)

	UserRequest := &models.UpdateJson{}
	err := ctx.Bind(&UserRequest)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "битый json на updateusername")
	}

	fmt.Println(UserRequest)

	user1, err := h.UseCase.Update(UserRequest, user.Id)
	if user == nil {
		return tools.CustomError(ctx, err, 0, "ошибка при получении пользователя")
	}

	result, _ := json.Marshal(user1)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

//func (h HandlerAuth) UpdatePassword(ctx echo.Context) error {
//	user := middleware.GetUserFromCtx(ctx)
//	if user == nil {
//		return tools.CustomError(ctx, errors.Errorf("пользователь не в системе"), 0, "ошибка при запросе пользователя")
//	}
//
//	UserRequest := &models.UpdatePasswordJson{}
//	err := ctx.Bind(&UserRequest)
//	if err != nil {
//		return tools.CustomError(ctx, err, 1, "битый json на updateusername")
//	}
//
//	user1, err := h.UseCase.UpdatePassword(UserRequest)
//	if err != nil {
//		return tools.CustomError(ctx, err, 0, "ошибка при обновлении пароля")
//	}
//
//	result, _ := json.Marshal(user1)
//	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
//	return ctx.JSONBlob(http.StatusOK, result)
//}
