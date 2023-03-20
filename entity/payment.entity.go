package entity

type Payment struct {
	ID          int64  `form:"id" json:"id"`
	Amount      int64  `form:"amount" json:"amount"`
	OrderID     int64  `form:"order_id" json:"order_id"`
	OrderStatus string `form:"order_status" json:"order_status"`
}
