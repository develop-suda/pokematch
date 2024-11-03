package main

import (
	"fmt"
	"pokematch/infra"
	"pokematch/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Pokemon{}); err != nil {
		panic("Failed to migrate database")
	}

	// テーブルの中身を全て削除
	db.Unscoped().Where("1=1").Delete(&models.Pokemon{})
	list := models.ReturnList()
	result := db.Create(&list)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
