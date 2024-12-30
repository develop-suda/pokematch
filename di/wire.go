// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The build tag makes sure the stub is not built in the final build.
package di

// import (
// 	"pokematch/controllers"
// 	"pokematch/repositories"
// 	"pokematch/services"

// 	"github.com/google/wire"
// 	"gorm.io/gorm"
// )

// func InitializeRouter(db *gorm.DB) controllers.IPokemonController {
// 	wire.Build(controllers.NewPokemonController, services.NewPokemonService, repositories.NewPokemonRepository)
// 	return &controllers.PokemonController{}
// }
