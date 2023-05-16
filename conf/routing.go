package conf

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tp-study-ai/backend/internal/app/auth/authHandler"
	"github.com/tp-study-ai/backend/internal/app/bonusSystem/bonusSystemHandler"
	"github.com/tp-study-ai/backend/internal/app/chatGPT/chatGPTHandler"
	"github.com/tp-study-ai/backend/internal/app/like/likeHandler"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/ml/mlHandler"
	"github.com/tp-study-ai/backend/internal/app/task/taskHandler"
	"github.com/tp-study-ai/backend/internal/app/testis/testisHandler"
)

type ServerHandlers struct {
	TaskHandler        *taskHandler.HandlerTask
	AuthHandler        *authHandler.HandlerAuth
	TestisHandler      *testisHandler.HandlerTestis
	LikeHandler        *likeHandler.HandlerLike
	ChatGPTHandler     *chatGPTHandler.HandlerChatGPT
	MLHandler          *mlHandler.HandlerML
	BonusSystemHandler *bonusSystemHandler.HandlerBonusSystem
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

	router.GET("metrics", echo.WrapHandler(promhttp.Handler()))

	// testis
	router.POST("/api/check_solution", sh.TestisHandler.CheckSolution, mwChain...)

	// ml
	router.POST("/api/get_similar", sh.MLHandler.GetSimilar, mwChain...)
	router.GET("/api/recommendations", sh.MLHandler.Recommendations, mwChain...)
	router.GET("/api/cold_start", sh.MLHandler.ColdStart, mwChain...)

	// like
	router.POST("/api/like_task", sh.LikeHandler.LikeTask, mwChain...)
	router.POST("/api/delete_like", sh.LikeHandler.DeleteLike, mwChain...)
	router.GET("/api/get_like_tasks", sh.LikeHandler.GetLikeTasks, mwChain...)

	// chatGPT
	router.POST("/api/chat_gpt", sh.ChatGPTHandler.ChatGPT, mwChain...)

	router.GET("/api/get_task", sh.TaskHandler.GetTask, mwChain...)
	router.GET("/api/get_task_by_id", sh.TaskHandler.GetTaskById, mwChain...)
	router.GET("/api/tasks_list", sh.TaskHandler.GetTaskByLimit, mwChain...)
	router.GET("/api/get_tags", sh.TaskHandler.GetTags, mwChain...)
	router.GET("/api/get_send_tasks", sh.TaskHandler.GetSendTasks, mwChain...)
	router.GET("/api/get_send_tasks_by_task_id", sh.TaskHandler.GetSendTaskByTaskId, mwChain...)
	router.GET("/api/get_done_task", sh.TaskHandler.GetDoneTask, mwChain...)
	router.GET("/api/get_not_done_task", sh.TaskHandler.GetNotDoneTask, mwChain...)
	router.POST("/api/set_difficulty", sh.TaskHandler.SetDifficultyTask, mwChain...)

	// bonusSystem
	router.GET("api/calendar", sh.BonusSystemHandler.GetCountTaskOfDate, mwChain...)
	router.GET("api/shock_mode", sh.BonusSystemHandler.GetChockMode, mwChain...)

	//auth
	router.POST("/api/register", sh.AuthHandler.Register, mwChain...)
	router.POST("/api/login", sh.AuthHandler.Login, mwChain...)
	router.GET("/api/logout", sh.AuthHandler.Logout, mwChain...)
	router.GET("/api/get_user", sh.AuthHandler.GetUserById, mwChain...)
	router.POST("/api/update", sh.AuthHandler.Update, mwChain...)
}
