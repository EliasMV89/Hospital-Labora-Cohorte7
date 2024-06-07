package main

import (
	"Hospital/servicios"
	"Hospital/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	// Establece la conexión con la base de datos
	db, err := utils.ConectarBaseDeDatos()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Opcines del sistema
	for {
		fmt.Print("Bienvenido!")
		fmt.Print("*******************************************\n")
		fmt.Print("Elige una opcion\n")

		fmt.Print("*******************************************\n")
		fmt.Println("1. Listar pacientes por gravedad de salud")
		fmt.Println("2. Buscar pacientes por nombre o número de seguro social")
		fmt.Println("3. Programar una nueva cita para un paciente con un médico específico")
		fmt.Println("4. Mostrar las citas próximas para un médico específico")
		fmt.Println("5. Actualizar el diagnóstico de un paciente en su historial médico")
		fmt.Println("6. Generar un informe de pacientes atendidos por un médico en un período de tiempo específico")
		fmt.Println("7. Agregar un nuevo paciente")
		fmt.Println("8. Agregar un nuevo medico ")
		fmt.Println("9. Salir del sistema")
		fmt.Println("*********************************")
		fmt.Print("Ingrese su opcion: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Lista de pacientes por gravedad de salud")
			servicios.ListarPacientePorGravedad(db)
		case 2:
			fmt.Println("Buscar paciente por nombre o seguro social")
			fmt.Printf("Ingrese el nombre o numero de seguro social del paciente: ")
			var buscar string
			fmt.Scanln(&buscar)
			servicios.BuscarPacientePorNombreONumero(db, buscar)
		case 3:
			fmt.Println("Programar una nueva cita")
			fmt.Printf("Ingrese el ID del paciente: ")
			var idPaciente int
			fmt.Scanln(&idPaciente)
			fmt.Printf("Ingrese el ID del Medico: ")
			var idMedico int
			fmt.Scanln(&idMedico)
			fmt.Printf("Ingrese la fecha y hora formato(YYYY-MM-DD HH:MM:SS): ")
			var fechaHora string
			fmt.Scanln(&fechaHora)
			nuevaCita := servicios.Cita{
				IdPaciente: idPaciente,
				IdMedico:   idMedico,
				FechaHora:  fechaHora,
			}
			servicios.ProgramarNuevaCita(db, nuevaCita)
		case 4:
			fmt.Println("Mostrar las citas próximas para un médico específico")
			fmt.Printf("Ingrese el ID del medico: ")
			var idMedico int
			fmt.Scanln(&idMedico)
			servicios.ListarCitasPorMedico(db, idMedico)
		case 5:
			fmt.Println("Actualizar el diagnóstico de un paciente en su historial médico")
			fmt.Printf("Ingrese el ID del paciente: ")
			var idPaciente int
			fmt.Scanln(&idPaciente)
			fmt.Printf("Ingrese el diagnostico: ")
			var diagnostico string
			fmt.Scanln(&diagnostico)
			servicios.ActualizarHistorialMedico(db, idPaciente, diagnostico)
		case 6:
			fmt.Println("Informes de pacientes atendidos por un medico en un periodo de tiempo")
			fmt.Printf("Ingrese el ID del medico: ")
			var idMedico int
			fmt.Scanln(&idMedico)
			fmt.Printf("Ingrese la fecha de incio del informe formato(YYYY-MM-DD): ")
			var fechaInicio string
			fmt.Scanln(&fechaInicio)
			fmt.Printf("Ingrese la fecha de fin del informe formato(YYYY-MM-DD): ")
			var fechaFin string
			fmt.Scanln(&fechaFin)
			servicios.GenerarInforme(db, idMedico, fechaInicio, fechaFin)
		case 7:
			fmt.Println("Agregar un nuevo paciente")
			fmt.Printf("Ingrese el nombre del paciente: ")
			var nombre string
			fmt.Scanln(&nombre)
			fmt.Printf("Ingrese el numero de seguro social: ")
			var numeroSS string
			fmt.Scanln(&numeroSS)
			fmt.Printf("Indique la graveda de salud 1, 2 0 3: ")
			var gravedad int
			fmt.Scanln(&gravedad)
			nuevoPaciente := servicios.Paciente{
				Nombre:          nombre,
				NumeroSegSocial: numeroSS,
				GravedadSalud:   gravedad,
			}
			id, err := servicios.AgregarPaciente(db, nuevoPaciente)
			if err != nil {
				log.Fatal(err)
			}
			servicios.AgregarHistorial(db, int(id))
		case 8:
			fmt.Println("Agregar un nuevo medico")
			fmt.Printf("Ingrese el nombre del medico: ")
			var nombre string
			fmt.Scanln(&nombre)
			nuevoMedico := servicios.Medico{
				Nombre: nombre,
			}
			servicios.AgragarMedico(db, nuevoMedico)
		case 9:
			os.Exit(0)
		default:
			fmt.Println("Opcion invalida")
		}
	}
}
