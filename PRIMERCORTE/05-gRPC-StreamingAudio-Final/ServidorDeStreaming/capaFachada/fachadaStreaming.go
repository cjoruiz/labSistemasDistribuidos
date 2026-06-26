package capaFachada

import (
	"fmt"
	"io"

	"ServidorDeStreaming/capaAccesoDatos"
	pb "ServidorDeStreaming/protos"
)

func ObtenerAudioStream(nombre string, stream pb.StreamingService_StreamAudioServer) error {
	fmt.Printf("Solicitando streaming de audio: %s\n", nombre)

	file, err := capaaccesodatos.AbrirArchivoAudio(nombre)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer file.Close()

	buf := make([]byte, 32*1024)
	chunkNum := 0

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			fmt.Println("Audio enviado completamente.")
			break
		}
		if err != nil {
			return fmt.Errorf("error al leer el archivo: %w", err)
		}

		chunkNum++

		if n > 0 {
			chunk := &pb.AudioChunk{Data: buf[:n]}
			if err := stream.Send(chunk); err != nil {
				return fmt.Errorf("error enviando chunk #%d: %w", chunkNum, err)
			}
			fmt.Printf("Chunk #%d enviado (%d bytes)\n", chunkNum, n)
		}
	}

	return nil
}
