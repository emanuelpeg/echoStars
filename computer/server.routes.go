package computer

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo) {
	group := e.Group("/servers")
	group.GET("", getServers)
	group.POST("", createServer)
}

func getServers(c echo.Context) error {
	servers, err := GetServers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"servers": servers,
	})
}

func createServer(c echo.Context) error {
	server := Server{}
	if err := c.Bind(&server); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := CreateServer(server)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "not accepted")
	}

	return c.JSON(http.StatusOK, server)
}
