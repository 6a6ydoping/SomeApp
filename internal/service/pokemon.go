package service

import (
	"errors"
	"regexp"
	"someApp/internal/entity"
)

var (
	ErrInvalidName = errors.New("great name") // error to return
)

func (m *Manager) CreatePokemon(pokemon *entity.Pokemon) error {
	pattern := "^[a-zA-Z]{1,20}$" // regular expression

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	if !regex.MatchString(pokemon.Name) {
		return ErrInvalidName // not a valid name
	}

	err = m.pokemonRepo.CreatePokemon(pokemon)
	return err
}
