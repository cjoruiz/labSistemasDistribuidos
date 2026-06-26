# LAB21ABRIL - Microservicios con RabbitMQ

## Descripción
Sistema distribuido de **almacenamiento de canciones y notificación por correo** compuesto por dos microservicios independientes que se comunican de forma asíncrona a través de **RabbitMQ**.

## Arquitectura
```
Cliente HTTP --> [Go Microservice :5000] --> RabbitMQ --> [Spring Boot :6000]
  (POST mp3 + metadata)                        |              (simula envío de correo)
                                          notificaciones_audios
```

## Componentes

### mi-proyecto-go (Go) - Almacenamiento de Canciones
- Endpoint `POST /canciones/almacenamiento` (multipart: archivo mp3 + título, artista, género)
- Guarda el archivo MP3 en disco (`audios/`)
- Publica notificación JSON en RabbitMQ cola `notificaciones_audios`

### microservicioEnvioCorreos (Spring Boot) - Notificación por Correo
- Escucha la cola `notificaciones_audios` de RabbitMQ
- Al recibir un mensaje, simula el envío de un correo (barra de progreso en consola)
- Muestra título, artista y género de la canción recibida

## Tecnologías
- **Go**: gorilla/mux, streadway/amqp
- **Java 17**, Spring Boot 3.3.4, Spring AMQP, Jackson, Lombok
- **RabbitMQ**: message broker, cola durable `notificaciones_audios`

## Ejecución
1. Iniciar RabbitMQ
2. Iniciar microservicio de correos: `mvn spring-boot:run` (puerto `6000`)
3. Iniciar servidor Go: `go run main/main.go` (puerto `5000`)
4. Enviar POST con multipart a `http://localhost:5000/canciones/almacenamiento`
