package conf

import (
	"github.com/labstack/echo/v4"
	"github.com/tp-study-ai/backend/internal/app/task"
)

type ServerHandlers struct {
	TaskHandler *task.HandlerTask
}

func (sh *ServerHandlers) ConfigureRouting(router *echo.Echo) {
	router.GET("get_task", sh.TaskHandler.GetTask)
	router.GET("get_task_by_id", sh.TaskHandler.GetTaskById)
}