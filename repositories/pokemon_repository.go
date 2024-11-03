package repositories

import (
	"errors"
	"pokematch/models"

	"gorm.io/gorm"
)

type IPokemonRepository interface {
	FindPokemon(minHeight float32, maxHeight float32, minWeight float32, maxWeight float32) (*[]models.Pokemon, error)
}

type PokemonRepository struct {
	db *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) IPokemonRepository {
	return &PokemonRepository{db: db}
}

func (r *PokemonRepository) FindPokemon(minHeight float32, maxHeight float32, minWeight float32, maxWeight float32) (*[]models.Pokemon, error) {
	var pokemons []models.Pokemon
	result := r.db.Where("height BETWEEN ? AND ? and weight BETWEEN ? AND ?", minHeight, maxHeight, minWeight, maxWeight).Find(&pokemons)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("Item not found")
		}
		return nil, result.Error
	}
	return &pokemons, nil
}
