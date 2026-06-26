package capaFachada

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

func GuardarAudio(nombre string, datos []byte, res *pb.AudioFileResponse) error {
	fmt.Printf("Guardando archivo de audio: %s\n", nombre)

	directorio := "audios"
	err := os.MkdirAll(directorio, 0755)
	if err != nil {
		res.Success = false
		res.Message = "Error al crear directorio: " + err.Error()
		return err
	}

	ruta := fmt.Sprintf("%s/%s", directorio, nombre)
	err = os.WriteFile(ruta, datos, 0644)
	if err != nil {
		res.Success = false
		res.Message = "Error al guardar archivo: " + err.Error()
		return err
	}

	res.Success = true
	res.Message = "Audio almacenado exitosamente"
	res.AudioId = nombre
	fmt.Printf("Archivo guardado: %s (%d bytes)\n", ruta, len(datos))
	return nil
}

type CallbackNotificacion struct {
	IdAudio          string `json:"idAudio"`
	TituloAudio      string `json:"tituloAudio"`
	FechaHoraReproduccion string `json:"fechaHoraReproduccion"`
}

func EnviarCallback(audioId string, fechaHora string, urlCallback string) {
	notificacion := CallbackNotificacion{
		IdAudio:               audioId,
		TituloAudio:          audioId,
		FechaHoraReproduccion: fechaHora,
	}

	jsonData, err := json.Marshal(notificacion)
	if err != nil {
		fmt.Printf("Error al convertir callback a JSON: %v\n", err)
		return
	}

	respuesta, err := http.Post(urlCallback, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error al enviar callback: %v\n", err)
		return
	}
	defer respuesta.Body.Close()

	if respuesta.StatusCode == 200 {
		fmt.Println("Callback enviado exitosamente al administrador")
	} else {
		fmt.Printf("Error en respuesta del callback: %d\n", respuesta.StatusCode)
	}
}

func EnviarCallbackMultiple(audioId string, fechaHora string, urlCallbacks []string) {
	notificacion := CallbackNotificacion{
		IdAudio:               audioId,
		TituloAudio:          audioId,
		FechaHoraReproduccion: fechaHora,
	}

	jsonData, err := json.Marshal(notificacion)
	if err != nil {
		fmt.Printf("Error al convertir callback a JSON: %v\n", err)
		return
	}

	for i, urlCallback := range urlCallbacks {
		respuesta, err := http.Post(urlCallback, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error al enviar callback al administrador #%d: %v\n", i+1, err)
			continue
		}
		defer respuesta.Body.Close()

		if respuesta.StatusCode == 200 {
			fmt.Printf("Callback #%d enviado exitosamente a %s\n", i+1, urlCallback)
		} else {
			fmt.Printf("Error en respuesta del callback #%d: %d\n", i+1, respuesta.StatusCode)
		}
	}
}
