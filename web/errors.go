package web

// validation error
type ValidationError struct {
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

// serialization error
type SerializationError struct {
	Message string `json:"message"`
}

func (e *SerializationError) Error() string {
	return e.Message
}

// messages
type NotFound struct {
	Message string `json:"-"`
}

func (e *NotFound) Error() string {
	return e.Message
}

type NoContent struct {
	Message string `json:"-"`
}

func (e *NoContent) Error() string {
	return e.Message
}

type Forbidden struct {
	Message string `json:"-"`
}

func (e *Forbidden) Error() string {
	return "Forbidden"
}

type BadRequest struct {
	Message string `json:"message"`
}

func (e *BadRequest) Error() string {
	return e.Message
}
