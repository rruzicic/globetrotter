package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rruzicic/globetrotter/flights/backend/models"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/http"
	"github.com/rruzicic/globetrotter/flights/backend/pkg/jwt"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpGin := http.Gin{Context: ctx}
		err := jwt.TokenValid(ctx, models.UserRole)
		if err != nil {
			ctx.Abort()
			httpGin.Unauthorized(nil)
			return
		}
		ctx.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		httpGin := http.Gin{Context: ctx}
		err := jwt.TokenValid(ctx, models.AdminRole)
		if err != nil {
			ctx.Abort()
			httpGin.Unauthorized(nil)
			return
		}
		ctx.Next()
	}
}
