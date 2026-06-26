package main

import (
	"fmt"

	"ServidorAudiosMetadatos/modelos"
)

var vectorMetadataAudios []modelos.MetadataAudio

func inicializarDatos() {
	vectorMetadataAudios = cargarMetadataAudios()
}

func cargarMetadataAudios() []modelos.MetadataAudio {
	var vec []modelos.MetadataAudio
	var objAudio1, objAudio2, objAudio3, objAudio4, objAudio5, objAudio6, objAudio7, objAudio8 modelos.MetadataAudio

	objAudio1.SetTitulo("Virgen")
	objAudio1.SetAutor("Adolescentes Orquesta")
	objAudio1.SetAlbum("Persona Ideal")
	objAudio1.SetGenero("Salsa")
	objAudio1.SetDuracion(294)
	objAudio1.SetTipo("Musica")
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
	objAudio2.SetTipo("Musica")
	objAudio2.SetTipoId(1)
	objAudio2.SetFechaLanzamiento("2020")
	objAudio2.SetDisponible(true)
	objAudio2.SetNombreArchivo("el_beneficio_de_la_duda_grupo_firme.mp3")
	objAudio2.SetSelloDiscografico("Music VIP Entertainment")

	objAudio3.SetTitulo("Abrazar HABITOS SALUDABLES")
	objAudio3.SetAutor("Molo Cebrian")
	objAudio3.SetAlbum("Entiende Tu Mente")
	objAudio3.SetGenero("Psicologia")
	objAudio3.SetDuracion(1083)
	objAudio3.SetTipo("Podcast")
	objAudio3.SetTipoId(2)
	objAudio3.SetFechaLanzamiento("2025")
	objAudio3.SetDisponible(true)
	objAudio3.SetNombreArchivo("etm_habitos_saludables.mp3")
	objAudio3.SetNombrePodcast("Entiende Tu Mente")
	objAudio3.SetNumeroTemporadaEpisodio("E385")
	objAudio3.SetNotasShow("Episodio sobre como transformar el bienestar en un disfrute")
	objAudio3.SetClasificacionContenido("Educativo")

	objAudio4.SetTitulo("Que pasaria si la tierra fuera plana de verdad?")
	objAudio4.SetAutor("Jordi Wild & Javi Santaolalla")
	objAudio4.SetAlbum("The Wild Project Clips")
	objAudio4.SetGenero("Ciencia")
	objAudio4.SetDuracion(60)
	objAudio4.SetTipo("Podcast")
	objAudio4.SetTipoId(2)
	objAudio4.SetFechaLanzamiento("2023")
	objAudio4.SetDisponible(true)
	objAudio4.SetNombreArchivo("wild_project_tierra_plana.mp3")
	objAudio4.SetNombrePodcast("The Wild Project")
	objAudio4.SetNumeroTemporadaEpisodio("E135")
	objAudio4.SetNotasShow("Analisis fisico y cientifico sobre las imposibilidades")
	objAudio4.SetClasificacionContenido("General")

	objAudio5.SetTitulo("Nature - Gentle Rain 07")
	objAudio5.SetAutor("Pixabay")
	objAudio5.SetAlbum("Natural Ambiences")
	objAudio5.SetGenero("Ruido Blanco")
	objAudio5.SetDuracion(51)
	objAudio5.SetTipo("Ruido Blanco")
	objAudio5.SetTipoId(4)
	objAudio5.SetFechaLanzamiento("2023")
	objAudio5.SetDisponible(true)
	objAudio5.SetNombreArchivo("nature-gentle-rain-07-437321.mp3")
	objAudio5.SetTipoDeSonido("Lluvia Suave")
	objAudio5.SetFuenteAudio("Grabacion de campo")
	objAudio5.SetUsoSugerido("Relajacion y aislamiento acustico")
	objAudio5.SetProveedorContenido("Pixabay Content License")
	objAudio5.SetDuracionBucle("51 segundos")
	objAudio5.SetFrecuenciaDominante("1 kHz - 6 kHz")

	objAudio6.SetTitulo("City 12 - Ambiente Cafeteria")
	objAudio6.SetAutor("Pixabay")
	objAudio6.SetAlbum("Urban Ambiances")
	objAudio6.SetGenero("Ruido Blanco")
	objAudio6.SetDuracion(53)
	objAudio6.SetTipo("Ruido Blanco")
	objAudio6.SetTipoId(4)
	objAudio6.SetFechaLanzamiento("2022")
	objAudio6.SetDisponible(true)
	objAudio6.SetNombreArchivo("city-12-ambiente-cafeteria-33701.mp3")
	objAudio6.SetTipoDeSonido("Ambiente Urbano / Murmullo")
	objAudio6.SetFuenteAudio("Grabacion de campo")
	objAudio6.SetUsoSugerido("Concentracion y estudio")
	objAudio6.SetProveedorContenido("Pixabay Content License")
	objAudio6.SetDuracionBucle("53 segundos")
	objAudio6.SetFrecuenciaDominante("250 Hz - 5 kHz")

	objAudio7.SetTitulo("El Corazon Delator")
	objAudio7.SetAutor("Edgar Allan Poe")
	objAudio7.SetAlbum("Cuentos de Terror")
	objAudio7.SetGenero("Clasico")
	objAudio7.SetDuracion(1058)
	objAudio7.SetTipo("Audiolibro")
	objAudio7.SetTipoId(3)
	objAudio7.SetFechaLanzamiento("1843")
	objAudio7.SetDisponible(true)
	objAudio7.SetNombreArchivo("poe_corazon_delator.mp3")
	objAudio7.SetNarrador("La voz que te cuenta")
	objAudio7.SetEditorial("Dominio Publico")
	objAudio7.SetIsbn("978-1503290563")
	objAudio7.SetCapitulo("Cuento Unico")

	objAudio8.SetTitulo("El Ruisenor y la Rosa")
	objAudio8.SetAutor("Oscar Wilde")
	objAudio8.SetAlbum("Cuentos de Hadas")
	objAudio8.SetGenero("Clasico")
	objAudio8.SetDuracion(1048)
	objAudio8.SetTipo("Audiolibro")
	objAudio8.SetTipoId(3)
	objAudio8.SetFechaLanzamiento("1888")
	objAudio8.SetDisponible(true)
	objAudio8.SetNombreArchivo("wilde_ruisenor_rosa.mp3")
	objAudio8.SetNarrador("La voz que te cuenta")
	objAudio8.SetEditorial("David Nutt")
	objAudio8.SetIsbn("978-1539404281")
	objAudio8.SetCapitulo("Cuento Unico")

	vec = append(vec, objAudio1, objAudio2, objAudio3, objAudio4, objAudio5, objAudio6, objAudio7, objAudio8)
	return vec
}

