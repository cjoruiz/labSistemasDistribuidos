package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	pb "servidor.local/grpc-servidor/serviciosCancion" // ruta generada por protoc
)

func main() {
	// Conectar al servidor gRPC (localhost:50053)
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("No se pudo conectar: %v", err)
		return
	}
	defer conn.Close()

	// Crear cliente
	c := pb.NewServiciosCancionesClient(conn)
	// Contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	//Se captura el titulo del audio a buscar
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese el título del audio a buscar: ")
	tituloLeido, _ := reader.ReadString('\n')
	tituloLeido = strings.TrimSpace(tituloLeido)
	//Se crea un objeto de tipo DTO que contiene el título del audio a buscar
	objPeticion := &pb.PeticionDTO{Titulo: tituloLeido}

	// Llamada al procedimiento remoto buscarAudio
	res, err := c.BuscarAudio(ctx, objPeticion)
	if err != nil {
		fmt.Printf("Error en la llamada gRPC: %v", err)
	}
	// Impresión de la respuesta
	fmt.Printf("\nMensaje: %s", res.Mensaje)
	fmt.Printf("\nCodigo: %d", res.Codigo)
	if res.Codigo == 200 {
		fmt.Printf("\nAudio: %s, Duracion: %d, Tipo: %s, Disponible: %v\n",
			res.ObjAudio.Titulo,
			res.ObjAudio.Duracion,
			res.ObjAudio.Tipo,
			res.ObjAudio.Disponible,
		)
	}
	fmt.Println()
}
