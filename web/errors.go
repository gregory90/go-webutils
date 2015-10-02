package web

// validation error
type ValidationError struct {
	ErrorType string      `json:"error,omitempty"`
	Fields    interface{} `json:"fields,omitempty"`
}

func (e *ValidationError) Error() string {
	return e.ErrorType
}

// serialization error
type SerializationError struct {
	ErrorType string `json:"error"`
}

func (e *SerializationError) Error() string {
	return e.ErrorType
}

// errors
type NotFound struct {
	ErrorType string `json:"-"`
}

func (e *NotFound) Error() string {
	return e.ErrorType
}

type NoContent struct {
	ErrorType string `json:"-"`
}

func (e *NoContent) Error() string {
	return e.ErrorType
}

type Forbidden struct {
	ErrorType string `json:"error"`
}

func (e *Forbidden) Error() string {
	return e.ErrorType
}

type BadRequest struct {
	ErrorType string `json:"error"`
}

func (e *BadRequest) Error() string {
	return e.ErrorType
}
