package domain

// Drink describe drink in system
type Drink struct {
	Model
	Name  string `json:"name"`
	Price int    `json:"price"`
}
