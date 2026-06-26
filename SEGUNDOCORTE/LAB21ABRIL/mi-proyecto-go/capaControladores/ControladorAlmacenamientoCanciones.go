package capacontroladores

import (
	"fmt"
	"io"
	"net/http"

	capafachadaservices "ServidorQueAlmacenaCanciones/capaFachadaServices/fachada"
	dtos "ServidorQueAlmacenaCanciones/capaFachadaServices/DTOs"
)

type ControladorAlmacenamientoCanciones struct {
	fachada capafachadaservices.FachadaAlmacenamiento
}

func NuevoControladorAlmacenamientoCanciones(fachada capafachadaservices.FachadaAlmacenamiento) *ControladorAlmacenamientoCanciones {
	return &ControladorAlmacenamientoCanciones{
		fachada: fachada,
	}
}

func (thisC *ControladorAlmacenamientoCanciones) AlmacenarAudioCancion(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Almacenado canción...\n")
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(50 << 20) // Límite de 50MB para el archivo
	file, _, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, "Error leyendo el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, _ := io.ReadAll(file)

	//leer los campos del DTO
	dto := dtos.CancionAlmacenarDTOInput{
		Titulo: r.FormValue("titulo"),
		Artista: r.FormValue("artista"),
		Genero: r.FormValue("genero"),
	}

	thisC.fachada.GuardarCancion(dto, data)
}
