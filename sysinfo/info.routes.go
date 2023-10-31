package sysinfo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init(e *echo.Echo) {
	group := e.Group("/info")
	group.GET("/database/check", DataBaseHealthCheck)
	group.GET("/save", SaveInfo)
	group.GET("/getFromDb", GetInfoFromDb)
}

// DataBaseHealthCheck godoc
// @Summary Show the status of the database.
// @Description get the status of the database.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /info/database/check [get]
func DataBaseHealthCheck(c echo.Context) error {
	service := NewInfoService()
	isOk, error := service.CheckDb()
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
func SaveInfo(c echo.Context) error {
	var service = NewInfoService()
	var info, error = service.SaveInfo()
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
func GetInfoFromDb(c echo.Context) error {
	var service = NewInfoService()
	var info, error = service.GetInfo()
	if error == nil {
		if info == nil {
			return c.JSON(http.StatusNotFound, info)
		} else {
			return c.JSON(http.StatusOK, info)
		}
	}
	return c.JSON(http.StatusInternalServerError, error)
}
