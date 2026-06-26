// Package vistas contiene los menús de la interfaz de consola del cliente.
package vistas

import (
	"fmt"

	capacontroladores "cliente.local/grpc-cliente/capaControladores"
	util "cliente.local/grpc-cliente/utilidades"
)

// MostrarListaDeAudios consulta y muestra los audios del tipo seleccionado.
func MostrarListaDeAudios(
	ctrlMeta *capacontroladores.ClienteMetadatos,
	ctrlStream *capacontroladores.ClienteStreaming,
	idTipo int32,
	nombreTipo string,
) {
	audios, err := ctrlMeta.ObtenerAudiosPorTipo(idTipo)
	if err != nil {
		fmt.Printf("\n  Error al obtener audios: %v\n", err)
		return
	}
	if len(audios) == 0 {
		fmt.Println("\n  No hay audios disponibles para este tipo.")
		return
	}

	for {
		util.ImprimirEncabezado("Spotify — " + nombreTipo)
		fmt.Printf("\n  Tipo: %s\n\n", nombreTipo)
		for i, audio := range audios {
			fmt.Printf("  %d. %s\n", i+1, audio.GetTitulo())
		}
		fmt.Printf("  %d. Atrás\n", len(audios)+1)
		fmt.Println()
		fmt.Print("  Seleccione un audio: ")

		opcion := util.LeerOpcion()

		if opcion == len(audios)+1 {
			return
		}
		if opcion >= 1 && opcion <= len(audios) {
			audioSeleccionado := audios[opcion-1]
			MostrarDetallesAudio(ctrlMeta, ctrlStream, audioSeleccionado.GetId(), audioSeleccionado.GetRutaArchivoMp3())
		} else {
			fmt.Println("\n  Opción no válida. Intente de nuevo.")
		}
	}
}
