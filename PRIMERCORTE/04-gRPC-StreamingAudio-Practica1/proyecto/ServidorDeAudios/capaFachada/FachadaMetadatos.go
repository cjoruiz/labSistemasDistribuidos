// Package capafachada actúa como punto de entrada único hacia la capa de acceso a datos.
// Los controladores solo interactúan con este paquete, sin conocer los detalles de almacenamiento.
package capafachada

import (
	"fmt"
	capaaccesodatos "servidor.audios.local/grpc-servidor-audios/capaAccesoDatos"
	"servidor.audios.local/grpc-servidor-audios/modelos"
)

// ConsultarTiposDeAudio delega al repositorio la consulta de tipos de audio.
func ConsultarTiposDeAudio() []modelos.TipoAudio {
	fmt.Println("[Fachada] ConsultarTiposDeAudio invocado")
	return capaaccesodatos.ObtenerTodos()
}

// ConsultarAudiosPorTipo delega al repositorio la consulta de audios filtrados por tipo.
func ConsultarAudiosPorTipo(idTipo int32) []modelos.AudioResumen {
	fmt.Printf("[Fachada] ConsultarAudiosPorTipo invocado con idTipo=%d\n", idTipo)
	return capaaccesodatos.ObtenerAudiosPorTipo(idTipo)
}

// ConsultarDetallesAudio delega al repositorio la consulta de los metadatos completos de un audio.
func ConsultarDetallesAudio(idAudio int32) *modelos.RegistroAudio {
	fmt.Printf("[Fachada] ConsultarDetallesAudio invocado con idAudio=%d\n", idAudio)
	return capaaccesodatos.ObtenerRegistroPorId(idAudio)
}
