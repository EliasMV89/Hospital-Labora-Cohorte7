package services

import (
	"database/sql"
	"log"
)

type Paciente struct {
	Nombre          string `json:"nombre"`
	NumeroSegSocial string `json:"numero_seg_social"`
	GravedadSalud   int    `json:"gravedad_salud"`
}

func AgregarPaciente(db *sql.DB, paciente Paciente) (int64, error) {
	// Consulta para agregar paciente
	queryPaciente := `INSERT INTO Pacientes (Nombre, NumeroSeguroSocial, GravedadSalud) VALUES (?,?,?)`
	// Ejecuta la consulta para agregar el paciente
	result, err := db.Exec(queryPaciente, paciente.Nombre, paciente.NumeroSegSocial, paciente.GravedadSalud)
	if err != nil {
		log.Printf("Error al agregar paciente: %v", err)
		return 0, err
	}

	// Obtiene el ID del paciente recién insertado
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID del paciente: %v", err)
		return 0, err
	}

	// Agrega un registro en el historial médico para el nuevo paciente
	err = AgregarHistorial(db, int(id))
	if err != nil {
		log.Printf("Error al agregar historial médico para el paciente: %v", err)
		return 0, err
	}

	log.Printf("Paciente agregado correctamente con ID: %d", id)
	return id, nil
}

func ListarPacientePorGravedad(db *sql.DB) ([]Paciente, error) {
	query := `SELECT Nombre, NumeroSeguroSocial, GravedadSalud FROM Pacientes ORDER BY GravedadSalud`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error al listar los pacientes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var paciente Paciente
		if err := rows.Scan(&paciente.Nombre, &paciente.NumeroSegSocial, &paciente.GravedadSalud); err != nil {
			log.Printf("Error al leer las filas: %v", err)
			continue
		}
		pacientes = append(pacientes, paciente)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return pacientes, nil
}

func BuscarPacientePorNombreONumero(db *sql.DB, buscar string) ([]Paciente, error) {
	buscar = "%" + buscar + "%"
	query := `SELECT Nombre, NumeroSeguroSocial, GravedadSalud FROM Pacientes WHERE Nombre LIKE ? OR NumeroSeguroSocial LIKE ?`
	rows, err := db.Query(query, buscar, buscar)
	if err != nil {
		log.Printf("Error al buscar Paciente: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var paciente Paciente
		if err := rows.Scan(&paciente.Nombre, &paciente.NumeroSegSocial, &paciente.GravedadSalud); err != nil {
			log.Printf("Error al leer fila: %v", err)
			continue
		}
		pacientes = append(pacientes, paciente)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return nil, err
	}
	return pacientes, nil
}
