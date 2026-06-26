// Package vistas contiene los menús de la interfaz de consola del cliente.
package vistas

import (
	"context"
	"fmt"
	"io"

	capacontroladores "cliente.local/grpc-cliente/capaControladores"
	util "cliente.local/grpc-cliente/utilidades"
)

// MostrarReproduccion inicia la transmisión y reproducción del audio seleccionado.
// El usuario puede detener la reproducción seleccionando la opción "Salir".
func MostrarReproduccion(ctrlStream *capacontroladores.ClienteStreaming, tituloAudio string, rutaArchivo string) {
	util.ImprimirEncabezado("Spotify — Reproducción")
	fmt.Printf("\n  %s\n", tituloAudio)
	util.ImprimirSeparador()
	fmt.Println("\n       Reproduciendo audio ...\n")
	fmt.Println("  1. Salir")
	fmt.Println()

	// Contexto cancelable: al pulsar "Salir" se detiene el stream
	ctx, cancelarStream := context.WithCancel(context.Background())

	reader, writer := io.Pipe()
	canalFin := make(chan struct{})

	// Goroutine: decodifica y reproduce el audio mientras llegan los fragmentos
	go util.DecodificarYReproducir(reader, canalFin)

	// Goroutine: recibe los fragmentos desde el ServidorDeStreaming
	go ctrlStream.RecibirAudio(ctx, rutaArchivo, writer)

	// Goroutine: espera a que el usuario pulse "Salir"
	canalSalida := make(chan struct{})
	go func() {
		for {
			opcion := util.LeerOpcion()
			if opcion == 1 {
				close(canalSalida)
				return
			}
		}
	}()

	// Espera a que termine la reproducción o el usuario salga
	select {
	case <-canalFin:
		cancelarStream()
		fmt.Println("\n  Reproducción finalizada.")
	case <-canalSalida:
		cancelarStream()
		fmt.Println("\n  Reproducción detenida.")
	}
}
