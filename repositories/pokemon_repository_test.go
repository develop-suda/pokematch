package repositories

import (
	"pokematch/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindPokemon(t *testing.T) {
	// SQLiteのインメモリデータベースを使用してテスト用のDBを作成
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// テーブルの自動マイグレーション
	db.AutoMigrate(&models.Pokemon{})

	// テストデータの作成
	db.Create(&models.Pokemon{Name: "Pikachu", Height: 0.4, Weight: 6.0})
	db.Create(&models.Pokemon{Name: "Charizard", Height: 1.7, Weight: 90.5})
	db.Create(&models.Pokemon{Name: "Bulbasaur", Height: 0.7, Weight: 6.9})

	r := NewPokemonRepository(db)

	tests := []struct {
		minHeight float32
		maxHeight float32
		minWeight float32
		maxWeight float32
		expected  int
	}{
		{0.3, 0.5, 5.0, 7.0, 1},    // Pikachuのみ
		{1.0, 2.0, 80.0, 100.0, 1}, // Charizardのみ
		{0.5, 1.0, 5.0, 10.0, 1},   // Bulbasaurのみ
		{0.0, 2.0, 0.0, 100.0, 3},  // 全てのポケモン
		{0.0, 0.0, 0.0, 0.0, 0},    // 該当なし
	}

	for _, test := range tests {
		pokemons, err := r.FindPokemon(test.minHeight, test.maxHeight, test.minWeight, test.maxWeight)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(*pokemons) != test.expected {
			t.Errorf("expected %d pokemons, got %d", test.expected, len(*pokemons))
		}
	}
}
