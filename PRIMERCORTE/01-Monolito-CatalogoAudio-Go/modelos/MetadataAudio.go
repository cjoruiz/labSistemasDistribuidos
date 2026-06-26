package modelos

type MetadataAudio struct {
	titulo string
	duracion int
	tipo string
	disponible bool
}

func(m *MetadataAudio) GetTitulo() string {
	return m.titulo
}

func (m *MetadataAudio) SetTitulo(titulo string){
	m.titulo = titulo
}
func(m *MetadataAudio) GetDuracion() int {
	return m.duracion
}

func (m *MetadataAudio) SetDuracion(duracion int){
	m.duracion = duracion
}
func(m *MetadataAudio) GetTipo() string {
	return m.tipo
}

func (m *MetadataAudio) SetTipo(tipo string){
	m.tipo = tipo
}
func(m *MetadataAudio) GetDisponible() bool {
	return m.disponible
}

func (m *MetadataAudio) SetDisponible(disponible bool){
	m.disponible = disponible
}