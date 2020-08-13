package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/mohemohe/parrot-webcam/server/controllers/v1"
	"github.com/mohemohe/parrot-webcam/server/middlewares"
	"github.com/mohemohe/parrot-webcam/server/models"
	"github.com/mohemohe/parrot-webcam/server/util"
	"os"
)

func main() {
	util.StartWebcam()

	rootID := os.Getenv("ROOT_ID")
	rootPassword := os.Getenv("ROOT_PASSWORD")
	models.UpsertUser(&models.User{
		Name:         "root",
		Email:        rootID,
		Password:     rootPassword,
		Role:         0,
	})

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "./assets/",
		HTML5: true,
	}))

	e.GET("/api/v1/auth", v1.AuthStatus, middlewares.Authorize, middlewares.Authorized)
	e.POST("/api/v1/auth", v1.Auth)
	e.GET("/api/v1/image", v1.GetImage, middlewares.Authorize, middlewares.Authorized)
	e.GET("/api/v1/image/taken", v1.GetImageTaken, middlewares.Authorize, middlewares.Authorized)

	e.Logger.Fatal(e.Start(":1323"))

}