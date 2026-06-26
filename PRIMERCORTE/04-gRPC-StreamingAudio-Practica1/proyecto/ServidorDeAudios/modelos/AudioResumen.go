// Package modelos define las estructuras de dominio del servidor de metadatos.
package modelos

// AudioResumen contiene los datos básicos identificatorios de un audio.
type AudioResumen struct {
	id               int32
	titulo           string
	idTipo           int32
	rutaArchivoMp3   string // ruta usada por el ServidorDeStreaming
}

// NewAudioResumen crea un AudioResumen con los valores dados.
func NewAudioResumen(id int32, titulo string, idTipo int32, ruta string) AudioResumen {
	return AudioResumen{id: id, titulo: titulo, idTipo: idTipo, rutaArchivoMp3: ruta}
}

func (a *AudioResumen) GetId() int32            { return a.id }
func (a *AudioResumen) GetTitulo() string        { return a.titulo }
func (a *AudioResumen) GetIdTipo() int32         { return a.idTipo }
func (a *AudioResumen) GetRutaArchivoMp3() string { return a.rutaArchivoMp3 }
