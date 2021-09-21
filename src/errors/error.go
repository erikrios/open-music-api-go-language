package errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// ResponseError an interface to define the Response Error
type ResponseError interface {
	Name() string
	Error() string
	StatusCode() uint16
}

// ErrorHandler is a function to handle the response error
func ErrorHandler(err ResponseError, c *fiber.Ctx) error {
	status := "fail"
	var statusCode uint16

	switch err.(type) {
	case BadRequest:
		statusCode = fiber.StatusBadRequest
	case NotFound:
		statusCode = fiber.StatusNotFound
	default:
		fmt.Println(err.Error())
		status = "error"
		statusCode = fiber.StatusInternalServerError
	}

	return c.Status(int(statusCode)).JSON(fiber.Map{
		"status":  status,
		"message": err.Error(),
	})
}
