package errors

// InternalServerError a struct for 400 Response ResponseError
type InternalServerError struct {
	name       string
	message    string
	statusCode uint16
}

// NewInternalServerError is a function to create the instance of InternalServerError struct
func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		name:       "Internal Server ResponseError",
		message:    message,
		statusCode: 500,
	}
}

// Name is a method to get the name of the ResponseError
func (b InternalServerError) Name() string {
	return b.name
}

// Message is a method to get the message of the ResponseError
func (b InternalServerError) Error() string {
	return b.message
}

// StatusCode is a method to get the status code of the ResponseError
func (b InternalServerError) StatusCode() uint16 {
	return b.statusCode
}
