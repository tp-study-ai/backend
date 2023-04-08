package conf

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/tp-study-ai/backend/internal/app/auth"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/task"
)

type ServerHandlers struct {
	TaskHandler *task.HandlerTask
	AuthHandler *auth.HandlerAuth
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo, mw *middleware.CommonMiddleware) {
	router.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST},
		MaxAge:           86400,
	}))
	//router.Use(echoMiddleware.CORSWithConfig(getCorsConfig()))
	mwChain := []echo.MiddlewareFunc{
		mw.AuthMiddleware,
	}

	router.GET("/api/get_task", sh.TaskHandler.GetTask, mwChain...)
	router.GET("/api/get_task_by_id", sh.TaskHandler.GetTaskById, mwChain...)
	router.GET("/api/tasks_list", sh.TaskHandler.GetTaskByLimit, mwChain...)
	router.POST("/api/check_solution", sh.TaskHandler.CheckSolution, mwChain...)
	router.GET("/api/get_tags", sh.TaskHandler.GetTags, mwChain...)

	router.POST("/api/register", sh.AuthHandler.Register, mwChain...)
	router.POST("/api/login", sh.AuthHandler.Login, mwChain...)
	router.GET("/api/logout", sh.AuthHandler.Logout, mwChain...)
	router.GET("/api/get_user", sh.AuthHandler.GetUserById, mwChain...)
}
