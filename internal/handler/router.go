package handler

import (
	"github.com/gorilla/mux"
	"someApp/internal/service"
)

type Handler struct {
	pokemonService service.PokemonService
}

func New(ps service.PokemonService) *Handler {
	return &Handler{
		pokemonService: ps,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/pokemon", h.CreatePokemon).Methods("HEAD")
	return r
}
