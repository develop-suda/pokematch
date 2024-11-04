package controllers

import (
	"fmt"
	"net/http"
	"pokematch/dto"
	"pokematch/services"

	"github.com/gin-gonic/gin"
)

type IPokemonController interface {
	FindPokemon(ctx *gin.Context)
	Index(ctx *gin.Context)
}

type PokemonController struct {
	service services.IPokemonService
}

func (c *PokemonController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "ポケマッチ",
	})
}

func (c *PokemonController) FindPokemon(ctx *gin.Context) {

	// バリデーションチェック
	var input dto.HeightWeightInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		fmt.Println(err)
		return
	}

	// この実装だとパラメータが空でも通ってしまう
	// コメントアウトしているものに変更すると空は通らなくなるが、その代わり普通の０が通らなくなる
	// if input.Height == nil || input.Weight == nil || *input.Height == 0 || *input.Weight == 0 {
	if input.Height == nil || input.Weight == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Height and Weight are required"})
		return
	}

	// ポケモンを取得
	pokemons, err := c.service.FindPokemon(*input.Height, *input.Weight)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.HTML(http.StatusOK, "result.tmpl", gin.H{
		"pokemons": pokemons,
	})
}

func NewPokemonController(service services.IPokemonService) IPokemonController {
	return &PokemonController{service: service}
}
