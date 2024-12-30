package services

import (
	"pokematch/constants"
	"pokematch/models"
	"pokematch/repositories"
)

type IPokemonService interface {
	// ポケモンを見つけるメソッド
	FindPokemon(height float32, weight float32) (*[]models.Pokemon, error)
}

type PokemonService struct {
	repository repositories.IPokemonRepository
}

func (s *PokemonService) FindPokemon(height float32, weight float32) (*[]models.Pokemon, error) {
	// 高さの許容範囲を計算
	minHeight := height - constants.HEIGHT_TOLERANCE
	maxHeight := height + constants.HEIGHT_TOLERANCE

	// 重さの許容範囲を計算
	minWeight := weight - constants.WEIGHT_TOLERANCE
	maxWeight := weight + constants.WEIGHT_TOLERANCE

	return s.repository.FindPokemon(minHeight, maxHeight, minWeight, maxWeight)
}

// 新しいPokemonServiceを作成する関数
func NewPokemonService(repository repositories.IPokemonRepository) IPokemonService {
	return &PokemonService{repository: repository}
}
