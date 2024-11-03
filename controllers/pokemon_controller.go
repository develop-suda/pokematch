package controllers

import (
	"net/http"
	"pokematch/services"
	"strconv"

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
	heightStr := ctx.Query("height")
	weightStr := ctx.Query("weight")

	height, err := strconv.ParseFloat(heightStr, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	weight, err := strconv.ParseFloat(weightStr, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	pokemons, err := c.service.FindPokemon(float32(height), float32(weight))
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
