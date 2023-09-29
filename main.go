package main

import (
	"echoStars/computer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/health/check", check)
	e.GET("/health/info", info)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func check(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func info(c echo.Context) error {
	return c.JSON(http.StatusOK, computer.Info())
}
