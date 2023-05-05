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
		AllowOrigins:     []string{"https://study-ai.ru/"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.POST},
		MaxAge:           86400,
	}))
	mwChain := []echo.MiddlewareFunc{
		mw.AuthMiddleware,
	}

	router.GET("/api/get_task", sh.TaskHandler.GetTask, mwChain...)
	router.GET("/api/get_task_by_id", sh.TaskHandler.GetTaskById, mwChain...)
	router.GET("/api/tasks_list", sh.TaskHandler.GetTaskByLimit, mwChain...)
	router.POST("/api/check_solution", sh.TaskHandler.CheckSolution, mwChain...)
	router.GET("/api/get_tags", sh.TaskHandler.GetTags, mwChain...)

	router.POST("/api/get_similar", sh.TaskHandler.GetSimilar, mwChain...)

	router.GET("/api/get_send_tasks", sh.TaskHandler.GetSendTasks, mwChain...)
	router.GET("/api/get_send_tasks_by_task_id", sh.TaskHandler.GetSendTaskByTaskId, mwChain...)
	router.POST("/api/like_task", sh.TaskHandler.LikeTask, mwChain...)
	router.POST("/api/delete_like", sh.TaskHandler.DeleteLike, mwChain...)
	router.GET("/api/get_like_tasks", sh.TaskHandler.GetLikeTasks, mwChain...)
	router.GET("/api/get_done_task", sh.TaskHandler.GetDoneTask, mwChain...)
	router.GET("/api/get_not_done_task", sh.TaskHandler.GetNotDoneTask, mwChain...)
	router.POST("/api/set_difficulty", sh.TaskHandler.SetDifficultyTask, mwChain...)

	router.GET("/api/recommendations", sh.TaskHandler.Recommendations, mwChain...)
	router.GET("/api/cold_start", sh.TaskHandler.ColdStart, mwChain...)

	router.GET("api/calendar", sh.TaskHandler.GetCountTaskOfDate, mwChain...)
	router.GET("api/shock_mode", sh.TaskHandler.GetChockMode, mwChain...)

	router.POST("/api/register", sh.AuthHandler.Register, mwChain...)
	router.POST("/api/login", sh.AuthHandler.Login, mwChain...)
	router.GET("/api/logout", sh.AuthHandler.Logout, mwChain...)
	router.GET("/api/get_user", sh.AuthHandler.GetUserById, mwChain...)
	router.POST("/api/update", sh.AuthHandler.Update, mwChain...)
}
