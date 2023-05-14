package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/conf"
	"github.com/tp-study-ai/backend/internal/app/auth"

	"github.com/tp-study-ai/backend/internal/app/like/likeHandler"
	"github.com/tp-study-ai/backend/internal/app/like/likeRepository"
	"github.com/tp-study-ai/backend/internal/app/like/likeUseCase"

	"github.com/tp-study-ai/backend/internal/app/metrics"
	"github.com/tp-study-ai/backend/internal/app/middleware"
	"github.com/tp-study-ai/backend/internal/app/task"

	"github.com/tp-study-ai/backend/internal/app/testis/testisHandler"
	"github.com/tp-study-ai/backend/internal/app/testis/testisRepository"
	"github.com/tp-study-ai/backend/internal/app/testis/testisUseCase"

	"github.com/tp-study-ai/backend/tools"
	"github.com/tp-study-ai/backend/tools/authManager/jwtManager"
	"log"
	"net/http"
)

func main() {
	configPath := flag.String("config", "./backConfig/conf.toml", "path to config file")
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

	jwtManager := jwtManager.NewJwtManager(config.JWT)

	taskRepo := task.NewRepositoryTask(pgxManager)
	taskUcase := task.NewUseCaseTask(taskRepo, config.Testis, config.Ml, config.MLRec, config.MLCS, config.CG)
	taskHandler := task.NewHandlerTask(taskUcase)

	testisRepo := testisRepository.NewRepositoryTask(pgxManager)
	testisUseCase := testisUseCase.NewUseCaseTestis(testisRepo, config.Testis, config.Ml, config.MLRec, config.MLCS, config.CG)
	testisHandler := testisHandler.NewHandlerTestis(testisUseCase)

	likeRepo := likeRepository.NewRepositoryLike(pgxManager)
	likeUseCase := likeUseCase.NewUseCaseLike(likeRepo)
	likeHandler := likeHandler.NewHandlerLike(likeUseCase)

	authRepo := auth.NewRepositoryAuth(pgxManager)
	authUcase := auth.NewUseCaseAuth(authRepo)
	authHandler := auth.NewHandlerAuth(authUcase, jwtManager)

	router := echo.New()

	m, err := metrics.CreateNewMetric("main")
	if err != nil {
		panic(err)
	}

	router.Use(m.CollectMetrics)

	serverRouting := conf.ServerHandlers{
		TaskHandler:   taskHandler,
		TestisHandler: testisHandler,
		AuthHandler:   authHandler,
		LikeHandler:   likeHandler,
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
