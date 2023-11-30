package sysinfo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type InfoController interface {
	Init(e *echo.Echo)
	DataBaseHealthCheck(c echo.Context) error
	SaveInfo(c echo.Context) error
	GetInfoFromDb(c echo.Context) error
}

type InfoControllerImpl struct {
	Service InfoService
}

func (controller InfoControllerImpl) Init(e *echo.Echo) {
	group := e.Group("/info")
	group.GET("/database/check", controller.DataBaseHealthCheck)
	group.GET("/save", controller.SaveInfo)
	group.GET("/getFromDb", controller.GetInfoFromDb)
}

// DataBaseHealthCheck godoc
// @Summary Show the status of the database.
// @Description get the status of the database.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /info/database/check [get]
func (controller InfoControllerImpl) DataBaseHealthCheck(c echo.Context) error {
	isOk, error := controller.Service.CheckDb()
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

// SaveInfo godoc
// @Summary It is a test of the database.
// @Description It is a test of the database.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} sysinfo.SysInfo
// @Router /info/save [get]
func (controller InfoControllerImpl) SaveInfo(c echo.Context) error {
	info, error := controller.Service.SaveInfo()
	if error == nil {
		return c.JSON(http.StatusOK, info)
	}
	return c.JSON(http.StatusInternalServerError, error)
}

// GetInfoFromDb godoc
// @Summary It is a test of the database.
// @Description It is a test of the database.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} sysinfo.SysInfo
// @Router /info/getFromDb [get]
func (controller InfoControllerImpl) GetInfoFromDb(c echo.Context) error {
	info, error := controller.Service.GetInfo()
	if error == nil {
		if info == nil {
			return c.JSON(http.StatusNotFound, info)
		} else {
			return c.JSON(http.StatusOK, info)
		}
	}
	return c.JSON(http.StatusInternalServerError, error)
}
