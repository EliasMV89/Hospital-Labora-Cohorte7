package controllers

import (
	"Hospital/services"
	"Hospital/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ProgramarCita(w http.ResponseWriter, r *http.Request) {
	var cita services.Cita
	if err := json.NewDecoder(r.Body).Decode(&cita); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	if err := services.ProgramarNuevaCita(db, cita); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cita)
}

func ListarCitasMedico(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMedico, err := strconv.Atoi(params["idMedico"])
	if err != nil {
		http.Error(w, "ID del médico debe ser un número entero", http.StatusBadRequest)
		return
	}

	db := utils.GetDB()
	citas, err := services.ListarCitasPorMedico(db, idMedico)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(citas)
}

func GenerarInformeMedico(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMedico, err := strconv.Atoi(params["idMedico"])
	if err != nil {
		http.Error(w, "ID del médico debe ser un número entero", http.StatusBadRequest)
		return
	}

	fechaInicio := r.URL.Query().Get("fechaInicio")
	fechaFin := r.URL.Query().Get("fechaFin")

	db := utils.GetDB()
	informe, err := services.GenerarInforme(db, idMedico, fechaInicio, fechaFin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(informe)
}
