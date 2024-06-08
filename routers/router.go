package routers

import (
	"Hospital/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Endpoints para Pacientes
	router.HandleFunc("/pacientes", controllers.AgregarPaciente).Methods("POST")
	router.HandleFunc("/pacientes", controllers.ListarPacientesPorGravedad).Methods("GET")
	router.HandleFunc("/pacientes/buscar/{buscar}", controllers.BuscarPacientePorNombreONumero).Methods("GET")

	// Endpoints para Medicos
	router.HandleFunc("/medicos", controllers.AgregarMedico).Methods("POST")

	// Endpoints para Citas
	router.HandleFunc("/citas", controllers.ProgramarCita).Methods("POST")
	router.HandleFunc("/citas/{idMedico}", controllers.ListarCitasMedico).Methods("GET")
	router.HandleFunc("/citas/informe/{idMedico}/{fechaInicio}/{fechaFin}", controllers.GenerarInformeMedico).Methods("GET")

	// Endpoints para Historial Medico
	router.HandleFunc("/historial/{idPaciente}", controllers.AgregarHistorial).Methods("POST")
	router.HandleFunc("/historial/{idPaciente}", controllers.ActualizarHistorial).Methods("PUT")

	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
	return router
}
