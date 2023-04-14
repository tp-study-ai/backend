package main

import (
	"github.com/joho/godotenv"
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
	"os"
	"strconv"
)

func main() {
	Db := &tools.DB{}

	var Secret1 string
	var Secret2 string

	if false {
		if err := godotenv.Load(".env"); err != nil {
			log.Print("No .env file found")
		}
		log.Print("Find .env ")

		Db.User = os.Getenv("User")
		Db.Dbname = os.Getenv("Dbname")
		Db.Password = os.Getenv("Password")
		Db.Host = os.Getenv("Host")
		Db.Port, _ = strconv.ParseInt(os.Getenv("Port"), 10, 64)
		Secret1 = os.Getenv("testis")
		Secret2 = os.Getenv("ml")
	} else {
		if err := godotenv.Load(".env.prod"); err != nil {
			log.Print("No .env.prod file found")
		}
		log.Print("Find .env.prod ")

		Db.User = os.Getenv("User")
		Db.Dbname = os.Getenv("Dbname")
		Db.Password = os.Getenv("Password")
		Db.Host = os.Getenv("Host")
		Db.Port, _ = strconv.ParseInt(os.Getenv("Port"), 10, 64)
		Secret1 = os.Getenv("testis")
		Secret2 = os.Getenv("ml")
	}

	pgxManager, err := tools.NewPostgres(Db)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	jwtManager := jwtManager.NewJwtManager()

	taskRepo := task.NewRepositoryTask(pgxManager)
	taskUcase := task.NewUseCaseTask(taskRepo, Secret1, Secret2)
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
