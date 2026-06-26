package modelos

type MetadataAudio struct {
	titulo                  string
	autor                   string
	album                   string
	genero                  string
	duracion                int
	tipo                    string
	tipoId                  int
	fechaLanzamiento        string
	disponible              bool
	nombreArchivo           string
	selloDiscografico       string
	nombrePodcast           string
	numeroTemporadaEpisodio string
	notasShow               string
	clasificacionContenido  string
	narrador                string
	editorial               string
	isbn                    string
	capitulo                string
	tipoDeSonido            string
	fuenteAudio             string
	usoSugerido             string
	proveedorContenido      string
	duracionBucle           string
	frecuenciaDominante     string
}

func (m *MetadataAudio) GetTitulo() string {
	return m.titulo
}

func (m *MetadataAudio) SetTitulo(titulo string) {
	m.titulo = titulo
}

func (m *MetadataAudio) GetAutor() string {
	return m.autor
}

func (m *MetadataAudio) SetAutor(autor string) {
	m.autor = autor
}

func (m *MetadataAudio) GetAlbum() string {
	return m.album
}

func (m *MetadataAudio) SetAlbum(album string) {
	m.album = album
}

func (m *MetadataAudio) GetGenero() string {
	return m.genero
}

func (m *MetadataAudio) SetGenero(genero string) {
	m.genero = genero
}

func (m *MetadataAudio) GetDuracion() int {
	return m.duracion
}

func (m *MetadataAudio) SetDuracion(duracion int) {
	m.duracion = duracion
}

func (m *MetadataAudio) GetTipo() string {
	return m.tipo
}

func (m *MetadataAudio) SetTipo(tipo string) {
	m.tipo = tipo
}

func (m *MetadataAudio) GetTipoId() int {
	return m.tipoId
}

func (m *MetadataAudio) SetTipoId(tipoId int) {
	m.tipoId = tipoId
}

func (m *MetadataAudio) GetFechaLanzamiento() string {
	return m.fechaLanzamiento
}

func (m *MetadataAudio) SetFechaLanzamiento(fechaLanzamiento string) {
	m.fechaLanzamiento = fechaLanzamiento
}

func (m *MetadataAudio) GetDisponible() bool {
	return m.disponible
}

func (m *MetadataAudio) SetDisponible(disponible bool) {
	m.disponible = disponible
}

func (m *MetadataAudio) GetNombreArchivo() string {
	return m.nombreArchivo
}

func (m *MetadataAudio) SetNombreArchivo(nombreArchivo string) {
	m.nombreArchivo = nombreArchivo
}

func (m *MetadataAudio) GetSelloDiscografico() string {
	return m.selloDiscografico
}

func (m *MetadataAudio) SetSelloDiscografico(selloDiscografico string) {
	m.selloDiscografico = selloDiscografico
}

func (m *MetadataAudio) GetNombrePodcast() string {
	return m.nombrePodcast
}

func (m *MetadataAudio) SetNombrePodcast(nombrePodcast string) {
	m.nombrePodcast = nombrePodcast
}

func (m *MetadataAudio) GetNumeroTemporadaEpisodio() string {
	return m.numeroTemporadaEpisodio
}

func (m *MetadataAudio) SetNumeroTemporadaEpisodio(numeroTemporadaEpisodio string) {
	m.numeroTemporadaEpisodio = numeroTemporadaEpisodio
}

func (m *MetadataAudio) GetNotasShow() string {
	return m.notasShow
}

func (m *MetadataAudio) SetNotasShow(notasShow string) {
	m.notasShow = notasShow
}

func (m *MetadataAudio) GetClasificacionContenido() string {
	return m.clasificacionContenido
}

func (m *MetadataAudio) SetClasificacionContenido(clasificacionContenido string) {
	m.clasificacionContenido = clasificacionContenido
}

func (m *MetadataAudio) GetNarrador() string {
	return m.narrador
}

func (m *MetadataAudio) SetNarrador(narrador string) {
	m.narrador = narrador
}

func (m *MetadataAudio) GetEditorial() string {
	return m.editorial
}

func (m *MetadataAudio) SetEditorial(editorial string) {
	m.editorial = editorial
}

func (m *MetadataAudio) GetIsbn() string {
	return m.isbn
}

func (m *MetadataAudio) SetIsbn(isbn string) {
	m.isbn = isbn
}

func (m *MetadataAudio) GetCapitulo() string {
	return m.capitulo
}

func (m *MetadataAudio) SetCapitulo(capitulo string) {
	m.capitulo = capitulo
}

func (m *MetadataAudio) GetTipoDeSonido() string {
	return m.tipoDeSonido
}

func (m *MetadataAudio) SetTipoDeSonido(tipoDeSonido string) {
	m.tipoDeSonido = tipoDeSonido
}

func (m *MetadataAudio) GetFuenteAudio() string {
	return m.fuenteAudio
}

func (m *MetadataAudio) SetFuenteAudio(fuenteAudio string) {
	m.fuenteAudio = fuenteAudio
}

func (m *MetadataAudio) GetUsoSugerido() string {
	return m.usoSugerido
}

func (m *MetadataAudio) SetUsoSugerido(usoSugerido string) {
	m.usoSugerido = usoSugerido
}

func (m *MetadataAudio) GetProveedorContenido() string {
	return m.proveedorContenido
}

func (m *MetadataAudio) SetProveedorContenido(proveedorContenido string) {
	m.proveedorContenido = proveedorContenido
}

func (m *MetadataAudio) GetDuracionBucle() string {
	return m.duracionBucle
}

func (m *MetadataAudio) SetDuracionBucle(duracionBucle string) {
	m.duracionBucle = duracionBucle
}

func (m *MetadataAudio) GetFrecuenciaDominante() string {
	return m.frecuenciaDominante
}

func (m *MetadataAudio) SetFrecuenciaDominante(frecuenciaDominante string) {
	m.frecuenciaDominante = frecuenciaDominante
}
