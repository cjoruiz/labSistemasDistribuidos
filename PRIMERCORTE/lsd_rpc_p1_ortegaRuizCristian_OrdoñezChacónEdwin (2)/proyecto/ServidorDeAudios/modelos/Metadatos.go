// Package modelos define las estructuras de metadatos específicos por tipo de audio.
package modelos

// ─── Música ──────────────────────────────────────────────────────────────────

// MetadataMusica almacena los metadatos de una canción musical.
type MetadataMusica struct {
	artistaPrincipal  string
	album             string
	generoMusical     string
	tituloCancion     string
	selloDiscografico string
	anioLanzamiento   int32
}

func NewMetadataMusica(artista, album, genero, titulo, sello string, anio int32) MetadataMusica {
	return MetadataMusica{
		artistaPrincipal:  artista,
		album:             album,
		generoMusical:     genero,
		tituloCancion:     titulo,
		selloDiscografico: sello,
		anioLanzamiento:   anio,
	}
}

func (m *MetadataMusica) GetArtistaPrincipal() string  { return m.artistaPrincipal }
func (m *MetadataMusica) GetAlbum() string             { return m.album }
func (m *MetadataMusica) GetGeneroMusical() string     { return m.generoMusical }
func (m *MetadataMusica) GetTituloCancion() string     { return m.tituloCancion }
func (m *MetadataMusica) GetSelloDiscografico() string { return m.selloDiscografico }
func (m *MetadataMusica) GetAnioLanzamiento() int32    { return m.anioLanzamiento }

// ─── Podcast ─────────────────────────────────────────────────────────────────

// MetadataPodcast almacena los metadatos de un episodio de podcast.
type MetadataPodcast struct {
	nombrePodcast          string
	tituloEpisodio         string
	host                   string
	temporadaEpisodio      string
	notasShow              string
	clasificacionContenido string
}

func NewMetadataPodcast(nombre, episodio, host, temporada, notas, clasificacion string) MetadataPodcast {
	return MetadataPodcast{
		nombrePodcast:          nombre,
		tituloEpisodio:         episodio,
		host:                   host,
		temporadaEpisodio:      temporada,
		notasShow:              notas,
		clasificacionContenido: clasificacion,
	}
}

func (p *MetadataPodcast) GetNombrePodcast() string          { return p.nombrePodcast }
func (p *MetadataPodcast) GetTituloEpisodio() string         { return p.tituloEpisodio }
func (p *MetadataPodcast) GetHost() string                   { return p.host }
func (p *MetadataPodcast) GetTemporadaEpisodio() string      { return p.temporadaEpisodio }
func (p *MetadataPodcast) GetNotasShow() string              { return p.notasShow }
func (p *MetadataPodcast) GetClasificacionContenido() string { return p.clasificacionContenido }

// ─── Audiolibro ──────────────────────────────────────────────────────────────

// MetadataAudiolibro almacena los metadatos de un audiolibro.
type MetadataAudiolibro struct {
	tituloLibro string
	autor       string
	narrador    string
	editorial   string
	isbn        string
	capitulo    int32
}

func NewMetadataAudiolibro(titulo, autor, narrador, editorial, isbn string, capitulo int32) MetadataAudiolibro {
	return MetadataAudiolibro{
		tituloLibro: titulo,
		autor:       autor,
		narrador:    narrador,
		editorial:   editorial,
		isbn:        isbn,
		capitulo:    capitulo,
	}
}

func (a *MetadataAudiolibro) GetTituloLibro() string { return a.tituloLibro }
func (a *MetadataAudiolibro) GetAutor() string       { return a.autor }
func (a *MetadataAudiolibro) GetNarrador() string    { return a.narrador }
func (a *MetadataAudiolibro) GetEditorial() string   { return a.editorial }
func (a *MetadataAudiolibro) GetIsbn() string        { return a.isbn }
func (a *MetadataAudiolibro) GetCapitulo() int32     { return a.capitulo }

// ─── Ruido Blanco ────────────────────────────────────────────────────────────

// MetadataRuidoBlanco almacena los metadatos de un audio de ruido blanco.
type MetadataRuidoBlanco struct {
	tipoSonido          string
	fuenteAudio         string
	usoSugerido         string
	proveedorContenido  string
	duracionBucle       string
	frecuenciaDominante string
}

func NewMetadataRuidoBlanco(tipo, fuente, uso, proveedor, duracion, frecuencia string) MetadataRuidoBlanco {
	return MetadataRuidoBlanco{
		tipoSonido:          tipo,
		fuenteAudio:         fuente,
		usoSugerido:         uso,
		proveedorContenido:  proveedor,
		duracionBucle:       duracion,
		frecuenciaDominante: frecuencia,
	}
}

func (r *MetadataRuidoBlanco) GetTipoSonido() string          { return r.tipoSonido }
func (r *MetadataRuidoBlanco) GetFuenteAudio() string         { return r.fuenteAudio }
func (r *MetadataRuidoBlanco) GetUsoSugerido() string         { return r.usoSugerido }
func (r *MetadataRuidoBlanco) GetProveedorContenido() string  { return r.proveedorContenido }
func (r *MetadataRuidoBlanco) GetDuracionBucle() string       { return r.duracionBucle }
func (r *MetadataRuidoBlanco) GetFrecuenciaDominante() string { return r.frecuenciaDominante }
