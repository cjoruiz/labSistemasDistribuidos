package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"ServidorDeAudios/capaAccesoDatos"
	capaControladores "ServidorDeAudios/capaControladores"
	componenterabbit "ServidorDeAudios/componenteConexionCola"
	pb "ServidorDeAudios/protos"
	"google.golang.org/grpc"
)

const (
	PuertoGRPC   = ":50051"
	PuertoHTTP   = ":8090"
	NombreCola   = "notificaciones_audios"
	NombrePaquete = "ServidorDeAudios"
)

func main() {
	inicializarSistema()
	iniciarServidor()
}

func inicializarSistema() {
	capaaccesodatos.InicializarDatos()
	fmt.Println("Sistema inicializado")
}

func iniciarServidor() {
	publicador := inicializarRabbitMQ()
	controladorHTTP := capaControladores.NuevoControladorHTTP(publicador)

	iniciarServidorHTTP(controladorHTTP)
	iniciarServidorGRPC()

	esperarCierreServidor(publicador)
}

func inicializarRabbitMQ() *componenterabbit.PublicadorRabbit {
	publicador, err := componenterabbit.NuevoPublicadorRabbit(NombreCola)
	if err != nil {
		fmt.Printf("Advertencia: %v\n", err)
		fmt.Println("Continuando sin notificaciones por correo...")
		return nil
	}
	return publicador
}

func iniciarServidorHTTP(controlador *capaControladores.ControladorHTTP) {
	http.HandleFunc("/metadatos", controlador.ManejarMetadatos)
	fmt.Printf("Servidor REST escuchando en puerto %s...\n", PuertoHTTP)
	go func() {
		log.Fatal(http.ListenAndServe(PuertoHTTP, nil))
	}()
}

func iniciarServidorGRPC() {
	lis, err := net.Listen("tcp", PuertoGRPC)
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMetadataServiceServer(grpcServer, &capaControladores.ServidorMetadata{})

	fmt.Println("===============================================")
	fmt.Println("  SERVIDOR DE METADATOS gRPC                 ")
	fmt.Println("===============================================")
	fmt.Printf("Puerto gRPC: %s\n", PuertoGRPC)
	fmt.Printf("Puerto REST: %s\n", PuertoHTTP)
	fmt.Println("Servicios disponibles:")
	fmt.Println("  - GetAudioTypes")
	fmt.Println("  - GetAudiosByType")
	fmt.Println("  - GetAudioDetails")
	fmt.Println("  - GetAudioFilename")
	fmt.Println("  - BuscarAudio")
	fmt.Println("  - POST /metadatos (REST)")
	fmt.Println("===============================================")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error al servir: %v", err)
		}
	}()
}

func esperarCierreServidor(publicador *componenterabbit.PublicadorRabbit) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nCerrando servidor...")

	if publicador != nil {
		publicador.Cerrar()
	}

	fmt.Println("Servidor detenido")
	os.Exit(0)
}