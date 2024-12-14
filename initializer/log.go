package initializer

import (
	"context"
	"io"
	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Log() {
	// gin.log,go.logファイル作成
	ginf, _ := os.Create("gin.log")
	gof, _ := os.Create("go.log")
	// ginでのログをファイルと標準出力に出力
	gin.DefaultWriter = io.MultiWriter(ginf, os.Stdout)
	// zerologでのログをファイルと標準出力に出力
	log.Logger = log.Output(io.MultiWriter(gof, os.Stdout))
}

func CreateLogger(c *gin.Context) *zerolog.Logger {
	ctx := context.Background()
	// configured-logger
	logger := log.With().
		Str("request_id", requestid.Get(c)).
		Str("ip", c.ClientIP()).
		Str("path", c.Request.URL.Path).
		Str("method", c.Request.Method).
		Str("params", c.Request.URL.RawQuery).
		Logger()

	// context.Contextにロガーを登録
	ctx = logger.WithContext(ctx)

	// context.Contextに設定したロガーを取り出し
	newLogger := zerolog.Ctx(ctx)
	return newLogger
}
