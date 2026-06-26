package capaControladores

import (
	"fmt"

	"ServidorDeStreaming/capaFachada"
	pb "ServidorDeStreaming/protos"
)

type ServidorStreaming struct {
	pb.UnimplementedStreamingServiceServer
}

func NewServidorStreaming() *ServidorStreaming {
	return &ServidorStreaming{}
}

func (s *ServidorStreaming) StreamAudio(req *pb.StreamRequest, stream pb.StreamingService_StreamAudioServer) error {
	fmt.Printf("RPC Call: StreamAudio - Archivo: %s\n", req.AudioId)
	return capaFachada.ObtenerAudioStream(req.AudioId, stream)
}
