package capaControladores

import (
	"context"
	"fmt"
	"time"

	"ServidorDeStreaming/capaFachada"
	pb "ServidorDeStreaming/protos"
)

type ServidorStreaming struct {
	pb.UnimplementedStreamingServiceServer
	callbackURL string
}

var urlCallbacks []string

func NewServidorStreaming() *ServidorStreaming {
	return &ServidorStreaming{}
}

func (s *ServidorStreaming) StreamAudio(req *pb.StreamRequest, stream pb.StreamingService_StreamAudioServer) error {
	fmt.Printf("RPC Call: StreamAudio - Archivo: %s\n", req.AudioId)

	// Enviar callback INMEDIATAMENTE al iniciar la reproducción a todos los administradores
	if len(urlCallbacks) > 0 {
		fechaHora := time.Now().Format("2006-01-02 15:04:05")
		go capaFachada.EnviarCallbackMultiple(req.AudioId, fechaHora, urlCallbacks)
		fmt.Printf("Callback enviado a %d administrador(es): ID=%s, Fecha=%s\n", len(urlCallbacks), req.AudioId, fechaHora)
	}

	// Luego proceder con el streaming del audio
	err := capaFachada.ObtenerAudioStream(req.AudioId, stream)
	return err
}

func (s *ServidorStreaming) AlmacenarAudio(ctx context.Context, req *pb.AudioFileRequest) (*pb.AudioFileResponse, error) {
	fmt.Printf("RPC Call: AlmacenarAudio - Archivo: %s\n", req.Filename)
	var res pb.AudioFileResponse
	err := capaFachada.GuardarAudio(req.Filename, req.Data, &res)
	return &res, err
}

func (s *ServidorStreaming) RegistrarCallback(ctx context.Context, req *pb.CallbackRegistroRequest) (*pb.CallbackRegistroResponse, error) {
	fmt.Printf("RPC Call: RegistrarCallback - URL: %s\n", req.CallbackUrl)
	
	// Agregar URL a la lista (no sobrescribir)
	yaRegistrada := false
	for _, url := range urlCallbacks {
		if url == req.CallbackUrl {
			yaRegistrada = true
			break
		}
	}
	
	if !yaRegistrada {
		urlCallbacks = append(urlCallbacks, req.CallbackUrl)
		fmt.Printf("Total de administradores registrados: %d\n", len(urlCallbacks))
	}
	
	return &pb.CallbackRegistroResponse{
		Success: true,
		Message: "Callback registrado exitosamente",
	}, nil
}
