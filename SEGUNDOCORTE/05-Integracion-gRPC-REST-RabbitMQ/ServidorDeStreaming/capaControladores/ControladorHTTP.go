package capaControladores

import (
	"io"
	"net/http"

	"ServidorDeStreaming/capaFachada"
	pb "ServidorDeStreaming/protos"
)

const PuertoHTTP = ":8091"

type ControladorHTTP struct{}

func NuevoControladorHTTP() *ControladorHTTP {
	return &ControladorHTTP{}
}

func (c *ControladorHTTP) ObtenerPuertoHTTP() string {
	return PuertoHTTP
}

func (c *ControladorHTTP) IniciarServidorHTTP() {
	http.HandleFunc("/almacenar", c.ManejarAlmacenamiento)
}

func (c *ControladorHTTP) ManejarAlmacenamiento(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	filename := r.Header.Get("X-Filename")
	if filename == "" {
		filename = "audio_subido.mp3"
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
		return
	}

	var respuesta pb.AudioFileResponse
	err = capaFachada.GuardarAudio(filename, data, &respuesta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if respuesta.Success {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, respuesta.Message, http.StatusInternalServerError)
	}
}