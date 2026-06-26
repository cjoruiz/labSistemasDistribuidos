package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"ServidorAudiosMetadatos/modelos"
)

const (
	PuertoHTTP = ":8090"
	DirAudios  = "../audio_files"
)

func main() {
	inicializarDatos()

	http.HandleFunc("/tipos", corsMiddleware(manejarTipos))
	http.HandleFunc("/audios", corsMiddleware(manejarAudios))
	http.HandleFunc("/audio/", corsMiddleware(manejarAudioDetalle))
	http.HandleFunc("/registrar", corsMiddleware(manejarRegistrar))

	fmt.Println("===============================================")
	fmt.Println("  SERVIDOR DE AUDIOS Y METADATOS")
	fmt.Println("===============================================")
	fmt.Printf("Puerto REST: %s\n", PuertoHTTP)
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /tipos           - Listar tipos de audio")
	fmt.Println("  GET  /audios?tipo=X   - Listar audios por tipo")
	fmt.Println("  GET  /audio/{id}      - Detalles de un audio")
	fmt.Println("  POST /registrar       - Registrar nuevo audio (multipart)")
	fmt.Println("===============================================")

	log.Fatal(http.ListenAndServe(PuertoHTTP, nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func manejarTipos(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	tipos := obtenerTiposAudio()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos)
}

func manejarAudios(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	tipoStr := r.URL.Query().Get("tipo")
	tipoId, err := strconv.Atoi(tipoStr)
	if err != nil {
		http.Error(w, "Parametro 'tipo' requerido", http.StatusBadRequest)
		return
	}

	audios, indices := obtenerAudiosPorTipoId(tipoId)
	var items []AudioItem
	for idx, a := range audios {
		ext := filepath.Ext(a.GetNombreArchivo())
		formato := strings.TrimPrefix(ext, ".")
		nameNoExt := strings.TrimSuffix(a.GetNombreArchivo(), ext)
		id := fmt.Sprintf("audio_%d", indices[idx]+1)
		items = append(items, AudioItem{
			Id:         id,
			Titulo:     a.GetTitulo(),
			Archivo:    nameNoExt,
			Formato:    formato,
			TipoId:     a.GetTipoId(),
			TipoNombre: a.GetTipo(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func manejarAudioDetalle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/audio/")
	audio := obtenerAudioPorId(id)
	if audio == nil {
		http.Error(w, "Audio no encontrado", http.StatusNotFound)
		return
	}

	ext := filepath.Ext(audio.GetNombreArchivo())
	formato := strings.TrimPrefix(ext, ".")

	resp := map[string]interface{}{
		"titulo":            audio.GetTitulo(),
		"autor":             audio.GetAutor(),
		"album":             audio.GetAlbum(),
		"genero":            audio.GetGenero(),
		"duracion":          audio.GetDuracion(),
		"tipo":              audio.GetTipo(),
		"tipoId":            audio.GetTipoId(),
		"archivo":           strings.TrimSuffix(audio.GetNombreArchivo(), ext),
		"formato":           formato,
		"fechaLanzamiento": audio.GetFechaLanzamiento(),
		"disponible":        audio.GetDisponible(),
		"selloDiscografico": audio.GetSelloDiscografico(),
		"nombrePodcast":     audio.GetNombrePodcast(),
		"numeroEpisodio":    audio.GetNumeroTemporadaEpisodio(),
		"notasShow":         audio.GetNotasShow(),
		"clasificacion":     audio.GetClasificacionContenido(),
		"narrador":          audio.GetNarrador(),
		"editorial":         audio.GetEditorial(),
		"isbn":              audio.GetIsbn(),
		"capitulo":          audio.GetCapitulo(),
		"tipoSonido":        audio.GetTipoDeSonido(),
		"fuenteAudio":       audio.GetFuenteAudio(),
		"usoSugerido":       audio.GetUsoSugerido(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

type RegistroRequest struct {
	Titulo       string `json:"titulo"`
	Autor        string `json:"autor"`
	Album        string `json:"album"`
	Genero       string `json:"genero"`
	TipoId       int    `json:"tipoId"`
	TipoNombre   string `json:"tipoNombre"`
	SelloDiscografico      string `json:"selloDiscografico"`
	NombrePodcast          string `json:"nombrePodcast"`
	NumeroEpisodio         string `json:"numeroEpisodio"`
	NotasShow              string `json:"notasShow"`
	Clasificacion          string `json:"clasificacion"`
	Narrador               string `json:"narrador"`
	Editorial              string `json:"editorial"`
	Isbn                   string `json:"isbn"`
	Capitulo               string `json:"capitulo"`
	TipoSonido             string `json:"tipoSonido"`
	FuenteAudio            string `json:"fuenteAudio"`
	UsoSugerido            string `json:"usoSugerido"`
}

func manejarRegistrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "multipart/form-data") {
		err := r.ParseMultipartForm(50 << 20)
		if err != nil {
			http.Error(w, "Error al parsear multipart", http.StatusBadRequest)
			return
		}

		metadataJSON := r.FormValue("metadata")
		if metadataJSON == "" {
			http.Error(w, "Campo 'metadata' requerido", http.StatusBadRequest)
			return
		}

		var req RegistroRequest
		if err := json.Unmarshal([]byte(metadataJSON), &req); err != nil {
			http.Error(w, "Error al parsear metadata JSON", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("audio")
		if err != nil {
			http.Error(w, "Campo 'audio' requerido", http.StatusBadRequest)
			return
		}
		defer file.Close()

		if err := os.MkdirAll(DirAudios, 0755); err != nil {
			http.Error(w, "Error al crear directorio", http.StatusInternalServerError)
			return
		}

		filename := header.Filename
		dst, err := os.Create(filepath.Join(DirAudios, filename))
		if err != nil {
			http.Error(w, "Error al guardar archivo", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Error al escribir archivo", http.StatusInternalServerError)
			return
		}

		ext := filepath.Ext(filename)
		registrarMetadata(req, strings.TrimSuffix(filename, ext), ext)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true, "message": "Audio registrado exitosamente",
		})
	} else {
		http.Error(w, "Content-Type debe ser multipart/form-data", http.StatusBadRequest)
	}
}

func registrarMetadata(req RegistroRequest, nombreSinExt string, ext string) {
	audio := modelos.MetadataAudio{}
	audio.SetTitulo(req.Titulo)
	audio.SetAutor(req.Autor)
	audio.SetAlbum(req.Album)
	audio.SetGenero(req.Genero)
	audio.SetTipo(req.TipoNombre)
	audio.SetTipoId(req.TipoId)
	audio.SetDisponible(true)
	audio.SetNombreArchivo(nombreSinExt + ext)
	audio.SetSelloDiscografico(req.SelloDiscografico)
	audio.SetNombrePodcast(req.NombrePodcast)
	audio.SetNumeroTemporadaEpisodio(req.NumeroEpisodio)
	audio.SetNotasShow(req.NotasShow)
	audio.SetClasificacionContenido(req.Clasificacion)
	audio.SetNarrador(req.Narrador)
	audio.SetEditorial(req.Editorial)
	audio.SetIsbn(req.Isbn)
	audio.SetCapitulo(req.Capitulo)
	audio.SetTipoDeSonido(req.TipoSonido)
	audio.SetFuenteAudio(req.FuenteAudio)
	audio.SetUsoSugerido(req.UsoSugerido)
	agregarAudio(audio)
}
