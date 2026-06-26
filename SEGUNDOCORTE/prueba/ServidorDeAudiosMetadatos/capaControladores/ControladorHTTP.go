package capaControladores

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"ServidorDeAudios/capaAccesoDatos"
	"ServidorDeAudios/componenteConexionCola"
	"ServidorDeAudios/modelos"
)

type MetadatosRequest struct {
	ID                      string `json:"id"`
	Titulo                  string `json:"titulo"`
	Autor                   string `json:"autor"`
	Album                   string `json:"album"`
	Genero                  string `json:"genero"`
	Duracion                int    `json:"duracion"`
	TipoId                  int    `json:"tipoId"`
	TipoNombre              string `json:"tipoNombre"`
	FechaLanzamiento        string `json:"fechaLanzamiento"`
	Disponible              bool   `json:"disponible"`
	NombreArchivo          string `json:"nombreArchivo"`
	SelloDiscografico       string `json:"selloDiscografico"`
	NombrePodcast           string `json:"nombrePodcast"`
	NumeroTemporadaEpisodio string `json:"numeroTemporadaEpisodio"`
	NotasShow               string `json:"notasShow"`
	ClasificacionContenido  string `json:"clasificacionContenido"`
	Narrador                string `json:"narrador"`
	Editorial               string `json:"editorial"`
	Isbn                    string `json:"isbn"`
	Capitulo                string `json:"capitulo"`
	TipoSonido              string `json:"tipoSonido"`
	FuenteAudio             string `json:"fuenteAudio"`
	UsoSugerido             string `json:"usoSugerido"`
	ProveedorContenido     string `json:"proveedorContenido"`
	DuracionBucle           string `json:"duracionBucle"`
	FrecuenciaDominante     string `json:"frecuenciaDominante"`
}

type ControladorHTTP struct {
	publicador *componenterabbit.PublicadorRabbit
}

func NuevoControladorHTTP(publicador *componenterabbit.PublicadorRabbit) *ControladorHTTP {
	return &ControladorHTTP{
		publicador: publicador,
	}
}

func (c *ControladorHTTP) ManejarMetadatos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.procesarAlmacenamiento(w, r)
	case http.MethodGet:
		c.procesarListado(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (c *ControladorHTTP) procesarAlmacenamiento(w http.ResponseWriter, r *http.Request) {
	cuerpo, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer cuerpo", http.StatusBadRequest)
		return
	}

	var metadatos MetadatosRequest
	err = json.Unmarshal(cuerpo, &metadatos)
	if err != nil {
		http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Recibiendo metadatos: %s\n", metadatos.Titulo)

	audioNuevo := construirAudio(metadatos)
	capaaccesodatos.AgregarAudio(audioNuevo)

	c.publicarNotificacion(metadatos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"success": true, "message": "Metadatos almacenados"}`))
}

func (c *ControladorHTTP) procesarListado(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"audios": []}`))
}

func (c *ControladorHTTP) publicarNotificacion(metadatos MetadatosRequest) {
	fechaRegistro := time.Now().Format("2006-01-02 15:04:05")
	fraseMotivadora := generarFraseMotivadora()

	notificacion := componenterabbit.NotificacionAudio{
		ID:              metadatos.ID,
		Titulo:          metadatos.Titulo,
		Artista:         metadatos.Autor,
		Genero:          metadatos.Genero,
		FechaRegistro:   fechaRegistro,
		FraseMotivadora: fraseMotivadora,
	}

	if c.publicador != nil && c.publicador.EstaConectado() {
		err := c.publicador.PublicarNotificacion(notificacion)
		if err != nil {
			fmt.Printf("Error al publicar notificación: %v\n", err)
		}
	} else {
		fmt.Println("RabbitMQ no disponible, saltando publicación")
	}
}

func construirAudio(metadatos MetadatosRequest) modelos.MetadataAudio {
	audio := modelos.MetadataAudio{}
	audio.SetTitulo(metadatos.Titulo)
	audio.SetAutor(metadatos.Autor)
	audio.SetAlbum(metadatos.Album)
	audio.SetGenero(metadatos.Genero)
	audio.SetDuracion(metadatos.Duracion)
	audio.SetTipo(metadatos.TipoNombre)
	audio.SetTipoId(metadatos.TipoId)
	audio.SetFechaLanzamiento(metadatos.FechaLanzamiento)
	audio.SetDisponible(metadatos.Disponible)
	audio.SetNombreArchivo(metadatos.NombreArchivo)
	audio.SetSelloDiscografico(metadatos.SelloDiscografico)
	audio.SetNombrePodcast(metadatos.NombrePodcast)
	audio.SetNumeroTemporadaEpisodio(metadatos.NumeroTemporadaEpisodio)
	audio.SetNotasShow(metadatos.NotasShow)
	audio.SetClasificacionContenido(metadatos.ClasificacionContenido)
	audio.SetNarrador(metadatos.Narrador)
	audio.SetEditorial(metadatos.Editorial)
	audio.SetIsbn(metadatos.Isbn)
	audio.SetCapitulo(metadatos.Capitulo)
	audio.SetTipoDeSonido(metadatos.TipoSonido)
	audio.SetFuenteAudio(metadatos.FuenteAudio)
	audio.SetUsoSugerido(metadatos.UsoSugerido)
	audio.SetProveedorContenido(metadatos.ProveedorContenido)
	audio.SetDuracionBucle(metadatos.DuracionBucle)
	audio.SetFrecuenciaDominante(metadatos.FrecuenciaDominante)
	return audio
}

func generarFraseMotivadora() string {
	frases := []string{
		"La música es el lenguaje de los pulmones.",
		"Cada canción es una aventura musical.",
		"El ritmo de la vida está en la música.",
		"Deja que la melodía guíe tu alma.",
		"La música conecta corazones.",
		"Un nuevo audio, una nueva historia por contar.",
		"Tu biblioteca musical crece con cada nota.",
	}
	return frases[time.Now().Second()%len(frases)]
}