package ots

import "fmt"

// Error API error.
type Error struct {
	Message    string
	StatusCode int
}

// Error returns error as a string.
func (e Error) Error() string {
	var msg string
	if e.Message != "" {
		msg = fmt.Sprintf(": %q", e.Message)
	}

	return fmt.Sprintf("api error%s (status code %d)", msg, e.StatusCode)
}
