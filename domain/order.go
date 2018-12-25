package domain

// Order describe order in system
type Order struct {
	Model
	TableID      UUID          `json:"table_id"`
	TotalPayment int           `json:"total_payment"`
	OrderDetails []OrderDetail `json:"order_details"`
}
