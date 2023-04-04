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
	//router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	//}))
	router.Use(echoMiddleware.CORSWithConfig(getCorsConfig()))
	mwChain := []echo.MiddlewareFunc{
		mw.AuthMiddleware,
	}

	router.GET("/api/get_task", sh.TaskHandler.GetTask, mwChain...)
	router.GET("/api/get_task_by_id", sh.TaskHandler.GetTaskById, mwChain...)
	router.GET("/api/tasks_list", sh.TaskHandler.GetTaskByLimit, mwChain...)
	router.POST("/api/check_solution", sh.TaskHandler.CheckSolution, mwChain...)

	router.POST("/api/register", sh.AuthHandler.Register, mwChain...)
}
