// Package utilidades provee funciones auxiliares del cliente,
// incluyendo la decodificación y reproducción de audio MP3 en tiempo real.
package utilidades

import (
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// DecodificarYReproducir decodifica el stream mp3 proveniente del reader y lo reproduce.
// Cierra canalFin cuando la reproducción termina, permitiendo sincronización con la goroutine llamante.
func DecodificarYReproducir(reader io.Reader, canalFin chan struct{}) {
	streamer, formato, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Printf("Error decodificando MP3: %v", err)
		close(canalFin)
		return
	}
	defer streamer.Close()

	speaker.Init(formato.SampleRate, formato.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalFin)
	})))
}
