package capaaccesodatos

import (
	"fmt"
	"os"
)

func AbrirArchivoAudio(nombre string) (*os.File, error) {
	directorioBase := "audios"
	ruta := fmt.Sprintf("%s/%s", directorioBase, nombre)

	file, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error al abrir archivo: %s\n", ruta)
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	fmt.Printf("Archivo abierto: %s\n", ruta)
	return file, nil
}
