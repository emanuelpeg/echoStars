package main

import (
	"echoStars/computer"
	"echoStars/dataBase"
	_ "echoStars/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

const fileName = "data.db"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health/check", HealthCheck)
	e.GET("/health/info", Info)

	e.GET("/database/health/check", DataBaseHealthCheck)

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

// DataBaseHealthCheck godoc
// @Summary Show the status of the database.
// @Description get the status of the database.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /database/health/check [get]
func DataBaseHealthCheck(c echo.Context) error {
	isOk, error := dataBase.CheckDb(fileName)
	if isOk {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":  "Database is up and running",
			"error": "",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":  "We have a problem!",
			"error": error.Error(),
		})
	}
}

// Info godoc
// @Summary Show the info of server.
// @Description get the info of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} computer.SysInfo
// @Router /health/info [get]
func Info(c echo.Context) error {
	return c.JSON(http.StatusOK, computer.Info())
}
