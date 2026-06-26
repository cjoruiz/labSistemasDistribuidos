package capaaccesoadatos

import (
	"fmt"
	"os"
	"sync"
	"path/filepath"
)

type RepositorioCanciones struct {
	mu sync.Mutex
}

var (
	instancia *RepositorioCanciones
	once      sync.Once
)

// GetRepositorioCanciones aplica patrón Singleton
func GetRepositorioCanciones() *RepositorioCanciones {
	once.Do(func() {
		instancia = &RepositorioCanciones{}
	})
	return instancia
}

// GuardarCancion guarda el audio con formato titulo_genero_artista.mp3
func (r *RepositorioCanciones) GuardarCancion(titulo string, genero string, artista string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Crear carpeta si no existe
	os.MkdirAll("audios", os.ModePerm)

	// Construir nombre del archivo
	fileName := fmt.Sprintf("%s_%s_%s.mp3", titulo, genero, artista)
	filePath := filepath.Join("audios", fileName)

	// Guardar archivo
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("Error al guardar archivo: %v\n", err)		
	}
	//crear registro en memoria
	return nil
}