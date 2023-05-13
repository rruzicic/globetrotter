package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/bnb/account-service/models"
)

func AnyUserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !ValidateRequestToken(ctx, models.GuestRole) && !ValidateRequestToken(ctx, models.HostRole) {
			ctx.JSON(401, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
