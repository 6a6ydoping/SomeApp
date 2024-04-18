package service

import "someApp/internal/repository"

type Manager struct {
	pokemonRepo repository.Pokemon
}

func New(pr repository.Pokemon) *Manager {
	return &Manager{
		pokemonRepo: pr,
	}
}
