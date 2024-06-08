package controllers

import (
	"Hospital/services"
	"Hospital/utils"
	"encoding/json"
	"net/http"
)

func AgregarMedico(w http.ResponseWriter, r *http.Request) {
	var medico services.Medico
	if err := json.NewDecoder(r.Body).Decode(&medico); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.AgragarMedico(db, medico); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(medico)
}
