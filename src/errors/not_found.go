package errors

// NotFound a struct for 404 Response ResponseError
type NotFound struct {
	name       string
	message    string
	statusCode uint16
}

// NewNotFound is a function to create the instance of NotFound struct
func NewNotFound(message string) NotFound {
	return NotFound{
		name:       "Not Found",
		message:    message,
		statusCode: 404,
	}
}

// Name is a method to get the name of the ResponseError
func (b NotFound) Name() string {
	return b.name
}

// Message is a method to get the message of the ResponseError
func (b NotFound) Error() string {
	return b.message
}

// StatusCode is a method to get the status code of the ResponseError
func (b NotFound) StatusCode() uint16 {
	return b.statusCode
}
