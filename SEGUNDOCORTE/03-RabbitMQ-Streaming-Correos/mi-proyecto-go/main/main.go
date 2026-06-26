package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"ServidorQueAlmacenaCanciones/capaControladores"
	"ServidorQueAlmacenaCanciones/capaFachadaServices/fachada"
)

func main() {
	fmt.Println("Iniciando microservicio de almacenamiento de canciones...")

	// Inicializar Fachada y Controlador
	fachada := capafachadaservices.NuevaFachadaAlmacenamiento()
	controlador := capacontroladores.NuevoControladorAlmacenamientoCanciones(*fachada)

	// Configurar router
	router := mux.NewRouter()
	router.HandleFunc("/canciones/almacenamiento", controlador.AlmacenarAudioCancion).Methods("POST")

	// Iniciar servidor HTTP
	fmt.Println("Servicio de almacenamiento escuchando en el puerto 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}