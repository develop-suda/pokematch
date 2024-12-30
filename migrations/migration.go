package main

import (
	"fmt"
	"pokematch/infra"
	"pokematch/models"
)

// pokemonsテーブルを削除し再作成後、データ投入
func main() {
	infra.Initialize()    // インフラの初期化
	db := infra.SetupDB() // データベースのセットアップ

	// Pokemonテーブルを削除し再作成
	db.Migrator().DropTable(&models.Pokemon{})                // 既存のPokemonテーブルを削除
	if err := db.AutoMigrate(&models.Pokemon{}); err != nil { // Pokemonテーブルを再作成
		panic("Failed to migrate database") // エラーが発生した場合、パニックを起こす
	}

	// Pokemonテーブルにデータを投入
	list := models.ReturnList() // データリストを取得
	result := db.Create(&list)  // データベースにデータを作成

	if result.Error != nil { // エラーが発生した場合
		fmt.Println(result.Error) // エラー内容を出力
	}
}
