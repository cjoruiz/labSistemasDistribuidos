package capaControladores

import (
	"context"
	"fmt"

	"ServidorDeAudios/capaFachada"
	pb "ServidorDeAudios/protos"
)

type ServidorMetadata struct {
	pb.UnimplementedMetadataServiceServer
}

func (s *ServidorMetadata) GetAudioTypes(ctx context.Context, req *pb.Empty) (*pb.AudioTypesResponse, error) {
	fmt.Printf("RPC Call: GetAudioTypes\n")
	tipos := capaFachada.ObtenerTiposAudio()
	return &pb.AudioTypesResponse{AudioTypes: tipos}, nil
}

func (s *ServidorMetadata) GetAudiosByType(ctx context.Context, req *pb.AudioTypeRequest) (*pb.AudiosListResponse, error) {
	fmt.Printf("RPC Call: GetAudiosByType - Tipo ID: %d\n", req.AudioTypeId)
	audios := capaFachada.ObtenerAudiosPorTipoId(req.AudioTypeId)
	return &pb.AudiosListResponse{Audios: audios}, nil
}

func (s *ServidorMetadata) GetAudioDetails(ctx context.Context, req *pb.AudioDetailsRequest) (*pb.AudioDetailsResponse, error) {
	fmt.Printf("RPC Call: GetAudioDetails - ID: %s\n", req.AudioId)
	audio := capaFachada.ObtenerAudioPorId(req.AudioId)
	if audio == nil {
		return nil, fmt.Errorf("audio no encontrado")
	}
	return audio, nil
}

func (s *ServidorMetadata) GetAudioFilename(ctx context.Context, req *pb.AudioFilenameRequest) (*pb.AudioFilenameResponse, error) {
	fmt.Printf("RPC Call: GetAudioFilename - ID: %s\n", req.AudioId)
	filename := capaFachada.ObtenerNombreArchivo(req.AudioId)
	if filename == "" {
		return nil, fmt.Errorf("audio no encontrado")
	}
	return &pb.AudioFilenameResponse{Filename: filename}, nil
}

func (s *ServidorMetadata) BuscarAudio(ctx context.Context, req *pb.PeticionDTO) (*pb.RespuestaMetadataAudioDTO, error) {
	fmt.Printf("RPC Call: BuscarAudio - Titulo: %s\n", req.GetTitulo())
	return capaFachada.BuscarAudio(req.GetTitulo()), nil
}
