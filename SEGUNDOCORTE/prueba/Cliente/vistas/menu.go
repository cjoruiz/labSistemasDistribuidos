package vistas

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	util "Cliente/utilidades"
	pb "ServidorDeAudios/protos"
	pbStreaming "ServidorDeStreaming/protos"
)

var clienteMetadata pb.MetadataServiceClient
var clienteStreaming pbStreaming.StreamingServiceClient
var ctx context.Context
var usuarioActual string
var estaLogueado bool

func InicializarClientes(metadataClient pb.MetadataServiceClient, streamingClient pbStreaming.StreamingServiceClient, context context.Context) {
	clienteMetadata = metadataClient
	clienteStreaming = streamingClient
	ctx = context
}

func MostrarMenuPrincipal() {
	if !estaLogueado {
		MostrarLogin()
		return
	}

	util.LimpiarPantalla()
	fmt.Println("----------------------------------------")
	fmt.Println("    APLICACIÓN DE AUDIO gRPC            ")
	fmt.Println("----------------------------------------")
	fmt.Printf("  Sesión: %s\n", usuarioActual)
	fmt.Println("----------------------------------------")
	fmt.Println("  1. Ver Tipos de Audio                 ")
	fmt.Println("  2. Cerrar Sesión                       ")
	fmt.Println("  3. Salir                               ")
	fmt.Println("----------------------------------------")
	fmt.Print("\nSeleccione una opción: ")

	reader := bufio.NewReader(os.Stdin)
	opcion, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			// Handle EOF gracefully - exit the application
			fmt.Println("\nEntrada cerrada. Saliendo de la aplicación...")
			os.Exit(0)
		}
		fmt.Printf("Error reading input: %v\n", err)
		fmt.Print("Presione Enter para continuar...")
		reader.ReadString('\n')
		MostrarMenuPrincipal()
		return
	}
	opcion = strings.TrimSpace(opcion)

	switch opcion {
	case "1":
		MostrarTiposAudio()
	case "2":
		estaLogueado = false
		usuarioActual = ""
		MostrarLogin()
	case "3":
		fmt.Println("¡Hasta luego!")
		os.Exit(0)
	default:
		fmt.Println("Opción inválida. Presione Enter para continuar...")
		reader.ReadString('\n')
		MostrarMenuPrincipal()
	}
}

func MostrarLogin() {
	util.LimpiarPantalla()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║         INICIAR SESIÓN                 ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
	fmt.Print("  Nickname: ")
	reader := bufio.NewReader(os.Stdin)
	nickname, _ := reader.ReadString('\n')
	nickname = strings.TrimSpace(nickname)

	fmt.Print("  Contraseña: ")
	contrasena, _ := reader.ReadString('\n')
	contrasena = strings.TrimSpace(contrasena)

	fmt.Println()

	if validarCredenciales(nickname, contrasena) {
		usuarioActual = nickname
		estaLogueado = true
		MostrarMenuPrincipal()
	} else {
		fmt.Println("  ✗ Nickname o contraseña incorrectos")
		fmt.Println("  Presione Enter para intentar de nuevo...")
		reader.ReadString('\n')
		MostrarLogin()
	}
}

