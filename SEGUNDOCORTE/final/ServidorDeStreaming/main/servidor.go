package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	capacontroladores "ServidorDeStreaming/capaControladores"
	pb "ServidorDeStreaming/protos"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStreamingServiceServer(grpcServer, capacontroladores.NewServidorStreaming())

	fmt.Println("Servidor de Streaming gRPC escuchando en puerto 50052...")
	fmt.Println("Servicios disponibles:")
	fmt.Println("  - StreamAudio")

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
