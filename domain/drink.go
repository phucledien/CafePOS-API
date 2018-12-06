package domain

// Drink describe drink in system
type Drink struct {
	Model
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
