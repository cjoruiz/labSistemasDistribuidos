// Package main es el punto de entrada del ServidorDeStreaming.
// Registra el servicio gRPC de streaming y empieza a escuchar conexiones.
package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	capacontroladores "servidor.streaming.local/grpc-servidor-streaming/capaControladores"
	pb "servidor.streaming.local/grpc-servidor-streaming/serviciosStreaming"
)

const puerto = ":50051"

func main() {
	listener, err := net.Listen("tcp", puerto)
	if err != nil {
		log.Fatalf("Error al abrir el puerto %s: %v", puerto, err)
	}

	// Se crea el servidor gRPC sin límite de tamaño de mensaje (necesario para chunks de audio)
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(10*1024*1024),
		grpc.MaxSendMsgSize(10*1024*1024),
	)

	// Registra el controlador de streaming como implementación del servicio
	pb.RegisterServicioStreamingServer(grpcServer, &capacontroladores.ControladorStreaming{})

	fmt.Printf("ServidorDeStreaming escuchando en %s...\n", puerto)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
