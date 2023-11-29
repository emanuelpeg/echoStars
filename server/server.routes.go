package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	group := e.Group("/servers")
	group.POST("", createServer)
	group.GET("", getServers)
	group.DELETE("/:url", deleteServer)
}

// CreateServer godoc
// @Summary Creates a new server.
// @Description Creates new server for monitoring using a server struct as input.
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} server.Server
// @Router /servers [post]
func createServer(c echo.Context) error {
	// define empty server
	var server = Server{}
	// link the received 'server' from the context and check the structure is ok.
	if err := c.Bind(&server); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// TODO check if this could be pass as 'class' variable
	var service, error = NewServerService()
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error)
	}

	// call the (Abstraction) interface in the service
	var createdServer, err = service.Upsert(&server)
	// if the result has an err, return the server error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	// no errors, but empty content, maybe this could be deleted
	if createdServer == nil {
		return c.JSON(http.StatusNotFound, server)
	}
	// no errors, return the content
	return c.JSON(http.StatusOK, server)

}

// GetServers godoc
// @Summary Lists registered servers.
// @Description List all registered servers.
// @Tags server
// @Accept */*
// @Produce json
// @Success 200 {object} []server.Server
// @Router /servers [get]
func getServers(c echo.Context) error {
	var service, error = NewServerService()
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error)
	}

	servers, serviceError := service.GetAll()
	if serviceError != nil {
		return c.JSON(http.StatusBadRequest, serviceError)
	}
	return c.JSON(http.StatusOK, servers)
}

// DeleteServer godoc
// @Summary delete a server.
// @Description delete  server from server id.
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {string} hostname
// @Router /servers [delete]
func deleteServer(c echo.Context) error {
	url := c.Param("url")

	var service, error = NewServerService()
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error)
	}

	if err := c.Bind(&Server{}); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	isDeleted, err := service.Delete(&url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	if isDeleted {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Server was deleted successfully",
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"message": "the server could not be deleted",
	})
}
