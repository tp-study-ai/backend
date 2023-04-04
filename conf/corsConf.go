package conf

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getCorsConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		//Skipper: middleware.DefaultSkipper,
		AllowOrigins: []string{
			"",
			//"http://localhost:3000",
			//"https://www.study-ai.ru",
			//"https://study-ai.ru",
			//"http://146.185.208.233:80",
			//"https://146.185.208.233:433",
		},
		AllowCredentials: true,
		//AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS, echo.PUT},
		//ExposeHeaders: []string{echo.HeaderXCSRFToken},
		MaxAge: 86400,
	}
}
