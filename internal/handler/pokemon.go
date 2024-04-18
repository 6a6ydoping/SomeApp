package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"someApp/api"
)

func (h *Handler) CreatePokemon(w http.ResponseWriter, r *http.Request) {
	var req api.RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println(req)
	err = h.pokemonService.CreatePokemon(&req.Pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
