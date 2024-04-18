package repository

import "someApp/internal/entity"

type Pokemon interface {
	CreatePokemon(p *entity.Pokemon) error
}
