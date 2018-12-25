package domain

// OrderDetail describe order detail in system
type OrderDetail struct {
	Model
	OrderID UUID `json:"order_id"`

	Drink   Drink `json:"drink"`
	DrinkID UUID  `json:"drink_id"`

	Quantity int `json:"quantity"`
}
