package controllers

import (
	"net/http"
	"pokematch/dto"
	"pokematch/services"

	"github.com/gin-gonic/gin"
)

type IPokemonController interface {
	FindPokemon(ctx *gin.Context)
	Index(ctx *gin.Context)
	Err(ctx *gin.Context)
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
	// 空が入っていても何故か通る。。。
	var input dto.HeightWeightInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": "身長or体重のどちらかは空もしくは数字以外が入っています!!!"})
		return
	}

	// この実装だとパラメータが空でも通ってしまう
	// コメントアウトしているものに変更すると空は通らなくなるが、その代わり普通の０が通らなくなる
	// if input.Height == nil || input.Weight == nil || *input.Height == 0 || *input.Weight == 0 {
	if input.Height == nil || input.Weight == nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": "身長と体重は必須です"})
		return
	}

	// ポケモンを取得
	pokemons, err := c.service.FindPokemon(*input.Height, *input.Weight)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "内部でエラーが発生しました。"})
		return
	}

	ctx.HTML(http.StatusOK, "result.tmpl", gin.H{
		"pokemons": pokemons,
	})
}

func (c *PokemonController) Err(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
		"error": "エラー用画面テスト表示",
	})
}

func NewPokemonController(service services.IPokemonService) IPokemonController {
	return &PokemonController{service: service}
}
