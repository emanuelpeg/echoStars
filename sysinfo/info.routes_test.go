package sysinfo

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesSaveInfo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/info/save", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	serviceImpl := NewMockInfoService(mockCtrl)
	var info = getSysInfoExample()
	serviceImpl.EXPECT().SaveInfo().Return(&info, nil)

	infoController := InfoControllerImpl{Service: serviceImpl}
	infoController.Init(e)

	infoJson := getInfoJson()

	// Assertions
	if assert.NoError(t, infoController.SaveInfo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, infoJson, rec.Body.String())
	}
}

func TestRoutesGetInfo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/info/getFromDb", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	serviceImpl := NewMockInfoService(mockCtrl)
	var info = getSysInfoExample()
	serviceImpl.EXPECT().GetInfo().Return(&info, nil)

	infoController := InfoControllerImpl{Service: serviceImpl}
	infoController.Init(e)

	infoJson := getInfoJson()

	// Assertions
	if assert.NoError(t, infoController.GetInfoFromDb(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, infoJson, rec.Body.String())
	}
}

func getInfoJson() string {
	return "{\"hostname\":\"test\",\"platform\":\"testOs\",\"uptime\":200,\"ram\":2000,\"ram available\":1000,\"ram free\":500,\"disk\":100}\n"
}
