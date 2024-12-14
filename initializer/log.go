package initializer

import (
	"context"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Log() {
	// gin.logファイル作成
	f, _ := os.Create("gin.log")
	// gin.logファイルと標準出力にログを出力
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.Logger = log.Output(gin.DefaultWriter)
}

func CreateLogger(c *gin.Context) *zerolog.Logger {
	ctx := context.Background()
	// configured-logger
	logger := log.With().
		Str("path", c.Request.URL.Path).
		Str("method", c.Request.Method).
		Logger()

	// context.Contextにロガーを登録
	ctx = logger.WithContext(ctx)

	// context.Contextに設定したロガーを取り出し
	newLogger := zerolog.Ctx(ctx)
	return newLogger
}
