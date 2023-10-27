package computer

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo) {
	group := e.Group("/servers")
	group.GET("", getServers)
	group.POST("", createServer)
	group.DELETE("/:hostname", deleteServer)
}

func getServers(c echo.Context) error {
	servers, err := GetServers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, servers)
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

func deleteServer(c echo.Context) error {
	hostname := c.Param("hostname")

	// TODO add find by 'id' in order to check if the server exists before delete.
	server := Server{}
	if err := c.Bind(&server); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := DeleteServer(&hostname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted server",
	})
}
