package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/conf"
	"github.com/tp-study-ai/backend/internal/app/auth"
	"github.com/tp-study-ai/backend/internal/app/task"
	"github.com/tp-study-ai/backend/tools"
	"github.com/tp-study-ai/backend/tools/authManager/jwtManager"
	"log"
	"net/http"
)

func main() {
	pgxManager, err := tools.NewPostgres()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	jwtManager := jwtManager.NewJwtManager()

	taskRepo := task.NewRepositoryTask(pgxManager)
	taskUcase := task.NewUseCaseTask(taskRepo)
	taskHandler := task.NewHandlerTask(taskUcase)

	authHandler := auth.NewHandlerAuth(jwtManager)

	router := echo.New()

	serverRouting := conf.ServerHandlers{
		TaskHandler: taskHandler,
		AuthHandler: authHandler,
	}

	serverRouting.ConfigureRouting(router)

	httpServ := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
