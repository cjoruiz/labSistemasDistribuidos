package capafachadaservices

import (
	"fmt"
	"io"
	"log"
	"os"
)

func StreamAudioFile(tituloCancion string, funcionParaEnviarFragmento func([]byte) error) error {
	log.Printf("Canci\u00f3n solicitada: %s", tituloCancion)
	file, err := os.Open("../audio_files/" + tituloCancion)
	if err != nil {
		file, err = os.Open("../ServidorStreaming/canciones/" + tituloCancion)
		if err != nil {
			file, err = os.Open("canciones/" + tituloCancion)
			if err != nil {
				return fmt.Errorf("no se pudo abrir el archivo: %w", err)
			}
		}
	}
	defer file.Close()

	buffer := make([]byte, 64*1024)
	fragmento := 0

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			log.Println("Canci\u00f3n enviada completamente desde la fachada.")
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo el archivo: %w", err)
		}

		fragmento++
		log.Printf("Fragmento #%d le\u00eddo (%d bytes) y enviando", fragmento, n)
		err = funcionParaEnviarFragmento(buffer[:n])
		if err != nil {
			return fmt.Errorf("error enviando fragmento #%d: %w", fragmento, err)
		}
	}

	return nil
}
