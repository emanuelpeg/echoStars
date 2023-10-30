package notification

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SendEmailRequest struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

type SendEmailResponse struct {
	Message string `json:"message"`
}

// SendEmailNotification godoc
// @Summary Send an email notification.
// @Description Send an email notification using the EmailNotification factory. Expects JSON data in the request body.
// @Tags email
// @Accept json
// @Produce json
// @Param request body SendEmailRequest true "Email request"
// @Success 200 {object} SendEmailResponse
// @Router /notification/email [post]
func SendEmailNotification(c echo.Context) error {
	var request SendEmailRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON request body"})
	}

	emailFactory := NewEmailNotificationFactory()
	if err := emailFactory.SendNotification(request.Subject, request.Body, request.Recipient); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Email sent successfully"})
}
