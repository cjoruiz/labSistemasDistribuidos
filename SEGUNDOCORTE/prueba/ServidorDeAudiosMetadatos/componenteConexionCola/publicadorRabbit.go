package componenterabbit

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PublicadorRabbit struct {
	canal       *amqp.Channel
	conexion    *amqp.Connection
	nombreCola  string
	conectado   bool
}

type NotificacionAudio struct {
	ID              string `json:"id"`
	Titulo          string `json:"titulo"`
	Artista         string `json:"artista"`
	Genero          string `json:"genero"`
	FechaRegistro   string `json:"fechaRegistro"`
	FraseMotivadora string `json:"fraseMotivadora"`
}

func NuevoPublicadorRabbit(nombreCola string) (*PublicadorRabbit, error) {
	direccionRabbit := "amqp://admin:1234@localhost:5672/"

	conexion, err := amqp.Dial(direccionRabbit)
	if err != nil {
		return nil, fmt.Errorf("error al conectar con RabbitMQ: %w", err)
	}

	canal, err := conexion.Channel()
	if err != nil {
		conexion.Close()
		return nil, fmt.Errorf("error al abrir canal: %w", err)
	}

	_, err = canal.QueueDeclare(
		nombreCola,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		canal.Close()
		conexion.Close()
		return nil, fmt.Errorf("error al declarar cola: %w", err)
	}

	fmt.Printf("RabbitMQ: Conexión exitosa - Cola '%s' lista\n", nombreCola)

	return &PublicadorRabbit{
		canal:      canal,
		conexion:   conexion,
		nombreCola: nombreCola,
		conectado:  true,
	}, nil
}

func (p *PublicadorRabbit) PublicarNotificacion(notificacion NotificacionAudio) error {
	if !p.conectado || p.canal == nil {
		fmt.Println("RabbitMQ no disponible, saltando publicación")
		return fmt.Errorf("conexión no disponible")
	}

	cuerpo, err := json.Marshal(notificacion)
	if err != nil {
		return fmt.Errorf("error al serializar notificación: %w", err)
	}

	err = p.canal.Publish(
		"",
		p.nombreCola,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        cuerpo,
		},
	)

	if err != nil {
		return fmt.Errorf("error al publicar notificación: %w", err)
	}

	fmt.Println("Notificación enviada a la cola de correos")
	return nil
}

func (p *PublicadorRabbit) EstaConectado() bool {
	return p.conectado
}

func (p *PublicadorRabbit) Cerrar() {
	if p.canal != nil {
		p.canal.Close()
	}
	if p.conexion != nil {
		p.conexion.Close()
	}
	p.conectado = false
}