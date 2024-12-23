package main

import (
	"fmt"
	"pokematch/infra"
	"pokematch/models"
)

// pokemonsテーブルを削除し再作成後、データ投入
func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// Pokemonテーブルを削除し再作成
	db.Migrator().DropTable(&models.Pokemon{})
	if err := db.AutoMigrate(&models.Pokemon{}); err != nil {
		panic("Failed to migrate database")
	}

	// Pokemonテーブルにデータを投入
	list := models.ReturnList()
	result := db.Create(&list)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
