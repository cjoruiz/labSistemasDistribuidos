package capaaccesodatos

import (
	. "ServidorDeAudios/modelos"
	"fmt"
	"os"
)

type AudioType struct {
	Id     int32
	Nombre string
}

var vectorMetadataAudios []MetadataAudio

func InicializarDatos() {
	vectorMetadataAudios = CargarMetadataAudios()
}

func CargarMetadataAudios() []MetadataAudio {
	var vec []MetadataAudio
	var objAudio1, objAudio2, objAudio3, objAudio4, objAudio5, objAudio6, objAudio7, objAudio8 MetadataAudio

	objAudio1.SetTitulo("Virgen")
	objAudio1.SetAutor("Adolescentes Orquesta")
	objAudio1.SetAlbum("Persona Ideal")
	objAudio1.SetGenero("Salsa")
	objAudio1.SetDuracion(294)
	objAudio1.SetTipo("Música")
	objAudio1.SetTipoId(1)
	objAudio1.SetFechaLanzamiento("1996")
	objAudio1.SetDisponible(true)
	objAudio1.SetNombreArchivo("virgen_adolescentes_orquesta.mp3")
	objAudio1.SetSelloDiscografico("Korta Record's")

	objAudio2.SetTitulo("El Beneficio de la Duda")
	objAudio2.SetAutor("Grupo Firme")
	objAudio2.SetAlbum("Nos Divertimos Logrando Lo Imposible")
	objAudio2.SetGenero("Regional Mexicano")
	objAudio2.SetDuracion(192)
	objAudio2.SetTipo("Música")
	objAudio2.SetTipoId(1)
	objAudio2.SetFechaLanzamiento("2020")
	objAudio2.SetDisponible(true)
	objAudio2.SetNombreArchivo("el_beneficio_de_la_duda_grupo_firme.mp3")
	objAudio2.SetSelloDiscografico("Music VIP Entertainment")

	objAudio3.SetTitulo("Abrazar HÁBITOS SALUDABLES")
	objAudio3.SetAutor("Molo Cebrián")
	objAudio3.SetAlbum("Entiende Tu Mente")
	objAudio3.SetGenero("Psicología")
	objAudio3.SetDuracion(1083) // 18:03 minutos
	objAudio3.SetTipo("Podcast")
	objAudio3.SetTipoId(2)
	objAudio3.SetFechaLanzamiento("2025")
	objAudio3.SetDisponible(true)
	objAudio3.SetNombreArchivo("etm_habitos_saludables.mp3")
	objAudio3.SetNombrePodcast("Entiende Tu Mente")
	objAudio3.SetNumeroTemporadaEpisodio("E385")
	objAudio3.SetNotasShow("Episodio sobre cómo transformar el bienestar en un disfrute en lugar de una obligación.")
	objAudio3.SetClasificacionContenido("Educativo")

	objAudio4.SetTitulo("¿Qué pasaría si la tierra fuera plana de verdad?")
	objAudio4.SetAutor("Jordi Wild & Javi Santaolalla")
	objAudio4.SetAlbum("The Wild Project Clips")
	objAudio4.SetGenero("Ciencia")
	objAudio4.SetDuracion(60) // 1:00 minuto
	objAudio4.SetTipo("Podcast")
	objAudio4.SetTipoId(2)
	objAudio4.SetFechaLanzamiento("2023")
	objAudio4.SetDisponible(true)
	objAudio4.SetNombreArchivo("wild_project_tierra_plana.mp3")
	objAudio4.SetNombrePodcast("The Wild Project")
	objAudio4.SetNumeroTemporadaEpisodio("E135")
	objAudio4.SetNotasShow("Análisis físico y científico sobre las imposibilidades de una tierra plana.")
	objAudio4.SetClasificacionContenido("General")

	objAudio5.SetTitulo("Nature - Gentle Rain 07")
	objAudio5.SetAutor("Pixabay")
	objAudio5.SetAlbum("Natural Ambiences")
	objAudio5.SetGenero("Ruido Blanco")
	objAudio5.SetDuracion(51) // 0:51 segundos
	objAudio5.SetTipo("Ruido Blanco")
	objAudio5.SetFechaLanzamiento("2023")
	objAudio5.SetDisponible(true)
	objAudio5.SetNombreArchivo("nature-gentle-rain-07-437321.mp3")
	objAudio5.SetTipoDeSonido("Lluvia Suave")
	objAudio5.SetFuenteAudio("Grabación de campo")
	objAudio5.SetUsoSugerido("Relajación y aislamiento acústico")
	objAudio5.SetProveedorContenido("Pixabay Content License")
	objAudio5.SetDuracionBucle("51 segundos")
	objAudio5.SetFrecuenciaDominante("1 kHz - 6 kHz")

	objAudio7.SetTitulo("City 12 - Ambiente Cafeteria")
	objAudio7.SetAutor("Pixabay")
	objAudio7.SetAlbum("Urban Ambiances")
	objAudio7.SetGenero("Ambiente / Ruido Blanco")
	objAudio7.SetDuracion(53) // 0:53 segundos
	objAudio7.SetTipo("Ruido Blanco")
	objAudio7.SetFechaLanzamiento("2022")
	objAudio7.SetDisponible(true)
	objAudio7.SetNombreArchivo("city-12-ambiente-cafeteria-33701.mp3")
	objAudio7.SetTipoDeSonido("Ambiente Urbano / Murmullo")
	objAudio7.SetFuenteAudio("Grabación de campo")
	objAudio7.SetUsoSugerido("Concentración y estudio (Efecto Coffee Shop)")
	objAudio7.SetProveedorContenido("Pixabay Content License")
	objAudio7.SetDuracionBucle("53 segundos")
	objAudio7.SetFrecuenciaDominante("250 Hz - 5 kHz")

	objAudio6.SetTitulo("El Corazón Delator")
	objAudio6.SetAutor("Edgar Allan Poe")
	objAudio6.SetAlbum("Cuentos de Terror")
	objAudio6.SetGenero("Clásico")
	objAudio6.SetDuracion(1058)
	objAudio6.SetTipo("Audiolibro")
	objAudio6.SetTipoId(3)
	objAudio6.SetFechaLanzamiento("1843")
	objAudio6.SetDisponible(true)
	objAudio6.SetNombreArchivo("poe_corazon_delator.mp3")
	objAudio6.SetNarrador("La voz que te cuenta")
	objAudio6.SetEditorial("Dominio Público")
	objAudio6.SetIsbn("978-1503290563")
	objAudio6.SetCapitulo("Cuento Único")

	objAudio8.SetTitulo("El Ruiseñor y la Rosa")
	objAudio8.SetAutor("Oscar Wilde")
	objAudio8.SetAlbum("Cuentos de Hadas")
	objAudio8.SetGenero("Clásico")
	objAudio8.SetDuracion(1048)
	objAudio8.SetTipo("Audiolibro")
	objAudio8.SetTipoId(3)
	objAudio8.SetFechaLanzamiento("1888")
	objAudio8.SetDisponible(true)
	objAudio8.SetNombreArchivo("wilde_ruisenor_rosa.mp3")
	objAudio8.SetNarrador("La voz que te cuenta")
	objAudio8.SetEditorial("David Nutt")
	objAudio8.SetIsbn("978-1539404281")
	objAudio8.SetCapitulo("Cuento Único")

	vec = append(vec, objAudio1)
	vec = append(vec, objAudio2)
	vec = append(vec, objAudio3)
	vec = append(vec, objAudio4)
	vec = append(vec, objAudio5)
	vec = append(vec, objAudio6)
	vec = append(vec, objAudio7)
	vec = append(vec, objAudio8)

	return vec
}

