package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"ServidorDeAudios/capaAccesoDatos"
	capacontroladores "ServidorDeAudios/capaControladores"
	pb "ServidorDeAudios/protos"
	"google.golang.org/grpc"
)

func main() {
	capaaccesodatos.InicializarDatos()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMetadataServiceServer(grpcServer, &capacontroladores.ServidorMetadata{})

	fmt.Println("Servidor de Metadatos gRPC escuchando en puerto 50051...")
	fmt.Println("Servicios disponibles:")
	fmt.Println("  - GetAudioTypes")
	fmt.Println("  - GetAudiosByType")
	fmt.Println("  - GetAudioDetails")
	fmt.Println("  - GetAudioFilename")
	fmt.Println("  - BuscarAudio")

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		fmt.Println("\nServidor detenido por el usuario")
		grpcServer.GracefulStop()
		os.Exit(0)
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
