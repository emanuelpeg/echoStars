package main

import (
	_ "echoStars/docs"
	"echoStars/notification"
	"echoStars/server"
	"echoStars/sysinfo"
	"echoStars/util"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Get the profile from environment variable
	activeProfile := os.Getenv("APP_PROFILE")
	if activeProfile == "" {
		log.Fatal("No application profile provided (APP_PROFILE)")
	}
	util.LoadConfig(activeProfile)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health/check", HealthCheck)
	e.GET("/health/info", Info)

	sysinfo.Init(e)
	server.Init(e)
	notification.Init(e)

	//swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health/check [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

// Info godoc
// @Summary Show the info of server.
// @Description get the info of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} sysinfo.SysInfo
// @Router /health/info [get]
func Info(c echo.Context) error {
	return c.JSON(http.StatusOK, sysinfo.Info())
}
