package controllers

import (
	"Hospital/services"
	"Hospital/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AgregarPaciente maneja la solicitud para agregar un nuevo paciente.
func AgregarPaciente(w http.ResponseWriter, r *http.Request) {
	var paciente services.Paciente
	if err := json.NewDecoder(r.Body).Decode(&paciente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	id, err := services.AgregarPaciente(db, paciente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

// ListarPacientesPorGravedad maneja la solicitud para listar pacientes por gravedad de salud.
func ListarPacientesPorGravedad(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	pacientes, err := services.ListarPacientePorGravedad(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pacientes)
}

// BuscarPacientePorNombreONumero maneja la solicitud para buscar un paciente por nombre o número de seguro social.
func BuscarPacientePorNombreONumero(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	buscar := params["buscar"]

	db := utils.GetDB()
	pacientes, err := services.BuscarPacientePorNombreONumero(db, buscar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pacientes)
}
