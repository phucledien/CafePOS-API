package domain

// Table describe table in system
type Table struct {
	Model
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
