package conf

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/tp-study-ai/backend/internal/app/auth"
	"github.com/tp-study-ai/backend/internal/app/task"
)

type ServerHandlers struct {
	TaskHandler *task.HandlerTask
	AuthHandler *auth.HandlerAuth
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	//router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"*"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	//}))
	router.Use(echoMiddleware.CORSWithConfig(getCorsConfig()))

	router.GET("/api/get_task", sh.TaskHandler.GetTask)
	router.GET("/api/get_task_by_id", sh.TaskHandler.GetTaskById)
	router.GET("/api/tasks_list", sh.TaskHandler.GetTaskByLimit)
	router.POST("/api/check_solution", sh.TaskHandler.CheckSolution)

	router.POST("/api/register", sh.AuthHandler.Register)
}
