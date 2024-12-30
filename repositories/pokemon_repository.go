package repositories

import (
	"pokematch/initializer"
	"pokematch/models"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// IPokemonRepository インターフェースの定義
type IPokemonRepository interface {
	// ポケモンを検索するメソッド
	FindPokemon(minHeight float32, maxHeight float32, minWeight float32, maxWeight float32) (*[]models.Pokemon, error)
}

// PokemonRepository 構造体の定義
type PokemonRepository struct {
	db     *gorm.DB
	logger zerolog.Logger
}

// NewPokemonRepository コンストラクタ関数
func NewPokemonRepository(db *gorm.DB) IPokemonRepository {
	// デフォルトのロガーを使用してPokemonRepositoryを初期化
	return &PokemonRepository{db: db, logger: initializer.DefaultLogger()}
}

// FindPokemon ポケモンを検索するメソッドの実装
func (r *PokemonRepository) FindPokemon(minHeight float32, maxHeight float32, minWeight float32, maxWeight float32) (*[]models.Pokemon, error) {
	var pokemons []models.Pokemon
	// 高さと重さの範囲でポケモンを検索
	result := r.db.Where("height BETWEEN ? AND ? and weight BETWEEN ? AND ?", minHeight, maxHeight, minWeight, maxWeight).Limit(10).Find(&pokemons)
	if result.Error != nil {
		// エラーログにリクエストIDとSQL文を含める
		r.logger.Error().Str("request_id", initializer.RequestID).Str("SQL", result.Sql).Send()
		return nil, result.Error
	}
	// 検索結果を返す
	return &pokemons, nil
}
