package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/conf"
	"github.com/tp-study-ai/backend/internal/app/auth/authHandler"
	"github.com/tp-study-ai/backend/internal/app/auth/authRepository"
	"github.com/tp-study-ai/backend/internal/app/auth/authUseCase"

	"github.com/tp-study-ai/backend/internal/app/bonusSystem/bonusSystemHandler"
	"github.com/tp-study-ai/backend/internal/app/bonusSystem/bonusSystemRepository"
	"github.com/tp-study-ai/backend/internal/app/bonusSystem/bonusSystemUseCase"

	"github.com/tp-study-ai/backend/internal/app/chatGPT/chatGPTHandler"
	"github.com/tp-study-ai/backend/internal/app/chatGPT/chatGPTRepository"
	"github.com/tp-study-ai/backend/internal/app/chatGPT/chatGPTUseCase"

	"github.com/tp-study-ai/backend/internal/app/ml/mlHandler"
	"github.com/tp-study-ai/backend/internal/app/ml/mlRepository"
	"github.com/tp-study-ai/backend/internal/app/ml/mlUseCase"

	"github.com/tp-study-ai/backend/internal/app/task/taskHandler"
	"github.com/tp-study-ai/backend/internal/app/task/taskRepository"
	"github.com/tp-study-ai/backend/internal/app/task/taskUseCase"

	"github.com/tp-study-ai/backend/internal/app/like/likeHandler"
	"github.com/tp-study-ai/backend/internal/app/like/likeRepository"
	"github.com/tp-study-ai/backend/internal/app/like/likeUseCase"

	"github.com/tp-study-ai/backend/internal/app/metrics"
	"github.com/tp-study-ai/backend/internal/app/middleware"

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

	taskRepository := taskRepository.NewRepositoryTask(pgxManager)
	taskUseCase := taskUseCase.NewUseCaseTask(taskRepository, config.Testis, config.Ml, config.MLRec, config.MLCS, config.CG)
	taskHandler := taskHandler.NewHandlerTask(taskUseCase)

	testisRepository := testisRepository.NewRepositoryTask(pgxManager)
	testisUseCase := testisUseCase.NewUseCaseTestis(testisRepository, config.Testis, config.Ml, config.MLRec, config.MLCS, config.CG)
	testisHandler := testisHandler.NewHandlerTestis(testisUseCase)

	likeRepository := likeRepository.NewRepositoryLike(pgxManager)
	likeUseCase := likeUseCase.NewUseCaseLike(likeRepository)
	likeHandler := likeHandler.NewHandlerLike(likeUseCase)

	chatGPTRepository := chatGPTRepository.NewRepositoryChatGPT(pgxManager)
	chatGPTUseCase := chatGPTUseCase.NewUseCaseChatGPT(chatGPTRepository, config.CG)
	chatGPTHandler := chatGPTHandler.NewHandlerChatGPT(chatGPTUseCase)

	mlRepository := mlRepository.NewRepositoryML(pgxManager)
	mlUseCase := mlUseCase.NewUseCaseML(mlRepository, config.Testis, config.Ml, config.MLRec, config.MLCS, config.CG)
	mlHandler := mlHandler.NewHandlerML(mlUseCase)

	bonusSystemRepository := bonusSystemRepository.NewRepositoryBonusSystem(pgxManager)
	bonusSystemUseCase := bonusSystemUseCase.NewUseCaseBonusSystem(bonusSystemRepository)
	bonusSystemHandler := bonusSystemHandler.NewHandlerBonusSystem(bonusSystemUseCase)

	authRepository := authRepository.NewRepositoryAuth(pgxManager)
	authUseCase := authUseCase.NewUseCaseAuth(authRepository)
	authHandler := authHandler.NewHandlerAuth(authUseCase, jwtManager)

	router := echo.New()

	m, err := metrics.CreateNewMetric("main")
	if err != nil {
		panic(err)
	}

	router.Use(m.CollectMetrics)

	serverRouting := conf.ServerHandlers{
		TaskHandler:        taskHandler,
		TestisHandler:      testisHandler,
		AuthHandler:        authHandler,
		LikeHandler:        likeHandler,
		ChatGPTHandler:     chatGPTHandler,
		MLHandler:          mlHandler,
		BonusSystemHandler: bonusSystemHandler,
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
