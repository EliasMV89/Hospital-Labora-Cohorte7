package services

import (
	"database/sql"
	"fmt"
	"log"
)

type Cita struct {
	IdPaciente int    `json:"id_paciente"`
	IdMedico   int    `json:"id_medico"`
	FechaHora  string `json:"fecha_hora"`
}

// Funcion para programar una nueva cita
func ProgramarNuevaCita(db *sql.DB, cita Cita) error {
	query := `INSERT INTO Citas (IDPaciente, IDMedico, FechaHora) VALUES (?,?,?)`
	_, err := db.Exec(query, cita.IdPaciente, cita.IdMedico, cita.FechaHora)
	if err != nil {
		log.Printf("Error al agregar cita: %v", err)
		return err
	}
	fmt.Println("Cita registrada correctamente.")
	return nil
}

// Funcion para mostrar las proximas citas de un medico especifico
func ListarCitasPorMedico(db *sql.DB, idMedico int) ([]Cita, error) {
	query := `SELECT IDPaciente, IDMedico, FechaHora FROM Citas WHERE IDMedico = ? AND FechaHora > NOW() ORDER BY FechaHora`
	rows, err := db.Query(query, idMedico)
	if err != nil {
		log.Printf("Error al buscar citas: %v", err)
		return nil, err
	}
	defer rows.Close()

	var citas []Cita
	for rows.Next() {
		var cita Cita
		if err := rows.Scan(&cita.IdPaciente, &cita.IdMedico, &cita.FechaHora); err != nil {
			log.Printf("Error al leer fila: %v", err)
			continue
		}
		citas = append(citas, cita)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return citas, nil
}

// Generar un informe de pacientes atendidos por un médico en un periodo de tiempo
func GenerarInforme(db *sql.DB, idMedico int, fechaInicio, fechaFin string) ([]Cita, error) {
	query := `SELECT IDPaciente, IDMedico, FechaHora FROM Citas WHERE IDMedico = ? AND FechaHora BETWEEN ? AND ?`
	rows, err := db.Query(query, idMedico, fechaInicio, fechaFin)
	if err != nil {
		log.Printf("Error al generar informe: %v", err)
		return nil, err
	}
	defer rows.Close()

	var informe []Cita
	for rows.Next() {
		var cita Cita
		if err := rows.Scan(&cita.IdPaciente, &cita.IdMedico, &cita.FechaHora); err != nil {
			log.Printf("Error al leer fila: %v", err)
			continue
		}
		informe = append(informe, cita)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return informe, nil
}
