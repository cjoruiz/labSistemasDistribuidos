// Package vistas contiene los menús de la interfaz de consola del cliente.
// Sigue el patrón MVC: las vistas muestran datos y capturan entradas,
// delegando la lógica a los controladores.
package vistas

import (
	"fmt"

	capacontroladores "cliente.local/grpc-cliente/capaControladores"
	util "cliente.local/grpc-cliente/utilidades"
)

// MostrarMenuPrincipal presenta el menú raíz de la aplicación y gestiona la navegación.
func MostrarMenuPrincipal(ctrlMeta *capacontroladores.ClienteMetadatos, ctrlStream *capacontroladores.ClienteStreaming) {
	for {
		util.ImprimirEncabezado("Spotify — Menú Principal")
		fmt.Println()
		fmt.Println("  1. Ver tipos de audio")
		fmt.Println("  2. Salir")
		fmt.Println()
		fmt.Print("  Seleccione una opción: ")

		opcion := util.LeerOpcion()

		switch opcion {
		case 1:
			MostrarTiposDeAudio(ctrlMeta, ctrlStream)
		case 2:
			fmt.Println("\n  ¡Hasta pronto!\n")
			return
		default:
			fmt.Println("\n  Opción no válida. Intente de nuevo.")
		}
	}
}
