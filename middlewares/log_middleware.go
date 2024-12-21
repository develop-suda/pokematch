package middlewares

import (
	"pokematch/initializer"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		initializer.SetRequestID(c)
		logger := initializer.MiddlewareLogger(c)
		logger.Info().Msg("REQUEST START")
		c.Next()
		logger.Info().Msg("REQUEST END")
	}
}
