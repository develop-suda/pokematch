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
	initializer.Init()
	initializer.Log()

	infra.Initialize()
	db := infra.SetupDB()

	pokemonController := di.InitializeRouter(db)

	r := gin.Default()

	// リクエストIDの生成
	r.Use(requestid.New())
	r.Use(middlewares.LogMiddleware())

	r.LoadHTMLGlob("templates/*")
	r.GET("/", pokemonController.Index)
	r.GET("/result", pokemonController.FindPokemon)
	r.GET("/error", pokemonController.Err)

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
