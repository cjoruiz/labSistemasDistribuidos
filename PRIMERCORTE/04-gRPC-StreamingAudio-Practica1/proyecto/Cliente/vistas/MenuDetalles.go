// Package vistas contiene los menús de la interfaz de consola del cliente.
package vistas

import (
	"fmt"

	capacontroladores "cliente.local/grpc-cliente/capaControladores"
	util "cliente.local/grpc-cliente/utilidades"
	pbMeta "servidor.audios.local/grpc-servidor-audios/serviciosMetadatos"
)

// MostrarDetallesAudio obtiene y presenta los metadatos completos de un audio,
// y permite al usuario reproducirlo o regresar al menú anterior.
func MostrarDetallesAudio(
	ctrlMeta *capacontroladores.ClienteMetadatos,
	ctrlStream *capacontroladores.ClienteStreaming,
	idAudio int32,
	rutaArchivoMp3 string,
) {
	detalles, err := ctrlMeta.ObtenerDetallesAudio(idAudio)
	if err != nil {
		fmt.Printf("\n  Error al obtener detalles: %v\n", err)
		return
	}
	if detalles.GetResumen() == nil {
		fmt.Println("\n  Audio no encontrado.")
		return
	}

	for {
		util.ImprimirEncabezado("Spotify — Detalles del Audio")
		fmt.Printf("\n  %s\n", detalles.GetResumen().GetTitulo())
		util.ImprimirSeparador()
		imprimirMetadatos(detalles)
		fmt.Println()
		fmt.Println("  1. Reproducir")
		fmt.Println("  2. Atrás")
		fmt.Println()
		fmt.Print("  Seleccione una opción: ")

		opcion := util.LeerOpcion()

		switch opcion {
		case 1:
			MostrarReproduccion(ctrlStream, detalles.GetResumen().GetTitulo(), rutaArchivoMp3)
		case 2:
			return
		default:
			fmt.Println("\n  Opción no válida. Intente de nuevo.")
		}
	}
}

// imprimirMetadatos muestra los campos específicos según el tipo de audio.
func imprimirMetadatos(detalles *pbMeta.RespuestaDetallesAudio) {
	switch {
	case detalles.GetMusica() != nil:
		m := detalles.GetMusica()
		fmt.Printf("  • Artista Principal  : %s\n", m.GetArtistaPrincipal())
		fmt.Printf("  • Álbum              : %s\n", m.GetAlbum())
		fmt.Printf("  • Género Musical     : %s\n", m.GetGeneroMusical())
		fmt.Printf("  • Título de la Canción: %s\n", m.GetTituloCancion())
		fmt.Printf("  • Sello Discográfico : %s\n", m.GetSelloDiscografico())
		fmt.Printf("  • Año de Lanzamiento : %d\n", m.GetAnioLanzamiento())

	case detalles.GetPodcast() != nil:
		p := detalles.GetPodcast()
		fmt.Printf("  • Nombre del Podcast     : %s\n", p.GetNombrePodcast())
		fmt.Printf("  • Título del Episodio    : %s\n", p.GetTituloEpisodio())
		fmt.Printf("  • Anfitrión (Host)       : %s\n", p.GetHost())
		fmt.Printf("  • Temporada/Episodio     : %s\n", p.GetTemporadaEpisodio())
		fmt.Printf("  • Notas del Show         : %s\n", p.GetNotasShow())
		fmt.Printf("  • Clasificación          : %s\n", p.GetClasificacionContenido())

	case detalles.GetAudiolibro() != nil:
		a := detalles.GetAudiolibro()
		fmt.Printf("  • Título del Libro : %s\n", a.GetTituloLibro())
		fmt.Printf("  • Autor            : %s\n", a.GetAutor())
		fmt.Printf("  • Narrador         : %s\n", a.GetNarrador())
		fmt.Printf("  • Editorial        : %s\n", a.GetEditorial())
		fmt.Printf("  • ISBN             : %s\n", a.GetIsbn())
		fmt.Printf("  • Capítulo         : %d\n", a.GetCapitulo())

	case detalles.GetRuidoBlanco() != nil:
		r := detalles.GetRuidoBlanco()
		fmt.Printf("  • Tipo de Sonido          : %s\n", r.GetTipoSonido())
		fmt.Printf("  • Fuente del Audio        : %s\n", r.GetFuenteAudio())
		fmt.Printf("  • Uso Sugerido            : %s\n", r.GetUsoSugerido())
		fmt.Printf("  • Proveedor de Contenido  : %s\n", r.GetProveedorContenido())
		fmt.Printf("  • Duración del Bucle      : %s\n", r.GetDuracionBucle())
		fmt.Printf("  • Frecuencia Dominante    : %s\n", r.GetFrecuenciaDominante())
	}
}
