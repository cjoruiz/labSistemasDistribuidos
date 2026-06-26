// Package capaaccesodatos provee acceso a los datos almacenados en el servidor de metadatos.
// Actúa como repositorio en memoria con los tipos de audio y sus registros.
package capaaccesodatos

import (
	"fmt"
	"servidor.audios.local/grpc-servidor-audios/modelos"
)

// ─── Datos en memoria ────────────────────────────────────────────────────────

var tiposDeAudio []modelos.TipoAudio
var registrosDeAudio []modelos.RegistroAudio

// InicializarDatos carga todos los tipos de audio y registros de ejemplo al arrancar el servidor.
func InicializarDatos() {
	cargarTiposDeAudio()
	cargarRegistrosDeAudio()
	fmt.Println("[Repositorio] Datos de ejemplo cargados correctamente.")
}

// cargarTiposDeAudio define los cuatro tipos soportados por el sistema.
func cargarTiposDeAudio() {
	tiposDeAudio = []modelos.TipoAudio{
		modelos.NewTipoAudio(1, "Música"),
		modelos.NewTipoAudio(2, "Podcasts"),
		modelos.NewTipoAudio(3, "Audiolibros"),
		modelos.NewTipoAudio(4, "Ruido Blanco"),
	}
}

// cargarRegistrosDeAudio define al menos dos audios por tipo.
func cargarRegistrosDeAudio() {
	// ── Música ──────────────────────────────────────────────────────────────
	m1 := modelos.NewMetadataMusica("Shakira", "Laundry Service", "Pop Latino",
		"Whenever, Wherever", "Epic Records", 2001)
	m2 := modelos.NewMetadataMusica("Carlos Vives", "Cumbiana", "Vallenato-Pop",
		"Canción Bonita", "Sony Music", 2020)

	// ── Podcasts ─────────────────────────────────────────────────────────────
	p1 := modelos.NewMetadataPodcast("Entiende Tu Mente", "La procrastinación",
		"Molo Cebrián", "T3E12", "Episodio sobre hábitos y productividad.", "Para toda la familia")
	p2 := modelos.NewMetadataPodcast("Cuentos para Dormir", "El bosque encantado",
		"Ana García", "T1E05", "Historia relajante para conciliar el sueño.", "Para toda la familia")

	// ── Audiolibros ──────────────────────────────────────────────────────────
	a1 := modelos.NewMetadataAudiolibro("Cien años de soledad", "Gabriel García Márquez",
		"Gustavo Bonfigli", "Penguin Random House", "978-0-06-088328-7", 1)
	a2 := modelos.NewMetadataAudiolibro("El principito", "Antoine de Saint-Exupéry",
		"Diego Pierri", "Salamandra", "978-84-9838-072-3", 1)

	// ── Ruido Blanco ─────────────────────────────────────────────────────────
	r1 := modelos.NewMetadataRuidoBlanco("Ruido Blanco", "Lluvia sobre techo",
		"Dormir", "RelaxSounds", "60 min", "Graves")
	r2 := modelos.NewMetadataRuidoBlanco("Ruido Rosa", "Ventilador de techo",
		"Concentración", "FocusAudio", "45 min", "Medios")

	registrosDeAudio = []modelos.RegistroAudio{
		// Música (idTipo=1)
		{Resumen: modelos.NewAudioResumen(1, "Whenever, Wherever - Shakira", 1, "audios/musica1.mp3"), Musica: &m1},
		{Resumen: modelos.NewAudioResumen(2, "Canción Bonita - Carlos Vives", 1, "audios/musica2.mp3"), Musica: &m2},
		// Podcasts (idTipo=2)
		{Resumen: modelos.NewAudioResumen(3, "La procrastinación - Entiende Tu Mente", 2, "audios/podcast1.mp3"), Podcast: &p1},
		{Resumen: modelos.NewAudioResumen(4, "El bosque encantado - Cuentos para Dormir", 2, "audios/podcast2.mp3"), Podcast: &p2},
		// Audiolibros (idTipo=3)
		{Resumen: modelos.NewAudioResumen(5, "Cien años de soledad - Cap. 1", 3, "audios/audiolibro1.mp3"), Audiolibro: &a1},
		{Resumen: modelos.NewAudioResumen(6, "El principito - Cap. 1", 3, "audios/audiolibro2.mp3"), Audiolibro: &a2},
		// Ruido Blanco (idTipo=4)
		{Resumen: modelos.NewAudioResumen(7, "Lluvia sobre techo - Ruido Blanco", 4, "audios/ruidoblanco1.mp3"), RuidoBlanco: &r1},
		{Resumen: modelos.NewAudioResumen(8, "Ventilador de techo - Ruido Rosa", 4, "audios/ruidoblanco2.mp3"), RuidoBlanco: &r2},
	}
}

// ─── Consultas ───────────────────────────────────────────────────────────────

// ObtenerTodos devuelve la lista completa de tipos de audio.
func ObtenerTodos() []modelos.TipoAudio {
	return tiposDeAudio
}

// ObtenerAudiosPorTipo devuelve los audios que pertenecen al tipo indicado por idTipo.
func ObtenerAudiosPorTipo(idTipo int32) []modelos.AudioResumen {
	var resultado []modelos.AudioResumen
	for _, registro := range registrosDeAudio {
		if registro.Resumen.GetIdTipo() == idTipo {
			resultado = append(resultado, registro.Resumen)
		}
	}
	return resultado
}

// ObtenerRegistroPorId devuelve el registro completo de un audio dado su id, o nil si no existe.
func ObtenerRegistroPorId(idAudio int32) *modelos.RegistroAudio {
	for i := range registrosDeAudio {
		if registrosDeAudio[i].Resumen.GetId() == idAudio {
			return &registrosDeAudio[i]
		}
	}
	return nil
}
