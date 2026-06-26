// Package capacontroladores implementa el procedimiento remoto gRPC del servidor de streaming.
package capacontroladores

import (
	"fmt"

	capafachada "servidor.streaming.local/grpc-servidor-streaming/capaFachada"
	pb "servidor.streaming.local/grpc-servidor-streaming/serviciosStreaming"
)

// ControladorStreaming implementa la interfaz ServicioStreamingServer generada por protoc.
type ControladorStreaming struct {
	pb.UnimplementedServicioStreamingServer
}

// ReproducirAudio es el procedimiento remoto que transmite un archivo mp3 en fragmentos al cliente.
func (c *ControladorStreaming) ReproducirAudio(
	req *pb.SolicitudStreaming,
	stream pb.ServicioStreaming_ReproducirAudioServer,
) error {
	fmt.Printf("[Servidor] RPC ReproducirAudio llamado con rutaArchivo=%s\n", req.GetRutaArchivo())
	return capafachada.EnviarFragmentosAudio(req.GetRutaArchivo(), stream)
}
