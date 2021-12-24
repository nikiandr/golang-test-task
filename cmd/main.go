package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nikiandr/golang-test-task/pkg/handler"
	"github.com/nikiandr/golang-test-task/pkg/repository"
	"github.com/nikiandr/golang-test-task/pkg/server"
	"github.com/nikiandr/golang-test-task/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {
	//setting up logger JSON format
	logrus.SetFormatter(new(logrus.JSONFormatter))

	var (
		dbAuthStr string
		db        *sqlx.DB
		err       error
	)

	if dbAuthStr = os.Getenv("DATABASE_URL"); dbAuthStr == "" {
		logrus.Fatalf("Error getting database authentication string: %s", err.Error())
	} else {
		db, err = repository.NewPostgresDBAuthString(dbAuthStr)
	}

	if err != nil {
		logrus.Fatalf("Error occured while connecting to DB server: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("ClubApp started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGTTIN)
	<-quit
	logrus.Print("ClubApp shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
