package componenteconexioncola

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue amqp.Queue
}

type NotificacionCancion struct {
	Titulo  string
	Artista string
	Genero  string
	Mensaje string
}

func NewRabbitPublisher() (*RabbitPublisher, error) {
	// Usar la IP de Windows donde corre RabbitMQ
	conn, err := amqp.Dial("amqp://admin:1234@10.252.124.198")
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("error abriendo canal: %v", err)
	}

	// Declarar la cola
	q, err := ch.QueueDeclare(
		"notificaciones_audios", // nombre de la cola
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("error declarando cola: %v", err)
	}

	return &RabbitPublisher{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *RabbitPublisher) PublicarNotificacion(msg NotificacionCancion) error {
	// Convertir a JSON
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error convirtiendo mensaje a JSON: %v", err)
	}

	// Publicar mensaje
	err = p.channel.Publish(
		"",                      // exchange
		p.queue.Name, 			 // routing key (nombre de la cola)
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("error publicando mensaje: %v", err)
	}

	fmt.Println("Notificacion enviada a RabbitMQ", string(body))
	return nil
}

// Cerrar conexión a la cola
func (p *RabbitPublisher) Cerrar() {	
	p.channel.Close()
	p.conn.Close()
}