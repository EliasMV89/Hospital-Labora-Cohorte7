package controllers

import (
	"Hospital/services"
	"Hospital/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AgregarHistorial(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPaciente, err := strconv.Atoi(params["idPaciente"])
	if err != nil {
		http.Error(w, "ID del paciente debe ser un número entero", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.AgregarHistorial(db, idPaciente); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func ActualizarHistorial(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPaciente, err := strconv.Atoi(params["idPaciente"])
	if err != nil {
		http.Error(w, "ID del paciente debe ser un número entero", http.StatusBadRequest)
		return
	}

	var historial struct {
		Diagnostico string `json:"diagnostico"`
	}
	if err := json.NewDecoder(r.Body).Decode(&historial); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.ActualizarHistorialMedico(db, idPaciente, historial.Diagnostico); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
