package error

// Error an interface to define the Response Error
type Error interface {
	Name() string
	Message() string
	StatusCode() uint16
}
