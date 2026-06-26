// Package modelos define las estructuras de dominio del servidor de metadatos.
package modelos

// TipoAudio representa una categoría de audio (Música, Podcast, etc.).
type TipoAudio struct {
	id     int32
	nombre string
}

// NewTipoAudio crea un nuevo TipoAudio con los valores proporcionados.
func NewTipoAudio(id int32, nombre string) TipoAudio {
	return TipoAudio{id: id, nombre: nombre}
}

func (t *TipoAudio) GetId() int32     { return t.id }
func (t *TipoAudio) GetNombre() string { return t.nombre }
