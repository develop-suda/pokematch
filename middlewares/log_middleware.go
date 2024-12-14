package middlewares

import (
	"pokematch/initializer"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := initializer.CreateLogger(c)
		logger.Info().Msg("request start")
		c.Next()
		logger.Info().Msg("request end")
	}
}
