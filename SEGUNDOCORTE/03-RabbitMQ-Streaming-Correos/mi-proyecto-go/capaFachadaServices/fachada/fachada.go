package capafachadaservices

import (
	capaaccesoadatos "ServidorQueAlmacenaCanciones/capaAccesoADatos"
	componenteconexioncola "ServidorQueAlmacenaCanciones/componnteConexionCola"
	dtos "ServidorQueAlmacenaCanciones/capaFachadaServices/DTOs"
	"fmt"
)

type FachadaAlmacenamiento struct {
	repo         *capaaccesoadatos.RepositorioCanciones
	conexionCola *componenteconexioncola.RabbitPublisher
}

// Constructor de la fachada
func NuevaFachadaAlmacenamiento() *FachadaAlmacenamiento {
	fmt.Println("Inicializando fachada de almacenamiento...")

	repo := capaaccesoadatos.GetRepositorioCanciones()
	conexionCola, err := componenteconexioncola.NewRabbitPublisher()
	if err != nil {
		fmt.Println("Error al conectar con RabbitMQ: ", err)
		conexionCola = nil
	}

	return &FachadaAlmacenamiento{
		repo:         repo,
		conexionCola: conexionCola,
	}
}

// Invocación a la operación asíncrona mediante cola de mensajes
func (thisF *FachadaAlmacenamiento) GuardarCancion(objCancion dtos.CancionAlmacenarDTOInput, data []byte) error {
	// Publicar notificación asíncrona
		thisF.conexionCola.PublicarNotificacion(componenteconexioncola.NotificacionCancion{
			Titulo:  objCancion.Titulo,
			Artista: objCancion.Artista,
			Genero:  objCancion.Genero,
			Mensaje: "Nueva canción almacenada: " + objCancion.Titulo + " de " + objCancion.Artista,
		})

	// Guardar archivo y registro en memoria 
	//Delegar el repositorio
	return thisF.repo.GuardarCancion(objCancion.Titulo, objCancion.Genero, objCancion.Artista, data)
}