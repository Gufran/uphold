package uphold

import (
	"fmt"
	"net/http"
	"time"
)

// ErrorResponse as returned by Uphold
type ErrorResponse struct {
	Response *http.Response
}

// Error returns the string representation of the error
func (e ErrorResponse) Error() string {
	code := e.Response.StatusCode
	return http.StatusText(code)

}

// RateLimitError is returned when API rate limits are exhausted
type RateLimitError struct {
	ErrorResponse
	RequestRate
}

// Error returns the string representation of the error
func (e RateLimitError) Error() string {
	msg := "Rate limit exhausted. Time window will reset on %s, retry after %s"
	return fmt.Sprintf(msg, e.ResetOn, time.Duration(e.RetryAfter)*time.Second)
}
