package services

import (
	"pokematch/constants"
	"pokematch/models"
	"pokematch/repositories"
)

type IPokemonService interface {
	FindPokemon(height float32, weight float32) (*[]models.Pokemon, error)
}

type PokemonService struct {
	repository repositories.IPokemonRepository
}

func (s *PokemonService) FindPokemon(height float32, weight float32) (*[]models.Pokemon, error) {
	minHeight := height - constants.HEIGHT_TOLERANCE
	maxHeight := height + constants.HEIGHT_TOLERANCE

	minWeight := weight - constants.WEIGHT_TOLERANCE
	maxWeight := weight + constants.WEIGHT_TOLERANCE

	return s.repository.FindPokemon(minHeight, maxHeight, minWeight, maxWeight)
}

func NewPokemonService(repository repositories.IPokemonRepository) IPokemonService {
	return &PokemonService{repository: repository}
}
