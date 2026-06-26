package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"proyecto/modelos"
	"proyecto/servicios"
)

func main() {
	vectorMetadataAudios := make([]modelos.MetadataAudio, 5)
	servicios.CargarMetadataAudio(vectorMetadataAudios)
	var opcion int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("==== Menu ====")
		fmt.Println("1. Buscar metadata de un audio")
		fmt.Println("2. Salir")
		fmt.Print("Opción: ")
		fmt.Scan(&opcion)
		//reader.ReadString('\n') // Limpiar el buffer
		switch opcion {
		case 1:
			fmt.Print("Digite el título del audio: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			objRespuesta := servicios.BuscarAudio(titulo, vectorMetadataAudios)

			switch objRespuesta.Codigo {
			case 200:
				fmt.Printf("\n %s", objRespuesta.Mensaje)
				audio := objRespuesta.ObjAudio
				fmt.Printf("\n Título del audio: %s", audio.GetTitulo())
				fmt.Printf("\n Tamaño de audio: %d", audio.GetDuracion())
				fmt.Printf("\n Tipo de audio: %s", audio.GetTipo())
				fmt.Printf("\n El audio está disponible: %t\n", audio.GetDisponible())
				fmt.Printf("\n")
			case 404:
				fmt.Println(objRespuesta.Mensaje)
			}
		case 2:
			fmt.Println("Programa terminado")
			return
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}
}
