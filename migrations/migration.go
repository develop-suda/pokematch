package main

import (
	"pokematch/infra"
	"pokematch/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Pokemon{}); err != nil {
		panic("Failed to migrate database")
	}
}