func MostrarTiposAudio() {
	util.LimpiarPantalla()
	fmt.Println("----------------------------------------")
	fmt.Println("        TIPOS DE AUDIO                  ")
	fmt.Println("----------------------------------------")

	resp, err := clienteMetadata.GetAudioTypes(ctx, &pb.Empty{})
	if err != nil {
		fmt.Printf("Error al obtener tipos de audio: %v\n", err)
		fmt.Print("\nPresione Enter para continuar...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		MostrarMenuPrincipal()
		return
	}

	fmt.Println()
	for i, tipo := range resp.AudioTypes {
		fmt.Printf("  %d. %s\n", i+1, tipo.Nombre)
	}
	fmt.Printf("  %d. Atrás\n", len(resp.AudioTypes)+1)
	fmt.Println()

	fmt.Print("Seleccione una opción: ")
	reader := bufio.NewReader(os.Stdin)
	opcion, _ := reader.ReadString('\n')
	opcion = strings.TrimSpace(opcion)

	idx := 0
	_, err = fmt.Sscanf(opcion, "%d", &idx)
	if err != nil {
		fmt.Println("Opción inválida.")
		fmt.Print("Presione Enter para continuar...")
		reader.ReadString('\n')
		MostrarTiposAudio()
		return
	}
	if idx == len(resp.AudioTypes)+1 {
		MostrarMenuPrincipal()
		return
	}
	if idx < 1 || idx > len(resp.AudioTypes)+1 {
		fmt.Println("Opción inválida.")
		fmt.Print("Presione Enter para continuar...")
		reader.ReadString('\n')
		MostrarTiposAudio()
		return
	}

	tipoSeleccionado := resp.AudioTypes[idx-1]
	MostrarListaAudios(tipoSeleccionado.Nombre, tipoSeleccionado.Id)
}

func MostrarListaAudios(tipo string, tipoId int32) {
	util.LimpiarPantalla()
	fmt.Println("----------------------------------------")
	fmt.Printf("        AUDIOS TIPO: %s                  \n", tipo)
	fmt.Println("----------------------------------------")

	resp, err := clienteMetadata.GetAudiosByType(ctx, &pb.AudioTypeRequest{AudioTypeId: tipoId})
	if err != nil {
		fmt.Printf("Error al obtener audios: %v\n", err)
		fmt.Print("\nPresione Enter para continuar...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		MostrarTiposAudio()
		return
	}

	if len(resp.Audios) == 0 {
		fmt.Println("No hay audios disponibles de este tipo.")
		fmt.Print("\nPresione Enter para continuar...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		MostrarTiposAudio()
		return
	}

	fmt.Println()
	for i, audio := range resp.Audios {
		fmt.Printf("  %d. %s\n", i+1, audio.Titulo)
	}
	fmt.Printf("  %d. Atrás\n", len(resp.Audios)+1)
	fmt.Println()

	fmt.Print("Seleccione una opción: ")
	reader := bufio.NewReader(os.Stdin)
	opcion, _ := reader.ReadString('\n')
	opcion = strings.TrimSpace(opcion)

	idx := 0
	_, errScan := fmt.Sscanf(opcion, "%d", &idx)
	if errScan != nil {
		fmt.Println("Opción inválida.")
		fmt.Print("Presione Enter para continuar...")
		reader.ReadString('\n')
		MostrarListaAudios(tipo, tipoId)
		return
	}
	if idx == len(resp.Audios)+1 {
		MostrarTiposAudio()
		return
	}
	if idx < 1 || idx > len(resp.Audios)+1 {
		fmt.Println("Opción inválida.")
		fmt.Print("Presione Enter para continuar...")
		reader.ReadString('\n')
		MostrarListaAudios(tipo, tipoId)
		return
	}

	audioSeleccionado := resp.Audios[idx-1]
	MostrarDetallesAudio(audioSeleccionado.Id, tipo, tipoId)
}

func MostrarDetallesAudio(id string, tipo string, tipoId int32) {
	util.LimpiarPantalla()
	fmt.Println("----------------------------------------")
	fmt.Println("        DETALLES DEL AUDIO             ")
	fmt.Println("----------------------------------------")

	resp, err := clienteMetadata.GetAudioDetails(ctx, &pb.AudioDetailsRequest{AudioId: id})
	if err != nil {
		fmt.Printf("Error al obtener detalles: %v\n", err)
		fmt.Print("\nPresione Enter para continuar...")
		bufio.NewReader(os.Stdin).ReadString('\n')
		MostrarListaAudios(tipo, tipoId)
		return
	}

	MostrarMenuDetallesConDatos(resp, tipo, tipoId)
}

func MostrarMenuDetallesConDatos(resp *pb.AudioDetailsResponse, tipo string, tipoId int32) {
	util.LimpiarPantalla()
	fmt.Println("----------------------------------------")
	fmt.Println("        DETALLES DEL AUDIO             ")
	fmt.Println("----------------------------------------")
	fmt.Println()

	switch tipo {
	case "Música":
		fmt.Printf("  Título:              %s\n", resp.Titulo)
		fmt.Printf("  Autor:               %s\n", resp.Autor)
		fmt.Printf("  Álbum:               %s\n", resp.Album)
		fmt.Printf("  Género Musical:      %s\n", resp.Genero)
		fmt.Printf("  Sello Discográfico:  %s\n", resp.SelloDiscografico)
		fmt.Printf("  Año de Lanzamiento: %s\n", resp.FechaLanzamiento)
		fmt.Printf("  Duración:            %d segundos\n", resp.Duracion)
	case "Podcast":
		fmt.Printf("  Título del Episodio:      %s\n", resp.Titulo)
		fmt.Printf("  Nombre del Podcast:       %s\n", resp.NombrePodcast)
		fmt.Printf("  Anfitrión:                %s\n", resp.Autor)
		fmt.Printf("  Temporada/Episodio:      %s\n", resp.NumeroTemporadaEpisodio)
		fmt.Printf("  Clasificación de Contenido: %s\n", resp.ClasificacionContenido)
		fmt.Printf("  Notas del Show:           %s\n", resp.NotasShow)
		fmt.Printf("  Año:                      %s\n", resp.FechaLanzamiento)
	case "Audiolibro":
		fmt.Printf("  Título del Libro:   %s\n", resp.Titulo)
		fmt.Printf("  Autor:              %s\n", resp.Autor)
		fmt.Printf("  Narrador:           %s\n", resp.Narrador)
		fmt.Printf("  Editorial:          %s\n", resp.Editorial)
		fmt.Printf("  ISBN:               %s\n", resp.Isbn)
		fmt.Printf("  Capítulo:           %s\n", resp.Capitulo)
		fmt.Printf("  Año:                %s\n", resp.FechaLanzamiento)
	case "Ruido Blanco":
		fmt.Printf("  Título:             %s\n", resp.Titulo)
		fmt.Printf("  Tipo de Sonido:     %s\n", resp.TipoSonido)
		fmt.Printf("  Fuente del Audio:   %s\n", resp.FuenteAudio)
		fmt.Printf("  Proveedor:          %s\n", resp.ProveedorContenido)
		fmt.Printf("  Uso Sugerido:       %s\n", resp.UsoSugerido)
		fmt.Printf("  Duración del Bucle: %s\n", resp.DuracionBucle)
		fmt.Printf("  Frecuencia Dominante: %s\n", resp.FrecuenciaDominante)
	}

	fmt.Println()
	disponible := "Sí"
	if !resp.Disponible {
		disponible = "No"
	}
	fmt.Printf("  Disponible:       %s\n", disponible)
	fmt.Println()

	fmt.Println("  1. Reproducir")
	fmt.Println("  2. Atrás")
	fmt.Print("\nSeleccione una opción: ")

	reader := bufio.NewReader(os.Stdin)
	opcion, _ := reader.ReadString('\n')
	opcion = strings.TrimSpace(opcion)

	switch opcion {
	case "1":
		if resp.Disponible {
			ReproducirAudio(resp.NombreArchivo)
			fmt.Print("Presione Enter para continuar...")
			bufio.NewReader(os.Stdin).ReadString('\n')
			MostrarMenuDetallesConDatos(resp, tipo, tipoId)
		} else {
			fmt.Println("El audio no está disponible.")
			fmt.Print("Presione Enter para continuar...")
			reader.ReadString('\n')
			MostrarListaAudios(tipo, tipoId)
		}
	case "2":
		MostrarListaAudios(tipo, tipoId)
	default:
		MostrarMenuDetallesConDatos(resp, tipo, tipoId)
	}
}

func ReproducirAudio(nombreArchivo string) {
	util.LimpiarPantalla()
	fmt.Println("----------------------------------------")
	fmt.Println("        REPRODUCIENDO AUDIO            ")
	fmt.Println("----------------------------------------")
	fmt.Println()
	fmt.Println("Presione ENTER para detener y volver al menú...")
	fmt.Println()

	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	canalSincronizacion := make(chan struct{}, 1)
	go util.ReproducirAudio(ctxCancel, clienteStreaming, nombreArchivo, canalSincronizacion)
	util.EsperarFinReproduccion(canalSincronizacion, cancel)
}
