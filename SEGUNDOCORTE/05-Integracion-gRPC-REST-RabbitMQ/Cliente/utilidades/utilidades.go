package utilidades

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	pb "ServidorDeStreaming/protos"
)

func LimpiarPantalla() {
	fmt.Print("\033[2J\033[H")
}

func DecodificarReproducir(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		fmt.Printf("Error decodificando MP3: %v\n", err)
		return
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))
	if err != nil {
		fmt.Printf("Error inicializando speaker: %v\n", err)
		return
	}

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		select {
		case canalSincronizacion <- struct{}{}:
		default:
		}
	})))
}

func RecibirAudio(
	ctx context.Context,
	stream pb.StreamingService_StreamAudioClient,
	writer *io.PipeWriter,
	canalSincronizacion chan struct{}) {
	noFragmento := 0
	for {
		select {
		case <-ctx.Done():
			writer.Close()
			return
		default:
		}

		fragmento, err := stream.Recv()
		if err == io.EOF {
			writer.Close()
			break
		}
		if err != nil {
			writer.Close()
			break
		}
		noFragmento++
		fmt.Printf("\nFragmento #%d recibido (%d bytes)", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			break
		}
	}

	select {
	case canalSincronizacion <- struct{}{}:
	default:
	}
}
