package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func extractKey(ctx *gin.Context) string {
	return ctx.Request.Header.Get("x-api-key")
}

func APIKeyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpGin := http.Gin{Context: ctx}
		key := extractKey(ctx)

		user, err := services.FindUserByAPIKey(models.APIKey{Key: key, Expiration: time.Now()})
		if err != nil {
			ctx.Abort()
			httpGin.Unauthorized("No user with this key")
			return
		}

		if services.APIKeyExpired(user.ApiKey) {
			ctx.Abort()
			httpGin.Unauthorized("Key expired")
			return
		}

		ctx.Next()
	}
}
