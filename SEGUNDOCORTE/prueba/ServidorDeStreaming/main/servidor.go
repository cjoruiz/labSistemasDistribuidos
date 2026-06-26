package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	capaControladores "ServidorDeStreaming/capaControladores"
	pb "ServidorDeStreaming/protos"
	"google.golang.org/grpc"
)

func main() {
	inicializarSistema()
	iniciarServidor()
}

func inicializarSistema() {
	fmt.Println("Sistema inicializado")
}

func iniciarServidor() {
	controladorHTTP := capaControladores.NuevoControladorHTTP()

	go iniciarServidorHTTP(controladorHTTP)
	iniciarServidorGRPC()

	esperarCierreServidor()
}

func iniciarServidorHTTP(controlador *capaControladores.ControladorHTTP) {
	controlador.IniciarServidorHTTP()
	log.Fatal(http.ListenAndServe(controlador.ObtenerPuertoHTTP(), nil))
}

func iniciarServidorGRPC() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStreamingServiceServer(grpcServer, capaControladores.NewServidorStreaming())

	fmt.Println("===============================================")
	fmt.Println("  SERVIDOR DE STREAMING gRPC                  ")
	fmt.Println("===============================================")
	fmt.Println("Puerto gRPC: 50052")
	fmt.Println("Puerto HTTP: 8091")
	fmt.Println("Servicios disponibles:")
	fmt.Println("  - StreamAudio")
	fmt.Println("  - AlmacenarAudio")
	fmt.Println("  - RegistrarCallback")
	fmt.Println("===============================================")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}

func esperarCierreServidor() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	fmt.Println("\nServidor detenido por el usuario")
	os.Exit(0)
}