package services

import (
	"database/sql"
	"fmt"
	"log"
)

type HistorialMedico struct {
	IDPaciente  int    `json:"idpaciente"`
	Diagnostico string `json:"diagnostico"`
	Fecha       string `json:"fecha"`
}

// Función para agregar un historial por cada paciente ingresado
func AgregarHistorial(db *sql.DB, idPaciente int) error {
	// Consulta para agregar un registro en historial
	query := `INSERT INTO HistorialMedico (IDPaciente, Diagnostico, Fecha) VALUES(?, 'Ingresado', NOW())`
	// Ejecuta la consulta
	_, err := db.Exec(query, idPaciente)

	if err != nil {
		log.Printf("Error al agregar historial: %v", err)
		return err
	}
	fmt.Println("Historial agregado correctamente.")
	return nil
}

// Función para actualizar el diagnóstico de un paciente
func ActualizarHistorialMedico(db *sql.DB, idPaciente int, diagnostico string) error {
	// Consulta para actualizar el historial médico
	query := `UPDATE HistorialMedico SET Diagnostico = ?, Fecha = NOW() WHERE IDPaciente = ?`
	// Ejecuta la consulta
	_, err := db.Exec(query, diagnostico, idPaciente)

	if err != nil {
		log.Printf("Error al actualizar el historial médico: %v", err)
		return err
	}
	fmt.Println("Historial actualizado correctamente.")
	return nil
}