func ObtenerVectorMetadataAudios() []MetadataAudio {
	return vectorMetadataAudios
}

func ObtenerTiposAudio() []AudioType {
	tiposOrdenados := []AudioType{
		{Id: 1, Nombre: "Música"},
		{Id: 2, Nombre: "Podcast"},
		{Id: 3, Nombre: "Audiolibro"},
		{Id: 4, Nombre: "Ruido Blanco"},
	}
	agregados := make(map[string]bool)
	var result []AudioType
	for _, tipo := range tiposOrdenados {
		for _, audio := range vectorMetadataAudios {
			if audio.GetTipo() == tipo.Nombre && !agregados[tipo.Nombre] {
				result = append(result, tipo)
				agregados[tipo.Nombre] = true
				break
			}
		}
	}
	return result
}

func ObtenerAudiosPorTipoId(tipoId int) []MetadataAudio {
	nombreTipo := ""
	switch tipoId {
	case 1:
		nombreTipo = "Música"
	case 2:
		nombreTipo = "Podcast"
	case 3:
		nombreTipo = "Audiolibro"
	case 4:
		nombreTipo = "Ruido Blanco"
	}
	var result []MetadataAudio
	for _, audio := range vectorMetadataAudios {
		if audio.GetTipo() == nombreTipo {
			result = append(result, audio)
		}
	}
	return result
}

func ObtenerAudioPorId(id string) *MetadataAudio {
	for i, audio := range vectorMetadataAudios {
		if fmt.Sprintf("audio_%d", i+1) == id || audio.GetTitulo() == id {
			return &vectorMetadataAudios[i]
		}
	}
	return nil
}

func ObtenerNombreArchivo(id string) string {
	audio := ObtenerAudioPorId(id)
	if audio != nil {
		return audio.GetNombreArchivo()
	}
	return ""
}

func BuscarAudio(titulo string) RespuestaMetadataAudioDTO {
	for i := 0; i < len(vectorMetadataAudios); i++ {
		if vectorMetadataAudios[i].GetTitulo() == titulo {
			var resp RespuestaMetadataAudioDTO
			resp.ObjAudio = vectorMetadataAudios[i]
			resp.Codigo = 200
			resp.Mensaje = "Metadata de audio encontrada"
			return resp
		}
	}
	var resp RespuestaMetadataAudioDTO
	resp.Codigo = 404
	resp.Mensaje = "Metadata de audio no se encontró"
	return resp
}

func AbrirArchivo(titulo string) (*os.File, error) {
	file, err := os.Open(titulo)
	if err != nil {
		fmt.Println("Error Audio abierto:", titulo)
		return nil, fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	fmt.Println("Audio Abierto:", titulo)
	return file, nil
}
