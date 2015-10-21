package web

// validation error
type ValidationError struct {
	ErrorType string      `json:"error,omitempty"`
	Fields    interface{} `json:"fields,omitempty"`
}

func (e *ValidationError) Error() string {
	return e.ErrorType
}

func NewValidationError(fields interface{}) *ValidationError {
	return &ValidationError{ErrorType: "validation_error", Fields: fields}
}

// serialization error
type SerializationError struct {
	ErrorType string `json:"error"`
}

func (e *SerializationError) Error() string {
	return e.ErrorType
}

func NewSerializationError() *SerializationError {
	return &SerializationError{ErrorType: "serialization_error"}
}

func NewDeserializationError() *SerializationError {
	return &SerializationError{ErrorType: "deserialization_error"}
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
