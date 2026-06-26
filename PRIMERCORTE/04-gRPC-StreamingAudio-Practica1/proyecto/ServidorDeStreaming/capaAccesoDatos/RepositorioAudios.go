// Package capaaccesodatos provee acceso a los archivos mp3 almacenados en el servidor de streaming.
package capaaccesodatos

import (
	"fmt"
	"os"
)

// AbrirArchivoMp3 abre el archivo mp3 indicado por su ruta y lo retorna como *os.File.
// Si no se puede abrir, retorna un error descriptivo.
func AbrirArchivoMp3(rutaArchivo string) (*os.File, error) {
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		fmt.Printf("[AccesoDatos] Error: no se pudo abrir el archivo '%s'\n", rutaArchivo)
		return nil, fmt.Errorf("no se pudo abrir el archivo '%s': %w", rutaArchivo, err)
	}
	fmt.Printf("[AccesoDatos] Archivo abierto: %s\n", rutaArchivo)
	return archivo, nil
}
