// Package vistas contiene los menús de la interfaz de consola del cliente.
package vistas

import (
	"fmt"

	capacontroladores "cliente.local/grpc-cliente/capaControladores"
	util "cliente.local/grpc-cliente/utilidades"
)

// MostrarTiposDeAudio consulta los tipos de audio disponibles y permite al usuario seleccionar uno.
func MostrarTiposDeAudio(ctrlMeta *capacontroladores.ClienteMetadatos, ctrlStream *capacontroladores.ClienteStreaming) {
	tipos, err := ctrlMeta.ObtenerTiposDeAudio()
	if err != nil {
		fmt.Printf("\n  Error al obtener tipos de audio: %v\n", err)
		return
	}

	for {
		util.ImprimirEncabezado("Spotify — Tipos de Audio")
		fmt.Println()
		for i, tipo := range tipos {
			fmt.Printf("  %d. %s\n", i+1, tipo.GetNombre())
		}
		fmt.Printf("  %d. Atrás\n", len(tipos)+1)
		fmt.Println()
		fmt.Print("  Seleccione un tipo: ")

		opcion := util.LeerOpcion()

		if opcion == len(tipos)+1 {
			return
		}
		if opcion >= 1 && opcion <= len(tipos) {
			tipoSeleccionado := tipos[opcion-1]
			MostrarListaDeAudios(ctrlMeta, ctrlStream, tipoSeleccionado.GetId(), tipoSeleccionado.GetNombre())
		} else {
			fmt.Println("\n  Opción no válida. Intente de nuevo.")
		}
	}
}
