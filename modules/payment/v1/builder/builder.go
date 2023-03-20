package builder

import (
	"gin-starter/config"
	"gin-starter/modules/auth/v1/service"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// BuildMasterHandler builds master handler
// starting from handler down to repository or tool.
func BuildPaymentHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
	pr := paymentRepo.NewPaymentRepository(db)

	pf := service.NewPaymentFinderService(cfg, pr)

	app.PaymentHandler(cfg, router, pf)
}