func obtenerTiposAudio() []TipoAudio {
	return []TipoAudio{
		{Id: 1, Nombre: "Musica"},
		{Id: 2, Nombre: "Podcast"},
		{Id: 3, Nombre: "Audiolibro"},
		{Id: 4, Nombre: "Ruido Blanco"},
	}
}

func obtenerNombreTipo(tipoId int) string {
	switch tipoId {
	case 1:
		return "Musica"
	case 2:
		return "Podcast"
	case 3:
		return "Audiolibro"
	case 4:
		return "Ruido Blanco"
	}
	return ""
}

func obtenerAudiosPorTipoId(tipoId int) ([]modelos.MetadataAudio, []int) {
	nombreTipo := obtenerNombreTipo(tipoId)
	var result []modelos.MetadataAudio
	var indices []int
	for i, audio := range vectorMetadataAudios {
		if audio.GetTipo() == nombreTipo {
			result = append(result, audio)
			indices = append(indices, i)
		}
	}
	return result, indices
}

func obtenerAudioPorId(id string) *modelos.MetadataAudio {
	for i, audio := range vectorMetadataAudios {
		idx := fmt.Sprintf("audio_%d", i+1)
		if idx == id || audio.GetTitulo() == id || audio.GetNombreArchivo() == id {
			return &vectorMetadataAudios[i]
		}
	}
	return nil
}

func encontrarIdEnVectorCompleto(nombreArchivo string) string {
	for i, a := range vectorMetadataAudios {
		if a.GetNombreArchivo() == nombreArchivo {
			return fmt.Sprintf("audio_%d", i+1)
		}
	}
	return "audio_0"
}

func agregarAudio(metadata modelos.MetadataAudio) {
	vectorMetadataAudios = append(vectorMetadataAudios, metadata)
	fmt.Printf("Audio agregado: %s (total: %d)\n", metadata.GetTitulo(), len(vectorMetadataAudios))
}

type TipoAudio struct {
	Id     int32  `json:"id"`
	Nombre string `json:"nombre"`
}

type AudioItem struct {
	Id       string `json:"id"`
	Titulo   string `json:"titulo"`
	Archivo  string `json:"archivo"`
	Formato  string `json:"formato"`
	TipoId   int    `json:"tipoId"`
	TipoNombre string `json:"tipoNombre"`
}
