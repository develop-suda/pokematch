package controllers

import (
	"net/http"
	"pokematch/dto"
	"pokematch/initializer"
	"pokematch/services"

	"github.com/rs/zerolog"

	"github.com/gin-gonic/gin"
)

// IPokemonControllerインターフェースの定義
type IPokemonController interface {
	FindPokemon(ctx *gin.Context)
	Index(ctx *gin.Context)
	Err(ctx *gin.Context)
}

// PokemonController構造体の定義
type PokemonController struct {
	service services.IPokemonService
	logger  zerolog.Logger
}

// Indexメソッドの実装
func (c *PokemonController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "ポケマッチ",
	})
}

// FindPokemonメソッドの実装
func (c *PokemonController) FindPokemon(ctx *gin.Context) {

	// ログにリクエストIDを記録
	c.logger.Info().Str("request_id", initializer.RequestID).Msg("FindPokemon実行")

	// バリデーションチェック
	var input dto.HeightWeightInput
	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": "身長or体重のどちらかは空もしくは数字以外が入っています!!!"})
		return
	}

	// 入力値のチェック
	if input.Height == nil || input.Weight == nil || *input.Height == 0 || *input.Weight == 0 {
		var errMsg string
		var resMsg string
		switch {
		case (input.Height == nil || *input.Height == 0) && (input.Weight == nil || *input.Weight == 0):
			errMsg = "身長体重どちらとも0もしくは空"
			resMsg = "身長と体重のどちらとも0もしくは空です。正しい値を入力してください。"
		case input.Height == nil || *input.Height == 0:
			errMsg = "身長が0もしくは空"
			resMsg = "身長が0もしくは空です。正しい値を入力してください。"
		case input.Weight == nil || *input.Weight == 0:
			errMsg = "体重が0もしくは空"
			resMsg = "体重が0もしくは空です。正しい値を入力してください。"
		}
		// エラーログの記録
		c.logger.Error().Float32("Height", *input.Height).Float32("Weight", *input.Weight).Msg(errMsg)
		ctx.HTML(http.StatusBadRequest, "error.tmpl", gin.H{"error": resMsg})
		return
	}

	// ポケモンを取得
	pokemons, err := c.service.FindPokemon(*input.Height, *input.Weight)
	if err != nil {
		// エラーログの記録
		c.logger.Error().Err(err).Str("request_id", initializer.RequestID).Send()
		ctx.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "内部でエラーが発生しました。"})
		return
	}

	// 結果をHTMLで返す
	ctx.HTML(http.StatusOK, "result.tmpl", gin.H{
		"pokemons": pokemons,
	})

}

// Errメソッドの実装
func (c *PokemonController) Err(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
		"error": "エラー用画面テスト表示",
	})
}

// NewPokemonControllerファクトリーメソッドの実装
func NewPokemonController(service services.IPokemonService) IPokemonController {
	// ロガーの初期化
	logger := initializer.DefaultLogger()
	return &PokemonController{service: service, logger: logger}
}
