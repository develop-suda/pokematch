// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject

package di

import (
	"gorm.io/gorm"
	"pokematch/controllers"
	"pokematch/repositories"
	"pokematch/services"
)

// Injectors from wire.go:

func InitializeRouter(db *gorm.DB) controllers.IPokemonController {
	iPokemonRepository := repositories.NewPokemonRepository(db)
	iPokemonService := services.NewPokemonService(iPokemonRepository)
	iPokemonController := controllers.NewPokemonController(iPokemonService)
	return iPokemonController
}
