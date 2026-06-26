# LAB28ABRIL - Sistema de Reacciones en Tiempo Real con WebSockets

## Descripción
Sistema de **reacciones musicales en tiempo real** que permite a múltiples usuarios conectados enviar y recibir reacciones (like, heart, sad, fun) al instante mediante **WebSockets con STOMP y SockJS**.

## Arquitectura
```
Navegador (STOMP/SockJS) <---> Spring Boot Server :5000 (STOMP Broker)
                                        |
                                  /brokerDeReacciones/reaccionesPorCancion
```

## Componentes

### servidorchat (Spring Boot)
- Servidor WebSocket con STOMP en puerto `5000`
- Endpoint SockJS: `/ws`
- Broker simple en `/brokerDeReacciones`
- Prefijo de aplicación: `/app`
- Controlador `ReaccionesController` que retransmite mensajes a todos los suscriptores

### vista/ (Frontend HTML/JS/CSS)
- Interfaz web con 4 iconos de reacción (Font Awesome)
- Botones Conectar/Desconectar
- Burbujas animadas que flotan y desaparecen al recibir reacciones
- Efecto de brillo dorado al hacer clic

## Tecnologías
- Java 17, Spring Boot 3.3.11, Spring WebSocket, STOMP, SockJS
- HTML5, CSS3, JavaScript vanilla, Font Awesome 6

## Flujo de Comunicación
1. Usuario se conecta via SockJS/STOMP a `http://localhost:5000/ws`
2. Se suscribe a `/brokerDeReacciones/reaccionesPorCancion`
3. Al hacer clic en una reacción, envía mensaje a `/app/enviarReaccion`
4. El servidor retransmite a todos los suscriptores
5. La reacción aparece como burbuja animada (3 segundos)

## Ejecución
1. Iniciar servidor: `mvn spring-boot:run` (puerto `5000`)
2. Abrir `vista/index.html` en el navegador
3. Conectar y enviar reacciones
