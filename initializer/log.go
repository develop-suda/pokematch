package initializer

import (
	"context"
	"errors"
	"os"
	"pokematch/public"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// TODO: 各処理のログを出力する
// TODO: 日付を跨いで実行している場合、前日のファイルに書き込みをしてしまう
func Log() {
	// gin.log,go.logファイル作成
	// ginf, err := getLogFile("gin", "log")
	// if err != nil {
	// 	log.Error().Msg("error")
	// }

	// zerologf, err := getLogFile("zerolog", "log")
	// if err != nil {
	// 	log.Error().Msg("error")
	// }

	// zerologjsonf, err := getLogFile("zerolog", "json")
	// if err != nil {
	// 	log.Error().Msg("error")
	// }
	// ginでのログをファイルと標準出力に出力
	// gin.DefaultWriter = io.MultiWriter(ginf, os.Stdout)
	// zerologでのログをファイルと標準出力に出力
	// log.Logger = log.Output(io.MultiWriter(zerologf, zerologjsonf))

	logstudy()
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

func getLogFile(fileName string, extension string) (*os.File, error) {
	yyyymmdd := public.FormatDate()

	filePath := "logs/" + fileName + "/" + yyyymmdd + "_" + fileName + "." + extension

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

func logstudy() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Trace().Send()
	log.Debug().Send()
	log.Info().Send()
	log.Warn().Send()
	log.Error().Send()
	// log.Fatal().Send()
	// log.Panic().Send()
	log.Info().Send()

	err := outer()
	// なんかスタックトレースでねえ
	log.Error().Stack().Err(err).Msg("")

	f()

}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}

// こいつは使えそうや
func f() {
	logger := zerolog.New(os.Stdout)
	ctx := context.Background()

	// Attach the Logger to the context.Context
	ctx = logger.WithContext(ctx)
	someFunc(ctx)
}

func someFunc(ctx context.Context) {
	// Get Logger from the go Context. if it's nil, then
	// `zerolog.DefaultContextLogger` is returned, if
	// `DefaultContextLogger` is nil, then a disabled logger is returned.
	logger := zerolog.Ctx(ctx)
	logger.Info().Msg("Hello")
}
