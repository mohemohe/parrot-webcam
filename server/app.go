package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v1 "github.com/mohemohe/parrot-webcam/server/controllers/v1"
	"github.com/mohemohe/parrot-webcam/server/util"
)

func main() {
	util.StartWebcam()

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/api/v1/image", v1.GetImage)
	e.GET("/api/v1/image/taken", v1.GetImageTaken)

	e.Logger.Fatal(e.Start(":1323"))

}