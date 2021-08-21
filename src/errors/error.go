package errors

import "github.com/gofiber/fiber/v2"

// Error an interface to define the Response Error
type Error interface {
	Name() string
	Message() string
	StatusCode() uint16
}

// ErrorHandler is a function to handle the response error
func ErrorHandler(err Error, c *fiber.Ctx) error {
	status := "fail"
	var statusCode uint16

	switch err.(type) {
	case BadRequest:
		statusCode = fiber.StatusBadRequest
	case NotFound:
		statusCode = fiber.StatusNotFound
	default:
		status = "error"
		statusCode = fiber.StatusInternalServerError
	}

	return c.Status(int(statusCode)).JSON(fiber.Map{
		"status":  status,
		"message": err.Message(),
	})
}
