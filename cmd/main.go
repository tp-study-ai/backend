package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/tp-study-ai/backend/tools"
	"log"
	"net/http"
)

func main() {
	pgxManager, err := db.NewPostgresqlX()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating postgres agent"))
	}
	defer pgxManager.Close()

	router := echo.New()

	httpServ := http.Server{
		Addr:         "127.0.0.1",
		//ReadTimeout:  time.Duration(config.ServConfig.ReadTimeout) * time.Second,
		//WriteTimeout: time.Duration(config.ServConfig.WriteTimeout) * time.Second,
		Handler:      router,
	}

	if err := router.StartServer(&httpServ); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
