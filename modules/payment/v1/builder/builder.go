package builder

import (
	"gin-starter/app"
	"gin-starter/config"
	paymentRepo "gin-starter/modules/payment/v1/repository"
	"gin-starter/modules/payment/v1/service"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// BuildMasterHandler builds master handler
// starting from handler down to repository or tool.
func BuildPaymentHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	pr := paymentRepo.NewPaymentRepository(db)

	pf := service.NewPaymentFinder(cfg, pr)

	app.PaymentFinderHTTPHandler(cfg, router, pf)
}
