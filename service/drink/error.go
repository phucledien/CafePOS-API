package drink

import (
	"net/http"
)

// ErrNotFound Declaration
var ErrNotFound = errNotFound{}

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}
