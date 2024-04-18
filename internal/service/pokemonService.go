package service

import "someApp/internal/entity"

type PokemonService interface {
	CreatePokemon(pokemon *entity.Pokemon) error
}
