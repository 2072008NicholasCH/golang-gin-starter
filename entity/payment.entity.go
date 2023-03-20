package entity

const (
	// paymentTableName is a variable for user table name
	paymentTableName = "main.payments"
)

type Payment struct {
	ID          int64  `form:"id" json:"id"`
	Amount      int64  `form:"amount" json:"amount"`
	OrderID     int64  `form:"order_id" json:"order_id"`
	OrderStatus string `form:"order_status" json:"order_status"`
}

// TableName represents table name on db, need to define it because the db has multi schema
func (u *Payment) TableName() string {
	return paymentTableName
}
