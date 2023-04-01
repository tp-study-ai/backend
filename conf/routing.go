package conf

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tp-study-ai/backend/internal/app/task"
	"net/http"
)

type ServerHandlers struct {
	TaskHandler *task.HandlerTask
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	router.GET("/api/get_task", sh.TaskHandler.GetTask)
	router.GET("/api/get_task_by_id", sh.TaskHandler.GetTaskById)
	router.POST("/api/check_solution", sh.TaskHandler.CheckSolution)
	router.POST("/api/my_send_solution", sh.TaskHandler.MySendSolution)
}
