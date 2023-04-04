package conf

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCorsConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
		AllowCredentials: true,
		//AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS, echo.PUT},
		//ExposeHeaders: []string{echo.HeaderXCSRFToken},
		MaxAge: 86400,
	}
}
