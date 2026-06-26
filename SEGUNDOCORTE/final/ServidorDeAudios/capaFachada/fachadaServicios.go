package capaFachada

import (
	"ServidorDeAudios/capaAccesoDatos"
	pb "ServidorDeAudios/protos"
)

func ObtenerTiposAudio() []*pb.AudioType {
	tipos := capaaccesodatos.ObtenerTiposAudio()
	result := make([]*pb.AudioType, len(tipos))
	for i, tipo := range tipos {
		result[i] = &pb.AudioType{
			Id:     tipo.Id,
			Nombre: tipo.Nombre,
		}
	}
	return result
}

func ObtenerAudiosPorTipoId(tipoId int32) []*pb.AudioSummary {
	audios := capaaccesodatos.ObtenerAudiosPorTipoId(int(tipoId))
	result := make([]*pb.AudioSummary, len(audios))
	for i, audio := range audios {
		result[i] = &pb.AudioSummary{
			Id:     audio.GetTitulo(),
			Titulo: audio.GetTitulo(),
		}
	}
	return result
}

func ObtenerAudioPorId(id string) *pb.AudioDetailsResponse {
	audio := capaaccesodatos.ObtenerAudioPorId(id)
	if audio == nil {
		return nil
	}

	return &pb.AudioDetailsResponse{
		Id:                      audio.GetTitulo(),
		Titulo:                  audio.GetTitulo(),
		Autor:                   audio.GetAutor(),
		Album:                   audio.GetAlbum(),
		Genero:                  audio.GetGenero(),
		Duracion:                int32(audio.GetDuracion()),
		TipoId:                  int32(audio.GetTipoId()),
		TipoNombre:              audio.GetTipo(),
		FechaLanzamiento:        audio.GetFechaLanzamiento(),
		Disponible:              audio.GetDisponible(),
		NombreArchivo:           audio.GetNombreArchivo(),
		SelloDiscografico:       audio.GetSelloDiscografico(),
		NombrePodcast:           audio.GetNombrePodcast(),
		NumeroTemporadaEpisodio: audio.GetNumeroTemporadaEpisodio(),
		NotasShow:               audio.GetNotasShow(),
		ClasificacionContenido:  audio.GetClasificacionContenido(),
		Narrador:                audio.GetNarrador(),
		Editorial:               audio.GetEditorial(),
		Isbn:                    audio.GetIsbn(),
		Capitulo:                audio.GetCapitulo(),
		TipoSonido:              audio.GetTipoDeSonido(),
		FuenteAudio:             audio.GetFuenteAudio(),
		UsoSugerido:             audio.GetUsoSugerido(),
		ProveedorContenido:      audio.GetProveedorContenido(),
		DuracionBucle:           audio.GetDuracionBucle(),
		FrecuenciaDominante:     audio.GetFrecuenciaDominante(),
	}
}

func ObtenerNombreArchivo(id string) string {
	return capaaccesodatos.ObtenerNombreArchivo(id)
}

func BuscarAudio(titulo string) *pb.RespuestaMetadataAudioDTO {
	respuesta := capaaccesodatos.BuscarAudio(titulo)

	metadataProto := &pb.MetadataAudio{
		Titulo:                  respuesta.ObjAudio.GetTitulo(),
		Autor:                   respuesta.ObjAudio.GetAutor(),
		Album:                   respuesta.ObjAudio.GetAlbum(),
		Genero:                  respuesta.ObjAudio.GetGenero(),
		Duracion:                int32(respuesta.ObjAudio.GetDuracion()),
		TipoId:                  int32(respuesta.ObjAudio.GetTipoId()),
		TipoNombre:              respuesta.ObjAudio.GetTipo(),
		FechaLanzamiento:        respuesta.ObjAudio.GetFechaLanzamiento(),
		Disponible:              respuesta.ObjAudio.GetDisponible(),
		NombreArchivo:           respuesta.ObjAudio.GetNombreArchivo(),
		SelloDiscografico:       respuesta.ObjAudio.GetSelloDiscografico(),
		NombrePodcast:           respuesta.ObjAudio.GetNombrePodcast(),
		NumeroTemporadaEpisodio: respuesta.ObjAudio.GetNumeroTemporadaEpisodio(),
		NotasShow:               respuesta.ObjAudio.GetNotasShow(),
		ClasificacionContenido:  respuesta.ObjAudio.GetClasificacionContenido(),
		Narrador:                respuesta.ObjAudio.GetNarrador(),
		Editorial:               respuesta.ObjAudio.GetEditorial(),
		Isbn:                    respuesta.ObjAudio.GetIsbn(),
		Capitulo:                respuesta.ObjAudio.GetCapitulo(),
		TipoSonido:              respuesta.ObjAudio.GetTipoDeSonido(),
		FuenteAudio:             respuesta.ObjAudio.GetFuenteAudio(),
		UsoSugerido:             respuesta.ObjAudio.GetUsoSugerido(),
		ProveedorContenido:      respuesta.ObjAudio.GetProveedorContenido(),
		DuracionBucle:           respuesta.ObjAudio.GetDuracionBucle(),
		FrecuenciaDominante:     respuesta.ObjAudio.GetFrecuenciaDominante(),
	}

	return &pb.RespuestaMetadataAudioDTO{
		ObjAudio: metadataProto,
		Codigo:   int32(respuesta.Codigo),
		Mensaje:  respuesta.Mensaje,
	}
}
