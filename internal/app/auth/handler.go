package auth

import (
	"github.com/labstack/echo/v4"
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
	AuthManager authManager.AuthManager
}

func NewHandlerAuth(authManager authManager.AuthManager) *HandlerAuth {
	return &HandlerAuth{
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
	var reg models.Register
	err := ctx.Bind(&reg)
	if err != nil {
		return tools.CustomError(ctx, err, 1, "битый json на авторизацию")
	}

	// тут должен быть поход в usecase

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
