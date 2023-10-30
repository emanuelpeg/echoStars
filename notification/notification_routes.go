package notification

import (
	"github.com/labstack/echo/v4"
)

func RegisterNotificationRoutes(e *echo.Echo) {
	group := e.Group("/notification")

	group.POST("/email", SendEmailNotification)
}
