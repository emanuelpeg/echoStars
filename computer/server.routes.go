package computer

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo) {
	group := e.Group("/servers")
	group.GET("", getServers)
	group.POST("", addServer)
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

func addServer(c echo.Context) error {
	servers := TestServers
	err := AddServer(servers[0])
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("{\"status\": %d}", http.StatusOK))
}
