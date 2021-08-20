package errors

// BadRequest a struct for 400 Response Error
type BadRequest struct {
	name       string
	message    string
	statusCode uint16
}

// NewBadRequest is a function to create the instance of BadRequest struct
func NewBadRequest(message string) BadRequest {
	return BadRequest{
		name:       "Bad Request",
		message:    message,
		statusCode: 400,
	}
}

// Name is a method to get the name of the Error
func (b BadRequest) Name() string {
	return b.name
}

// Message is a method to get the message of the Error
func (b BadRequest) Message() string {
	return b.message
}

// StatusCode is a method to get the status code of the Error
func (b BadRequest) StatusCode() uint16 {
	return b.statusCode
}
