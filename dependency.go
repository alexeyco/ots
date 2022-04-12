package ots

import "net/http"

// HTTPClient interface.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}
