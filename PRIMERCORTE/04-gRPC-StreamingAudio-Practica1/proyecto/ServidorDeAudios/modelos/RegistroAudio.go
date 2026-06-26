// Package modelos define las estructuras de dominio del servidor de metadatos.
package modelos

// RegistroAudio agrupa el resumen de un audio junto con sus metadatos específicos
// según el tipo (solo uno de los campos de metadata estará inicializado).
type RegistroAudio struct {
	Resumen     AudioResumen
	Musica      *MetadataMusica
	Podcast     *MetadataPodcast
	Audiolibro  *MetadataAudiolibro
	RuidoBlanco *MetadataRuidoBlanco
}
