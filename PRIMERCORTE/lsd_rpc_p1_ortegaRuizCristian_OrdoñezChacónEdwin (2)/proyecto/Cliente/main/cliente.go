// Package main es el punto de entrada del Cliente.
// Establece las conexiones gRPC con ambos servidores y lanza el menú principal.
package main

import (
	"fmt"
	"log"

	capacontroladores "cliente.local/grpc-cliente/capaControladores"
	vistas "cliente.local/grpc-cliente/vistas"
)

const (
	direccionServidorAudios    = "localhost:50052"
	direccionServidorStreaming = "localhost:50051"
)

func main() {
	// Conectar al ServidorDeAudios (metadatos)
	ctrlMeta, err := capacontroladores.NuevoClienteMetadatos(direccionServidorAudios)
	if err != nil {
		log.Fatalf("Error conectando al ServidorDeAudios: %v", err)
	}
	defer ctrlMeta.Cerrar()

	// Conectar al ServidorDeStreaming
	ctrlStream, err := capacontroladores.NuevoClienteStreaming(direccionServidorStreaming)
	if err != nil {
		log.Fatalf("Error conectando al ServidorDeStreaming: %v", err)
	}
	defer ctrlStream.Cerrar()

	fmt.Println("  Conexión establecida con ambos servidores.")

	// Iniciar el menú principal
	vistas.MostrarMenuPrincipal(ctrlMeta, ctrlStream)
}
