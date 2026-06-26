// Package capacontroladores implementa los procedimientos remotos gRPC del servidor de metadatos.
// Cada método recibe la solicitud protobuf, delega la lógica a la fachada y devuelve la respuesta protobuf.
package capacontroladores

import (
	"context"
	"fmt"

	capafachada "servidor.audios.local/grpc-servidor-audios/capaFachada"
	pb "servidor.audios.local/grpc-servidor-audios/serviciosMetadatos"
)

// ControladorMetadatos implementa la interfaz ServicioMetadatosServer generada por protoc.
type ControladorMetadatos struct {
	pb.UnimplementedServicioMetadatosServer
}

// ObtenerTiposDeAudio es el procedimiento remoto que devuelve todos los tipos de audio.
func (c *ControladorMetadatos) ObtenerTiposDeAudio(
	ctx context.Context,
	req *pb.SolicitudVacia,
) (*pb.RespuestaTiposAudio, error) {

	fmt.Println("[Servidor] RPC ObtenerTiposDeAudio llamado")

	tipos := capafachada.ConsultarTiposDeAudio()
	var pbTipos []*pb.TipoAudio
	for _, t := range tipos {
		pbTipos = append(pbTipos, &pb.TipoAudio{
			Id:     t.GetId(),
			Nombre: t.GetNombre(),
		})
	}
	return &pb.RespuestaTiposAudio{Tipos: pbTipos}, nil
}

// ObtenerAudiosPorTipo es el procedimiento remoto que devuelve la lista de audios de un tipo.
func (c *ControladorMetadatos) ObtenerAudiosPorTipo(
	ctx context.Context,
	req *pb.SolicitudTipo,
) (*pb.RespuestaListaAudios, error) {

	fmt.Printf("[Servidor] RPC ObtenerAudiosPorTipo llamado con idTipo=%d\n", req.GetIdTipo())

	audios := capafachada.ConsultarAudiosPorTipo(req.GetIdTipo())
	var pbAudios []*pb.AudioResumen
	for _, a := range audios {
		pbAudios = append(pbAudios, &pb.AudioResumen{
			Id:             a.GetId(),
			Titulo:         a.GetTitulo(),
			RutaArchivoMp3: a.GetRutaArchivoMp3(),
		})
	}
	return &pb.RespuestaListaAudios{Audios: pbAudios}, nil
}

// ObtenerDetallesAudio es el procedimiento remoto que devuelve los metadatos completos de un audio.
func (c *ControladorMetadatos) ObtenerDetallesAudio(
	ctx context.Context,
	req *pb.SolicitudAudio,
) (*pb.RespuestaDetallesAudio, error) {

	fmt.Printf("[Servidor] RPC ObtenerDetallesAudio llamado con idAudio=%d\n", req.GetIdAudio())

	registro := capafachada.ConsultarDetallesAudio(req.GetIdAudio())
	if registro == nil {
		return &pb.RespuestaDetallesAudio{}, nil
	}

	resp := &pb.RespuestaDetallesAudio{
		Resumen: &pb.AudioResumen{
			Id:             registro.Resumen.GetId(),
			Titulo:         registro.Resumen.GetTitulo(),
			RutaArchivoMp3: registro.Resumen.GetRutaArchivoMp3(),
		},
	}

	// Copia el bloque de metadatos correspondiente al tipo del audio
	if registro.Musica != nil {
		resp.Musica = &pb.MetadataMusica{
			ArtistaPrincipal:  registro.Musica.GetArtistaPrincipal(),
			Album:             registro.Musica.GetAlbum(),
			GeneroMusical:     registro.Musica.GetGeneroMusical(),
			TituloCancion:     registro.Musica.GetTituloCancion(),
			SelloDiscografico: registro.Musica.GetSelloDiscografico(),
			AnioLanzamiento:   registro.Musica.GetAnioLanzamiento(),
		}
	} else if registro.Podcast != nil {
		resp.Podcast = &pb.MetadataPodcast{
			NombrePodcast:          registro.Podcast.GetNombrePodcast(),
			TituloEpisodio:         registro.Podcast.GetTituloEpisodio(),
			Host:                   registro.Podcast.GetHost(),
			TemporadaEpisodio:      registro.Podcast.GetTemporadaEpisodio(),
			NotasShow:              registro.Podcast.GetNotasShow(),
			ClasificacionContenido: registro.Podcast.GetClasificacionContenido(),
		}
	} else if registro.Audiolibro != nil {
		resp.Audiolibro = &pb.MetadataAudiolibro{
			TituloLibro: registro.Audiolibro.GetTituloLibro(),
			Autor:       registro.Audiolibro.GetAutor(),
			Narrador:    registro.Audiolibro.GetNarrador(),
			Editorial:   registro.Audiolibro.GetEditorial(),
			Isbn:        registro.Audiolibro.GetIsbn(),
			Capitulo:    registro.Audiolibro.GetCapitulo(),
		}
	} else if registro.RuidoBlanco != nil {
		resp.RuidoBlanco = &pb.MetadataRuidoBlanco{
			TipoSonido:          registro.RuidoBlanco.GetTipoSonido(),
			FuenteAudio:         registro.RuidoBlanco.GetFuenteAudio(),
			UsoSugerido:         registro.RuidoBlanco.GetUsoSugerido(),
			ProveedorContenido:  registro.RuidoBlanco.GetProveedorContenido(),
			DuracionBucle:       registro.RuidoBlanco.GetDuracionBucle(),
			FrecuenciaDominante: registro.RuidoBlanco.GetFrecuenciaDominante(),
		}
	}

	return resp, nil
}
