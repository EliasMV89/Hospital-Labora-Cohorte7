package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Paciente struct {
	Nombre          string
	NumeroSegSocial string
	GravedadSalud   int
}

func AgregarPaciente(db *sql.DB, paciente Paciente) (int64, error) {
	// Consulta para agregar paciente
	queryPaciente := `INSERT INTO Pacientes (Nombre, NumeroSeguroSocial, GravedadSalud) VALUES (?,?,?)`
	// Ejecuta la consulta
	result, err := db.Exec(queryPaciente, paciente.Nombre, paciente.NumeroSegSocial, paciente.GravedadSalud)

	if err != nil {
		log.Printf("Error al agregar paciente: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error al obtener el ID del paciente: %v", err)
		return 0, err
	}

	fmt.Printf("Paciente agregado correctamente con ID: %d\n", id)
	return id, nil
}

// Funcion para listar pacienten por gravedad de salud
func ListarPacientePorGravedad(db *sql.DB) error {
	// Consulta para listar pacientes
	query := `SELECT * FROM Pacientes ORDER BY GravedadSalud`
	rows, err := db.Query(query)

	if err != nil {
		log.Printf("Error al listar los pacientes: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nombre, numeroSegSocial string
		var id, gravedadSalud int

		err := rows.Scan(&id, &nombre, &numeroSegSocial, &gravedadSalud)
		if err != nil {
			log.Printf("Error al leer las filas: %v", err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Numero de seguro social: %s, Grevedad salud: %d\n", id, nombre, numeroSegSocial, gravedadSalud)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Pacientes listados correctamente.")
	return nil
}

// Buscar paciente por nombre o numero de seguro social
func BuscarPacientePorNombreONumero(db *sql.DB, buscar string) error {
	// Agrega los caracteres de comodín al inicio y al final de la cadena de búsqueda
	buscar = "%" + buscar + "%"

	// Consulta para buscar paciente
	query := `SELECT * FROM Pacientes WHERE Nombre LIKE ? OR NumeroSeguroSocial = ?`
	rows, err := db.Query(query, buscar, buscar)

	if err != nil {
		log.Printf("Error al buscar Paciente: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var nombre, numeroSegSocial string
		var id, gravedadSalud int

		err := rows.Scan(&id, &nombre, &numeroSegSocial, &gravedadSalud)

		if err != nil {
			log.Printf("Error al leer fila: %v", err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Numero de seguro social: %s, Grevedad salud: %d\n", id, nombre, numeroSegSocial, gravedadSalud)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error al iterar filas: %v", err)
		return err
	}
	fmt.Println("Pacientes listados correctamente.")
	return nil
}
