package initializer

import (
	"context"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// TODO: 各処理のログを出力する
func Log() {
	// gin.log,go.logファイル作成
	ginf, err := getLogFile("gin")
	if err != nil {
		log.Error().Msg("error")
	}

	zerologf, err := getLogFile("zerolog")
	if err != nil {
		log.Error().Msg("error")
	}
	// ginでのログをファイルと標準出力に出力
	gin.DefaultWriter = io.MultiWriter(ginf, os.Stdout)
	// zerologでのログをファイルと標準出力に出力
	log.Logger = log.Output(io.MultiWriter(zerologf, os.Stdout))
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

func getLogFile(fileName string) (*os.File, error) {
	t := time.Now()

	y := strconv.Itoa(t.Year())
	m := strconv.Itoa(int(t.Month()))
	d := strconv.Itoa(t.Day())

	yyyymmdd := y + m + d

	filePath := "logs/" + fileName + "/" + yyyymmdd + "_" + fileName + ".log"

	// ファイルが存在しない場合　err != nil
	// ファイルが存在する場合　　　　err = nil
	_, err := os.Stat(filePath)
	if err != nil {
		return createLogFile(filePath)
	} else {
		return openLogFile(filePath)
	}
}

func createLogFile(filePath string) (*os.File, error) {
	return os.Create(filePath)
}

func openLogFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
