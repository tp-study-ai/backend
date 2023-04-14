package main

import (
	"fmt"
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
		fmt.Println(Db.User)
		Db.Dbname = os.Getenv("Dbname")
		fmt.Println(Db.Dbname)
		Db.Password = os.Getenv("Password")
		fmt.Println(Db.Password)
		Db.Host = os.Getenv("Host")
		fmt.Println(Db.Host)
		Db.Port = os.Getenv("Port")
		fmt.Println(Db.Port)
		Secret1 = os.Getenv("testis")
		fmt.Println(Secret1)
		Secret2 = os.Getenv("ml")
		fmt.Println(Secret2)
	} else {
		if err := godotenv.Load(".env.prod"); err != nil {
			log.Print("No .env.prod file found")
		}
		log.Print("Find .env.prod ")

		Db.User = os.Getenv("User")
		fmt.Println(Db.User)
		Db.Dbname = os.Getenv("Dbname")
		fmt.Println(Db.Dbname)
		Db.Password = os.Getenv("Password")
		fmt.Println(Db.Password)
		Db.Host = os.Getenv("Host")
		fmt.Println(Db.Host)
		Db.Port = os.Getenv("Port")
		fmt.Println(Db.Port)
		Secret1 = os.Getenv("testis")
		fmt.Println(Secret1)
		Secret2 = os.Getenv("ml")
		fmt.Println(Secret2)
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
