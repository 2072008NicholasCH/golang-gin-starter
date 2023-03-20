package resource

import "gin-starter/entity"

type GetPaymentByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type Payment struct {
	ID          int64  `form:"id" json:"id"`
	Amount      int64  `form:"amount" json:"amount"`
	OrderID     int64  `form:"order_id" json:"order_id"`
	OrderStatus string `form:"order_status" json:"order_status"`
}

func NewPaymentResponse(payment *entity.Payment) *Payment {
	return &Payment{
		ID:          payment.ID,
		Amount:      payment.Amount,
		OrderID:     payment.OrderID,
		OrderStatus: payment.OrderStatus,
	}
}

type PaymentListResponse struct {
	List  []*Payment `json:"list"`
	Total int64      `json:"total"`
}
