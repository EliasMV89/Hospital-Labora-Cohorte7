package servicios

import (
	"database/sql"
	"fmt"
	"log"
)

type Medico struct {
	Nombre string
}

// Funcion para agregar un medico
func AgragarMedico(db *sql.DB, medico Medico) error {
	// Consulta para agregar un medico
	query := `INSERT INTO Medicos (Nombre) VALUES(?)`
	// Ejecuta la consulta
	_, err := db.Exec(query, medico.Nombre)

	if err != nil {
		log.Printf("Error al agragar medico: %v", err)
		return err
	}
	fmt.Println("Medico agregado correctamente.")
	return nil
}
