package main

import (
	"github.com/labstack/echo/v4"
	"go-hackaton/config"
	usercontroller "go-hackaton/internal/app/user/controller"
	userrepo "go-hackaton/internal/app/user/repository/repoimpl"
	userservice "go-hackaton/internal/app/user/service/serviceimpl"
	"go-hackaton/internal/pkg/custom/earn"
)

func main() {

	defer config.App.Close()

	e := echo.New()

	//user usecase
	monetizeService := earn.NewServiceMonetize()
	rulesRepo := userrepo.CreateUserRepoImpl()
	rulesService := userservice.CreateUserServiceImpl(rulesRepo, monetizeService)
	usercontroller.CreateUserController(e, rulesService)

	err := e.Start("")
	if err != nil {
		return
	}

	// Start server
	//go func() {
	//	if err := e.Start(":" + config.App.Config.Port); err != nil && err != http.ErrServerClosed {
	//		e.Logger.Fatal("shutting down the server")
	//	}
	//}()
	//
	//// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	//// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//if err := e.Shutdown(ctx); err != nil {
	//	e.Logger.Fatal(err)
	//}
	//os.Exit(0)
}
