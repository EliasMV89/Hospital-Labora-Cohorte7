package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Cita struct {
	IdPaciente int
	IdMedico   int
	FechaHora  string
}

// Funcion para programar una nueva cita
func ProgramarNuevaCita(db *sql.DB, cita Cita) error {
	// Consulta para insertar una cita
	query := `INSERT INTO Citas (IDPaciente, IDMedico, FechaHora) VALUES (?,?,?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, cita.IdPaciente, cita.IdMedico, cita.FechaHora)

	if err != nil {
		log.Printf("Error al agregar cita: %v", err)
		return err
	}
	fmt.Println("Cita registrada correctamente.")
	return nil
}

// Funcion para mostrar las proximas citas de un medico especifico
func ListarCitasPorMedico(db *sql.DB, idMedico int) error {
	// Consulta para listar proximas citas de un medico
	query := `SELECT * FROM Citas WHERE IDMedico = ? AND FechaHora > NOW() ORDER BY FechaHora`
	// Ejecuta la consulta
	rows, err := db.Query(query, idMedico)

	if err != nil {
		log.Printf("Error al buscar Paciente: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, idPaciente, idMedico int
		var fechaHora string

		err := rows.Scan(&id, &idPaciente, &idMedico, &fechaHora)

		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID: %d, ID del paciente: %d, ID del medico: %d, Fecha y hora: %s\n", id, idPaciente, idMedico, fechaHora)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Pacientes listados correctamente.")
	return nil
}

// Generar un informe de pacientes atendidos por un medico en un periodo de tiempo
func GenerarInforme(db *sql.DB, idMedico int, fechaInicio, fechaFin string) error {
	// Consulta para generar un informe
	query := `SELECT * FROM Citas WHERE IDMedico = ? AND FechaHora BETWEEN ? AND ?`

	// Ejecuta la consulta
	rows, err := db.Query(query, idMedico, fechaInicio, fechaFin)

	if err != nil {
		log.Printf("Error al generar informe: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id, idPaciente, idMedico int
		var fechaHora string

		err := rows.Scan(&id, &idPaciente, &idMedico, &fechaHora)

		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID: %d, ID del paciente: %d, ID del medico: %d, Fecha y hora: %s\n", id, idPaciente, idMedico, fechaHora)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Informe generado correctamente.")
	return nil
}
