package server

import (
	"log"
	"os"

	"github.com/k-ueki/sublog/app/server/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	if err := controller.FetchBlogs(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := rubServe(); err != nil {
		log.Fatal(err)
	}
}

func rubServe() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	setRoute(e)

	return e.Start(":5000")
}

func setRoute(e *echo.Echo) {
	//e.GET("", controller.List)

	//api.POST("/emergency", controller.StopAll)

}
