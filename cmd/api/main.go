// Entry point for application's API
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vbetsun/space-trouble/configs"
	"github.com/vbetsun/space-trouble/internal/service"
	"github.com/vbetsun/space-trouble/internal/storage/psql"
	"github.com/vbetsun/space-trouble/internal/transport/rest"
	"github.com/vbetsun/space-trouble/internal/transport/rest/handler"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	conf, err := configs.LoadConfig("configs")
	if err != nil {
		logger.Fatal(fmt.Sprintf("can't read config: %v", err))
	}

	db, err := psql.NewDB(psql.Config{
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		Username: conf.DB.Username,
		DBName:   conf.DB.DBName,
		Password: conf.DB.Password,
		SSLMode:  conf.DB.SSLMode,
		Logger:   logger,
	})
	if err != nil {
		logger.Fatal(fmt.Sprintf("can't connect to the DB %v", err))
	}
	store := psql.NewStorage(db)
	services := service.NewService(service.Deps{
		LaunchpadStorage: store.Launchpad,
		OrderStorage:     store.Order,
		TripStorage:      store.Trip,
		UserStorage:      store.User,
	})
	h := handler.New(handler.Deps{
		Services: handler.Services{
			Launchpad: services.Launchpad,
			Order:     services.Order,
			Trip:      services.Trip,
			User:      services.User,
		},
		Log: logger,
	})
	srv := new(rest.Server)
	go func() {
		if err := srv.Run(conf.Port, h.Routes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Sprintf("can't start server on port %s, err: %v", conf.Port, err))
		} else {
			logger.Info("Server stopped gracefully")
		}
	}()
	logger.Info("Server is starting on port: " + conf.Port)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-exit
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error("Error occurred while server is shutting down " + err.Error())
	}
	if err := db.Close(); err != nil {
		logger.Error("Error occurred while db is closing " + err.Error())
	}
}
