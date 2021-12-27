package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nikiandr/golang-test-task/pkg/handler"
	"github.com/nikiandr/golang-test-task/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := new(server.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
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
}
