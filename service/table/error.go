package table

import (
	"net/http"
)

// Error Declaration
var (
	ErrNotFound = errNotFound{}
)

type errNotFound struct{}

func (errNotFound) Error() string {
	return "record not found"
}
func (errNotFound) StatusCode() int {
	return http.StatusNotFound
}
