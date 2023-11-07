package notification

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init(e *echo.Echo) {
	group := e.Group("/notification")

	group.POST("/email", sendEmailNotification)
}

type sendEmailRequest struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

type sendEmailResponse struct {
	Message string `json:"message"`
}

// sendEmailNotification godoc
// @Summary Send an email notification.
// @Description Send an email notification using the EmailService factory. Expects JSON data in the request body.
// @Tags email
// @Accept json
// @Produce json
// @Param request body sendEmailRequest true "Email request"
// @Success 200 {object} sendEmailResponse
// @Router /notification/email [post]
func sendEmailNotification(c echo.Context) error {
	var request sendEmailRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON request body"})
	}

	emailService := NewEmailNotificationService()
	if err := emailService.SendNotification(request.Subject, request.Body, request.Recipient); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Email sent successfully"})
}
