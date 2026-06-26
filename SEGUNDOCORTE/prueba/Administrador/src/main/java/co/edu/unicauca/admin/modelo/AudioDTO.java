package co.edu.unicauca.admin.modelo;

public class AudioDTO {

    private String id;
    private String titulo;
    private String autor;
    private String album;
    private String genero;
    private Integer duracion;
    private Integer tipoId;
    private String tipoNombre;
    private String fechaLanzamiento;
    private Boolean disponible;
    private String nombreArchivo;

    // Musica
    private String selloDiscografico;

    // Podcast
    private String nombrePodcast;
    private String numeroTemporadaEpisodio;
    private String notasShow;
    private String clasificacionContenido;

    // Audiolibro
    private String narrador;
    private String editorial;
    private String isbn;
    private String capitulo;

    // Ruido Blanco
    private String tipoSonido;
    private String fuenteAudio;
    private String usoSugerido;
    private String proveedorContenido;
    private String duracionBucle;
    private String frecuenciaDominante;

    public AudioDTO() {
    }

    // Getters y Setters
    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getTitulo() {
        return titulo;
    }

    public void setTitulo(String titulo) {
        this.titulo = titulo;
    }

    public String getAutor() {
        return autor;
    }

    public void setAutor(String autor) {
        this.autor = autor;
    }

    public String getAlbum() {
        return album;
    }

    public void setAlbum(String album) {
        this.album = album;
    }

    public String getGenero() {
        return genero;
    }

    public void setGenero(String genero) {
        this.genero = genero;
    }

    public Integer getDuracion() {
        return duracion;
    }

    public void setDuracion(Integer duracion) {
        this.duracion = duracion;
    }

    public Integer getTipoId() {
        return tipoId;
    }

    public void setTipoId(Integer tipoId) {
        this.tipoId = tipoId;
    }

    public String getTipoNombre() {
        return tipoNombre;
    }

    public void setTipoNombre(String tipoNombre) {
        this.tipoNombre = tipoNombre;
    }

    public String getFechaLanzamiento() {
        return fechaLanzamiento;
    }

    public void setFechaLanzamiento(String fechaLanzamiento) {
        this.fechaLanzamiento = fechaLanzamiento;
    }

    public Boolean getDisponible() {
        return disponible;
    }

    public void setDisponible(Boolean disponible) {
        this.disponible = disponible;
    }

    public String getNombreArchivo() {
        return nombreArchivo;
    }

    public void setNombreArchivo(String nombreArchivo) {
        this.nombreArchivo = nombreArchivo;
    }

    public String getSelloDiscografico() {
        return selloDiscografico;
    }

    public void setSelloDiscografico(String selloDiscografico) {
        this.selloDiscografico = selloDiscografico;
    }

    public String getNombrePodcast() {
        return nombrePodcast;
    }

    public void setNombrePodcast(String nombrePodcast) {
        this.nombrePodcast = nombrePodcast;
    }

    public String getNumeroTemporadaEpisodio() {
        return numeroTemporadaEpisodio;
    }

    public void setNumeroTemporadaEpisodio(String numeroTemporadaEpisodio) {
        this.numeroTemporadaEpisodio = numeroTemporadaEpisodio;
    }

    public String getNotasShow() {
        return notasShow;
    }

    public void setNotasShow(String notasShow) {
        this.notasShow = notasShow;
    }

    public String getClasificacionContenido() {
        return clasificacionContenido;
    }

    public void setClasificacionContenido(String clasificacionContenido) {
        this.clasificacionContenido = clasificacionContenido;
    }

    public String getNarrador() {
        return narrador;
    }

    public void setNarrador(String narrador) {
        this.narrador = narrador;
    }

    public String getEditorial() {
        return editorial;
    }

    public void setEditorial(String editorial) {
        this.editorial = editorial;
    }

    public String getIsbn() {
        return isbn;
    }

    public void setIsbn(String isbn) {
        this.isbn = isbn;
    }

    public String getCapitulo() {
        return capitulo;
    }

    public void setCapitulo(String capitulo) {
        this.capitulo = capitulo;
    }

    public String getTipoSonido() {
        return tipoSonido;
    }

    public void setTipoSonido(String tipoSonido) {
        this.tipoSonido = tipoSonido;
    }

    public String getFuenteAudio() {
        return fuenteAudio;
    }

    public void setFuenteAudio(String fuenteAudio) {
        this.fuenteAudio = fuenteAudio;
    }

    public String getUsoSugerido() {
        return usoSugerido;
    }

    public void setUsoSugerido(String usoSugerido) {
        this.usoSugerido = usoSugerido;
    }

    public String getProveedorContenido() {
        return proveedorContenido;
    }

    public void setProveedorContenido(String proveedorContenido) {
        this.proveedorContenido = proveedorContenido;
    }

    public String getDuracionBucle() {
        return duracionBucle;
    }

    public void setDuracionBucle(String duracionBucle) {
        this.duracionBucle = duracionBucle;
    }

    public String getFrecuenciaDominante() {
        return frecuenciaDominante;
    }

    public void setFrecuenciaDominante(String frecuenciaDominante) {
        this.frecuenciaDominante = frecuenciaDominante;
    }
}