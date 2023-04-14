package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/conf"
	"github.com/tp-study-ai/backend/internal/app/auth"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/task"
	"github.com/tp-study-ai/backend/tools"
	"github.com/tp-study-ai/backend/tools/authManager/jwtManager"
	"log"
	"net/http"
)

func main() {
	var configPath *string

	if false {
		configPath = flag.String("config", "./tools/conf.toml", "path to config file")
	} else {
		configPath = flag.String("config", "./tools/prod.yml", "path to config file")
	}
	flag.Parse()

	config := tools.NewConfig()

	err := tools.ReadConfigFile(*configPath, config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error reading config"))
	}
	fmt.Println(config)

	pgxManager, err := tools.NewPostgres(config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	jwtManager := jwtManager.NewJwtManager()

	taskRepo := task.NewRepositoryTask(pgxManager)
	taskUcase := task.NewUseCaseTask(taskRepo, config.Testis, config.Ml)
	taskHandler := task.NewHandlerTask(taskUcase)

	authRepo := auth.NewRepositoryAuth(pgxManager)
	authUcase := auth.NewUseCaseAuth(authRepo)
	authHandler := auth.NewHandlerAuth(authUcase, jwtManager)

	router := echo.New()

	serverRouting := conf.ServerHandlers{
		TaskHandler: taskHandler,
		AuthHandler: authHandler,
	}

	comonMw := middleware.NewCommonMiddleware(jwtManager)
	serverRouting.ConfigureRouting(router, &comonMw)

	httpServ := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
