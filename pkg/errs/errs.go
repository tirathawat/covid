package errs

import "time"

// Error represents the error response.
type Error struct {
	Msg       string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// Error returns the error message.
func (e *Error) Error() string {
	return e.Msg
}

// New creates a new error.
func New(msg string) *Error {
	return &Error{
		Msg:       msg,
		Timestamp: time.Now().Unix(),
	}
}
