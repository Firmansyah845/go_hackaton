package main

import (
	"context"
	"github.com/Firmansyah845/go_hackaton/config"
	usercontroller "github.com/Firmansyah845/go_hackaton/internal/app/user/controller"
	userrepo "github.com/Firmansyah845/go_hackaton/internal/app/user/repository/repoimpl"
	userservice "github.com/Firmansyah845/go_hackaton/internal/app/user/service/serviceimpl"
	"github.com/Firmansyah845/go_hackaton/internal/pkg/custom/forecast"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	defer config.App.Close()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	//user usecase
	monetizeService := forecast.NewServiceMonetize()
	rulesRepo := userrepo.CreateUserRepoImpl()
	rulesService := userservice.CreateUserServiceImpl(rulesRepo, monetizeService)
	usercontroller.CreateUserController(e, rulesService)

	//Start server
	go func() {
		if err := e.Start(":" + config.App.Config.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	os.Exit(0)
}
