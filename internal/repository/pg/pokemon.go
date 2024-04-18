package pg

import (
	"context"
	"fmt"
	"someApp/internal/entity"
)

func (p *Postgres) CreatePokemon(pokemon *entity.Pokemon) error {
	query := fmt.Sprintf(`
			INSERT INTO %s VALUES (123, $1)`, pokemonTable)

	_, err := p.Pool.Exec(context.Background(), query, pokemon.Name)
	if err != nil {
		return err
	}

	return nil
}
