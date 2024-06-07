package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type historialMedico struct {
	idPaciente  int
	diagnostico string
	fecha       string
}

// Funcion para agregar un historial por cada paciente ingresado
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

// Funcion para actualizar el diagnostico de un paciente
func ActualizarHistorialMedico(db *sql.DB, idPaciente int, diagnostico string) error {
	// Consulta para actualizar el historial medico
	query := `UPDATE HistorialMedico SET Diagnostico = ?, Fecha = NOW() WHERE IDPaciente = ?`
	// Ejecuta la consulta
	_, err := db.Exec(query, diagnostico, idPaciente)

	if err != nil {
		log.Printf("Error al actualizar el historial medico: %v", err)
		return err
	}
	fmt.Println("Historial actualizado correctamente.")
	return nil
}
