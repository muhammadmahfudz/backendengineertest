package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Backend_Engineer_Interview_Assignment/common/config"
	psgr "Backend_Engineer_Interview_Assignment/common/database/postgresql"
	"Backend_Engineer_Interview_Assignment/handler/users"
	repository "Backend_Engineer_Interview_Assignment/repository/users"

	"github.com/labstack/echo/v4"
)

func verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Validation RS256
		return next(c)
	}
}

func Run(cfg *config.Config) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*60, "the duration for which the server gracefully wait for existing connections to finish - e.g. 1m")
	flag.Parse()

	postgr := psgr.PostgreSQLConfig(cfg)

	e := echo.New()

	repositoryUsers := repository.NewRepositoryUsers(postgr)
	handlerUsers := users.NewUsers(repositoryUsers)

	e.POST("/user/registration", handlerUsers.Registration)
	e.PUT("/user/update", verify(handlerUsers.Update))
	e.GET("user/profile/:id", verify(handlerUsers.Profile))
	e.POST("user/login", handlerUsers.Login)

	srv := &http.Server{
		Addr:         cfg.Http.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      e,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
