// Package main es el punto de entrada del ServidorDeAudios.
// Inicializa los datos, registra el servicio gRPC y empieza a escuchar.
package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	capaaccesodatos "servidor.audios.local/grpc-servidor-audios/capaAccesoDatos"
	capacontroladores "servidor.audios.local/grpc-servidor-audios/capaControladores"
	pb "servidor.audios.local/grpc-servidor-audios/serviciosMetadatos"
)

const puerto = ":50052"

func main() {
	// Carga los datos de ejemplo en el repositorio en memoria
	capaaccesodatos.InicializarDatos()

	// Abre el socket TCP en el puerto configurado
	listener, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Error al abrir el puerto %s: %v", puerto, err)
	}

	// Crea e inicializa el servidor gRPC
	grpcServer := grpc.NewServer()

	// Registra el controlador de metadatos como implementación del servicio
	pb.RegisterServicioMetadatosServer(grpcServer, &capacontroladores.ControladorMetadatos{})

	fmt.Printf("ServidorDeAudios escuchando en %s...\n", puerto)

	// Bloquea sirviendo solicitudes
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
