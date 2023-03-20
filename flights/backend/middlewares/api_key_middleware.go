package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/services"
)

func APIKeyAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpGin := http.Gin{Context: ctx}
		api_key := httpGin.Context.Request.Header.Get("X-API-Key")

		user, err := services.FindUserByAPIKey(api_key)

		if err != nil {
			log.Println("Unknown api key")
			ctx.Abort()
			httpGin.Unauthorized(nil)
			return
		}

		if !services.CheckAPIKeyExpiration(user.APIKey) {
			log.Println("Expired api key")
			ctx.Abort()
			httpGin.Unauthorized(nil)
			return
		}

		ctx.Next()
	}
}
