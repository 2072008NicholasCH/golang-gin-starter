package handler

import (
	"gin-starter/common/errors"
	"gin-starter/modules/payment/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Payment finder handler

type PaymentFinderHandler struct {
	paymentFinder service.PaymentFinderUseCase
}

// NewPaymentFinderHandler

func NewPaymentFinderHandler(paymentFinder service.PaymentFinderUseCase) *PaymentFinderHandler {
	return &PaymentFinderHandler{
		paymentFinder: paymentFinder,
	}
}

// GetPayments

func (pf *PaymentFinderHandler) GetPayments(c *gin.Context) {
	payments, err := pf.paymentFinder.GetPayments(c.Request.Context())
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, c.Err().Error()))
		c.Abort()
		return
	}

	res := make([]*resource.Payment, 0)

	for _, payment := range payments {
		res = append(res, resource.NewPaymentResponse(payment))
	}
	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", &resource.PaymentListResponse{
		List:  res,
		Total: int64(len(res)),
	}))
}

func (pf *PaymentFinderHandler) GetPaymentByID(c *gin.Context) {
	var req resource.GetPaymentByIDRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	payment, err := pf.paymentFinder.GetPaymentByID(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(errors.ErrInternalServerError.Code, response.ErrorAPIResponse(errors.ErrInternalServerError.Code, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewPaymentResponse(payment)))
}
