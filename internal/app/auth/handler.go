package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/models"
	"github.com/tp-study-ai/backend/tools"
	"github.com/tp-study-ai/backend/tools/authManager"
	"net"
	"net/http"
	"time"
)

const (
	tokenCookieKey    = "token"
	CSRFCookieName    = "_csrf"
	avatarMaxSize     = 4000000
	updateUserMaxSize = 1000
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
	}
}

type OK struct {
	Message string `json:"message"`
}

func (h HandlerAuth) Register(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return tools.CustomError(ctx, errors.Errorf("пользователь уже зарегестрирован"), 0, "пользователь уже зарегестрирован")
	}

	var reg models.UserJson
	err := ctx.Bind(&reg)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "битый json на авторизацию")
	}

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

	var che = OK{
		Message: "ok",
	}

	ctx.SetCookie(tokenCookie)
	return ctx.JSON(http.StatusOK, che)
	//return ctx.JSONBlob(http.StatusOK, result)
}
