package main

import (
	"context"
	"log"
	"time"

	menu "Cliente/vistas"
	pb "ServidorDeAudios/protos"
	pbStreaming "ServidorDeStreaming/protos"
	"google.golang.org/grpc"
)

func main() {
	connMetadata, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error al conectar con ServidorDeAudios:", err)
	}
	defer connMetadata.Close()

	connStreaming, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error al conectar con ServidorDeStreaming:", err)
	}
	defer connStreaming.Close()

	clienteMetadata := pb.NewMetadataServiceClient(connMetadata)
	clienteStreaming := pbStreaming.NewStreamingServiceClient(connStreaming)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	menu.InicializarClientes(clienteMetadata, clienteStreaming, ctx)
	menu.MostrarMenuPrincipal()
}
