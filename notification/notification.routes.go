package notification

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init(e *echo.Echo) {
	group := e.Group("/notification")

	group.POST("/email", sendEmailNotification)
	group.POST("/sms", sendSMSNotification)
}

type sendEmailRequest struct {
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Recipient string `json:"recipient"`
}

type sendEmailResponse struct {
	Message string `json:"message"`
}

type sendSMSRequest struct {
	Message     string `json:"message"`
	PhoneNumber string `json:"phoneNumber"`
}

// sendEmailNotification godoc
// @Summary Send an email notification.
// @Description Send an email notification using the EmailService factory. Expects JSON data in the request body.
// @Tags notification
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

// sendSMSNotification godoc
// @Summary Send an SMS notification.
// @Description Using short message service from Twilio, sends a test notification to any number. Phone number in international format, example:"+123456789"
// @Tags notification
// @Accept json
// @Produce json
// @Param request body sendSMSRequest true "SMS request"
// @Success 200 {object} sendEmailResponse
// @Router /notification/sms [post]
func sendSMSNotification(c echo.Context) error {
	var request sendSMSRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON request body"})
	}

	smsService := NewSMSNotificationService()
	if err := smsService.SendNotification("", request.Message, request.PhoneNumber); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "SMS sent successfully"})
}
