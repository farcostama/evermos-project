package utils

import "github.com/gofiber/fiber/v2"

// Struktur standar response API
type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c *fiber.Ctx, statusCode int, message string, data interface{}, errors interface{}) error {
	status := true
	if statusCode >= 400 {
		status = false
	}

	return c.Status(statusCode).JSON(ApiResponse{
		Status:  status,
		Message: message,
		Errors:  errors,
		Data:    data,
	})
}