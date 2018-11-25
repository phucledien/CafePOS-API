package domain

// Order describe order in system
type Order struct {
	Model
	TableID      UUID          `json:"table_id"`
	TotalPayment float64       `json:"total_payment"`
	OrderDetails []OrderDetail `json:"order_details"`
}
