// Package capacontroladores del cliente encapsula la llamada RPC de streaming al ServidorDeStreaming.
package capacontroladores

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	pbStream "servidor.streaming.local/grpc-servidor-streaming/serviciosStreaming"
)

// ClienteStreaming agrupa la conexión y el stub gRPC al ServidorDeStreaming.
type ClienteStreaming struct {
	conexion *grpc.ClientConn
	stub     pbStream.ServicioStreamingClient
}

// NuevoClienteStreaming crea y devuelve un ClienteStreaming conectado a la dirección indicada.
func NuevoClienteStreaming(direccion string) (*ClienteStreaming, error) {
	conn, err := grpc.Dial(direccion,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al ServidorDeStreaming: %w", err)
	}
	return &ClienteStreaming{
		conexion: conn,
		stub:     pbStream.NewServicioStreamingClient(conn),
	}, nil
}

// Cerrar libera la conexión gRPC.
func (c *ClienteStreaming) Cerrar() {
	c.conexion.Close()
}

// RecibirAudio invoca el RPC ReproducirAudio y escribe cada fragmento recibido en el writer dado.
// Retorna cuando el servidor cierra el stream (EOF) o el contexto es cancelado.
func (c *ClienteStreaming) RecibirAudio(
	ctx context.Context,
	rutaArchivo string,
	writer *io.PipeWriter,
) {
	stream, err := c.stub.ReproducirAudio(ctx, &pbStream.SolicitudStreaming{RutaArchivo: rutaArchivo})
	if err != nil {
		log.Printf("Error iniciando stream: %v", err)
		writer.CloseWithError(err)
		return
	}

	numeroFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("\nTransmisión recibida completa.")
			writer.Close()
			return
		}
		if err != nil {
			// Si el contexto fue cancelado (usuario salió), no es un error grave
			select {
			case <-ctx.Done():
				writer.CloseWithError(ctx.Err())
				return
			default:
				log.Printf("Error recibiendo fragmento: %v", err)
				writer.CloseWithError(err)
				return
			}
		}
		numeroFragmento++
		fmt.Printf("\r  Fragmento #%d recibido (%d bytes) ...", numeroFragmento, len(fragmento.GetDatos()))
		if _, err := writer.Write(fragmento.GetDatos()); err != nil {
			writer.CloseWithError(err)
			return
		}
	}
}
