package utilidades

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	pb "ServidorDeStreaming/protos"
)

func ReproducirAudio(
	ctx context.Context,
	clienteStreaming pb.StreamingServiceClient,
	nombreArchivo string,
	canalSincronizacion chan struct{}) {

	fmt.Println("Recibiendo y reproduciendo audio en vivo...")
	reader, writer := io.Pipe()

	stream, err := clienteStreaming.StreamAudio(ctx, &pb.StreamRequest{AudioId: nombreArchivo})
	if err != nil {
		fmt.Printf("Error al iniciar streaming: %v\n", err)
		return
	}

	go DecodificarReproducir(reader, canalSincronizacion)
	go RecibirAudio(ctx, stream, writer, canalSincronizacion)
}

func EsperarFinReproduccion(canalSincronizacion chan struct{}, cancel context.CancelFunc) {
	entrada := make(chan struct{}, 1)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		entrada <- struct{}{}
	}()

	select {
	case <-canalSincronizacion:
		fmt.Println("\n¡Reproducción completada!")
	case <-entrada:
		fmt.Println("\nDeteniendo reproducción...")
		cancel()
		<-canalSincronizacion
	}

	time.Sleep(300 * time.Millisecond)
}