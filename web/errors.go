package web

// validation error
type ValidationError struct {
	Error  string      `json:"error,omitempty"`
	Fields interface{} `json:"fields,omitempty"`
}

func (e *ValidationError) Error() string {
	return e.Error
}

// serialization error
type SerializationError struct {
	Error string `json:"error"`
}

func (e *SerializationError) Error() string {
	return e.Error
}

// errors
type NotFound struct {
	Error string `json:"-"`
}

func (e *NotFound) Error() string {
	return e.Error
}

type NoContent struct {
	Error string `json:"-"`
}

func (e *NoContent) Error() string {
	return e.Error
}

type Forbidden struct {
	Error string `json:"error"`
}

func (e *Forbidden) Error() string {
	return e.Error
}

type BadRequest struct {
	Error string `json:"error"`
}

func (e *BadRequest) Error() string {
	return e.Error
}
