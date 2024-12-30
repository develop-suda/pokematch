package main

import (
	"net/http"
	"pokematch/di"
	"pokematch/infra"
	"pokematch/initializer"
	"pokematch/middlewares"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {
	// アプリケーション設定の初期化
	initializer.Init()
	initializer.Log()

	// インフラ設定の初期化
	infra.Initialize()
	db := infra.SetupDB()

	// データベース接続を使用してルーターを初期化
	pokemonController := di.InitializeRouter(db)

	// 新しいGinルーターを作成
	r := gin.Default()

	// 各リクエストにリクエストIDを生成
	r.Use(requestid.New())
	// カスタムロギングミドルウェアを使用
	r.Use(middlewares.LogMiddleware())

	// HTMLテンプレートを読み込む
	r.LoadHTMLGlob("templates/*")
	// ルートハンドラーを定義
	r.GET("/", pokemonController.Index)
	r.GET("/result", pokemonController.FindPokemon)
	r.GET("/error", pokemonController.Err)

	// 未定義のルートをホームページにリダイレクト
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	// ポート8080でサーバーを起動
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
