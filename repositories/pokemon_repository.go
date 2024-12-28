package repositories

import (
	"pokematch/initializer"
	"pokematch/models"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type IPokemonRepository interface {
	FindPokemon(minHeight float32, maxHeight float32, minWeight float32, maxWeight float32) (*[]models.Pokemon, error)
}

type PokemonRepository struct {
	db     *gorm.DB
	logger zerolog.Logger
}

func NewPokemonRepository(db *gorm.DB) IPokemonRepository {
	return &PokemonRepository{db: db, logger: initializer.DefaultLogger()}
}

func (r *PokemonRepository) FindPokemon(minHeight float32, maxHeight float32, minWeight float32, maxWeight float32) (*[]models.Pokemon, error) {
	var pokemons []models.Pokemon
	result := r.db.Where("height BETWEEN ? AND ? and weight BETWEEN ? AND ?", minHeight, maxHeight, minWeight, maxWeight).Limit(10).Find(&pokemons)
	if result.Error != nil {
		r.logger.Error().Str("request_id", initializer.RequestID).Str("SQL", result.Sql).Send()
		return nil, result.Error
	}
	return &pokemons, nil
}
