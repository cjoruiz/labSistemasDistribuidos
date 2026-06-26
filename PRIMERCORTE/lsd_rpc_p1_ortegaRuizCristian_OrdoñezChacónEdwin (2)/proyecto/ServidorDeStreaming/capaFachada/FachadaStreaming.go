// Package capafachada actúa como intermediario entre el controlador y la capa de acceso a datos
// del servidor de streaming. Encapsula la lógica de lectura y envío de fragmentos de audio.
package capafachada

import (
	"fmt"
	"io"
	"log"

	capaaccesodatos "servidor.streaming.local/grpc-servidor-streaming/capaAccesoDatos"
	pb "servidor.streaming.local/grpc-servidor-streaming/serviciosStreaming"
)

const tamanoFragmento = 32 * 1024 // 32 KB por fragmento

// EnviarFragmentosAudio abre el archivo mp3, lo lee en fragmentos y los envía por el stream gRPC.
// Es invocado por el controlador para mantener la lógica de negocio fuera del controlador.
func EnviarFragmentosAudio(rutaArchivo string, stream pb.ServicioStreaming_ReproducirAudioServer) error {
	log.Printf("[Fachada] EnviarFragmentosAudio invocado: rutaArchivo=%s", rutaArchivo)

	archivo, err := capaaccesodatos.AbrirArchivoMp3(rutaArchivo)
	if err != nil {
		return fmt.Errorf("fachada: %w", err)
	}
	defer archivo.Close()

	buffer := make([]byte, tamanoFragmento)
	numeroFragmento := 0

	for {
		n, err := archivo.Read(buffer)
		if err == io.EOF {
			log.Println("[Fachada] Transmisión completa.")
			break
		}
		if err != nil {
			return fmt.Errorf("fachada: error leyendo archivo: %w", err)
		}

		if n > 0 {
			numeroFragmento++
			fragmento := &pb.FragmentoAudio{Datos: buffer[:n]}
			if err := stream.Send(fragmento); err != nil {
				return fmt.Errorf("fachada: error enviando fragmento #%d: %w", numeroFragmento, err)
			}
			log.Printf("[Fachada] Fragmento #%d enviado (%d bytes)", numeroFragmento, n)
		}
	}

	return nil
}
