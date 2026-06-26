# Explicación del Proyecto - Sistema de Streaming de Audio con Reacciones

## Arquitectura General

```
┌─────────────────────────────────────────────────────────────────────────┐
│                        NAVEGADOR WEB                                     │
│                                                                          │
│  ┌─────────────────────┐    ┌──────────────────────────────────────────┐ │
│  │   Cliente Web       │    │   Admin Web                               │ │
│  │   (index.html)      │    │   (adminweb/index.html)                   │ │
│  ├─────────────────────┤    ├──────────────────────────────────────────┤ │
│  │ funciones.js        │    │ admin.js                                  │ │
│  │ cliente_receptor_   │    │                                           │ │
│  │ fragmentos.js       │    │ POST /registrar (multipart)               │ │
│  │ animaciones.js      │    │   → ServidorAudiosMetadatos :8090         │ │
│  │ estilos.css         │    │                                           │ │
│  └────────┬────────────┘    └──────────────────────────────────────────┘ │
│           │                                                              │
│  ┌────────┴─────────────────────────────────────────────────────────┐    │
│  │  WEB-GRPC (HTTP/1.1)  │  REST (HTTP/1.1)  │  STOMP/WebSocket    │    │
│  │  → localhost:8080     │  → localhost:8090  │  → localhost:5000   │    │
│  └────────┬─────────────────────┬──────────────────┬────────────────┘    │
└───────────┼─────────────────────┼──────────────────┼────────────────────┘
            │                     │                  │
            ▼                     ▼                  ▼
┌──────────────────────┐ ┌─────────────────┐ ┌──────────────────────────┐
│   PROXY ENVOY        │ │  ServidorAudios │ │  ServidorReacciones      │
│   (envoyConfig2.yaml) │ │  Metadatos      │ │  (Spring Boot WebSocket) │
│   Puerto: 8080        │ │  Puerto: 8090   │ │  Puerto: 5000            │
│                       │ │  Go - HTTP      │ │                          │
│  Traduce WEB-GRPC     │ │                 │ │  Broker: /brokerDe-      │
│  (HTTP/1.1) a GRPC    │ │  GET /tipos     │ │  Reacciones              │
│  (HTTP/2)             │ │  GET /audios    │ │  Dests: /app/*           │
│                       │ │  GET /audio/{id}│ │  Endpoint: /ws (SockJS)  │
│                       │ │  POST /registrar│ │                          │
│                       │ │                 │ │  Reenvía reacciones y    │
│                       │ │  Lee/escribe    │ │  notificaciones a        │
│                       │ │  audio_files/   │ │  todos los suscritos     │
└──────────┬────────────┘ └─────────────────┘ └──────────────────────────┘
           │
           ▼
┌─────────────────────────────────────┐
│  ServidorStreaming (Go gRPC)        │
│  Puerto: 50051                      │
│                                     │
│  Recibe petición con título+formato │
│  Lee archivo en fragmentos de 64KB  │
│  Envía streaming de fragmentos      │
└─────────────────────────────────────┘
```

---

## 1. FLUJO DE STREAMING DE AUDIO (VIA ENVOY)

### 1.1 Definición del contrato (Proto)

**`clienteweb/servicios.proto`** — Define el servicio gRPC y los mensajes:

```protobuf
syntax = "proto3";

service AudioService {
  rpc enviarCancionMedianteStream(peticionDTO) returns (stream fragmentoCancion);
}

message peticionDTO {
  string titulo = 1;
  string formato = 2;
}

message fragmentoCancion {
  bytes data = 1;
}
```

A partir de este `.proto` se generan:
- `clienteweb/stubs_generados/servicios_pb.js` — clases JS para serializar/deserializar mensajes
- `clienteweb/stubs_generados/servicios_grpc_web_pb.js` — cliente gRPC-Web para navegador
- `ServidorStreaming/serviciosCancion/servicios.pb.go` — tipos Go
- `ServidorStreaming/serviciosCancion/servicios_grpc.pb.go` — stub servidor gRPC Go

### 1.2 Cliente Web: inicia la solicitud

**`clienteweb/cliente_receptor_fragmentos.js`** — Este archivo se compila con webpack a `bundle.js`:

```javascript
const ENVOY_HOST = 'http://localhost:8080';
const client = new AudioServicePromiseClient(ENVOY_HOST);  // ← apunta al proxy Envoy

function iniciar_streaming_cancion(titulo, formato) {
    const request = new peticionDTO();
    request.setTitulo(titulo);
    request.setFormato(formato);

    const stream = client.enviarCancionMedianteStream(request, {});
    // ↑ LLAMADA GRPC-WEB hacia Envoy en puerto 8080

    stream.on('data', (response) => {
        const data = response.getData_asU8();
        audioChunks.push(data);  // ← acumula fragmentos de 64KB
        if (fragmentCount === 1) {
            audioPlayer.style.display = '';  // ← muestra el reproductor
        }
    });

    stream.on('end', () => {
        const audioBlob = new Blob(audioChunks, { type: 'audio/mpeg' });
        const audioUrl = URL.createObjectURL(audioBlob);
        audioPlayer.src = audioUrl;
        audioPlayer.play();  // ← reproduce el audio completo
    });
}
```

Esta función se exporta a `window` al final:
```javascript
window.iniciar_streaming_cancion = iniciar_streaming_cancion;
```

### 1.3 Quien llama a iniciar_streaming_cancion

**`clienteweb/funciones.js`** — función `reproducirAudioSeleccionado()`:

```javascript
function reproducirAudioSeleccionado() {
    if (!audioSeleccionado) return;
    // 1. Se suscribe a reacciones del audio
    suscribirReaccionCancion(audioSeleccionado.archivo);
    // 2. Llama al streaming via Envoy
    window.iniciar_streaming_cancion(audioSeleccionado.archivo, audioSeleccionado.formato);
}
```

Esta función se asocia al botón "Reproducir" en el `index.html`.

### 1.4 Proxy Envoy: traduce WEB-GRPC a GRPC

**`ProxyEnvoy/envoyConfig2.yaml`** — Configuración del proxy:

```yaml
static_resources:
  listeners:
  - name: listener_0
    address:
      socket_address:
        address: 0.0.0.0
        port_value: 8080    # ← Envoy escucha en puerto 8080
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          # Habilita CORS para permitir peticiones del navegador
          route_config:
            virtual_hosts:
            - name: backend
              domains: ["*"]
              cors:
                allow_origin_string_match: [{ exact: "*" }]
                allow_headers: "x-grpc-web,grpc-timeout,..."
                expose_headers: "grpc-status, grpc-message"
              routes:
              - match: { prefix: "/" }
                route:
                  cluster: local_grpc_service  # ← redirige al cluster
          http_filters:
          - name: envoy.filters.http.cors      # ← filtro CORS
          - name: envoy.filters.http.grpc_web  # ← TRADUCE WEB-GRPC → GRPC
          - name: envoy.filters.http.router    # ← enruta la petición
  clusters:
  - name: local_grpc_service
    http2_protocol_options: {}  # ← usa HTTP/2 para hablar con el backend
    load_assignment:
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 50051  # ← ServidorStreaming en Go
```

**Flujo dentro de Envoy:**
1. Recibe petición WEB-GRPC (HTTP/1.1) del navegador en puerto `8080`
2. Filtro `grpc_web` traduce el mensaje a GRPC nativo (HTTP/2)
3. Filtro `router` envía la petición al cluster `local_grpc_service`
4. El cluster reenvía a `127.0.0.1:50051` usando HTTP/2

### 1.5 Servidor Streaming: recibe y responde con streaming

**`ServidorStreaming/main/servidor.go`** — Arranca el servidor gRPC:

```go
func main() {
    listener, _ := net.Listen("tcp", ":50051")  // ← escucha gRPC
    grpcServer := grpc.NewServer()
    objControlador := capacontroladores.NewControladorServidor(objLogger)
    pb.RegisterAudioServiceServer(grpcServer, objControlador)  // ← registra servicio
    grpcServer.Serve(listener)
}
```

**`ServidorStreaming/capaControladores/controladorEnvioAudio.go`** — Controlador que maneja la solicitud de streaming:

```go
func (thisC *ControladorServidor) EnviarCancionMedianteStream(
    req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {

    // Obtiene IP del cliente desde el contexto gRPC
    direcionCliente := ObtenerDireccionCliente(stream.Context())
    // Registra la solicitud en el log (de forma asíncrona)
    go thisC.logger.AlmacenarSolicitud(req.Titulo, direcionCliente)
    // Combina título y formato: "nombre.mp3"
    combined := req.GetTitulo() + "." + req.GetFormato()
    // Delega a la fachada que lee el archivo y envía fragmentos
    return capafachadaservices.StreamAudioFile(
        combined,
        func(data []byte) error {
            return stream.Send(&pb.FragmentoCancion{Data: data})
        })
}
```

**`ServidorStreaming/capaFachadaServices/audioInt.go`** — Lee el archivo y envía fragmentos de 64KB:

```go
func StreamAudioFile(tituloCancion string, funcionParaEnviarFragmento func([]byte) error) error {
    // Busca el archivo en audio_files/ o canciones/
    file, _ := os.Open("../audio_files/" + tituloCancion)

    buffer := make([]byte, 64*1024)  // ← fragmentos de 64KB
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break  // ← archivo leído completamente
        }
        // Envía el fragmento via el callback (stream.Send)
        err = funcionParaEnviarFragmento(buffer[:n])
        if err != nil {
            return fmt.Errorf("error enviando fragmento: %w", err)
        }
    }
    return nil
}
```

### 1.6 Resumen del flujo de datos (envoy)

```
Navegador                    Envoy :8080                  ServidorStreaming :50051
   │                            │                              │
   │  HTTP/1.1 WEB-GRPC         │                              │
   │  POST (grpc-web)           │                              │
   ├──────────────────────────► │                              │
   │                            │  HTTP/2 gRPC                 │
   │                            │  enviarCancionMedianteStream │
   │                            ├─────────────────────────────►│
   │                            │                              ├── busca archivo
   │                            │                              ├── lee 64KB
   │                            │  ← fragmento 1 (stream)     │
   │                            ├──────────────────────────────┤
   │  ← fragmento 1 (grpc-web)  │                              │
   │◄───────────────────────────┤                              │
   │                            │  ← fragmento 2 (stream)      │
   │                            ├──────────────────────────────┤
   │  ← fragmento 2 (grpc-web)  │                              │
   │◄───────────────────────────┤                              │
   │          ...               │           ...                │
   │                            │  ← stream end                │
   │                            ├──────────────────────────────┤
   │  ← stream end              │                              │
   │◄───────────────────────────┤                              │
   │                            │                              │
   ├─── acumula chunks          │                              │
   ├─── crea Blob URL           │                              │
   ├─── audioPlayer.play()      │                              │
```

---

## 2. FLUJO DE REACCIONES (WebSocket vía STOMP)

### 2.1 Configuración del servidor WebSocket

**`ServidorReacciones/.../WebSocketConfig.java`**:

```java
@Configuration
@EnableWebSocketMessageBroker
public class WebSocketConfig implements WebSocketMessageBrokerConfigurer {
    @Override
    public void configureMessageBroker(MessageBrokerRegistry config) {
        config.enableSimpleBroker("/brokerDeReacciones");  // ← topicos para suscribirse
        config.setApplicationDestinationPrefixes("/app");  // ← prefijo para enviar msgs
    }
    @Override
    public void registerStompEndpoints(StompEndpointRegistry registry) {
        registry.addEndpoint("/ws").setAllowedOriginPatterns("*").withSockJS();
    }
}
```

### 2.2 Cliente: conexión y suscripción

**`clienteweb/funciones.js`** — `conectarReacciones()`:

```javascript
function conectarReacciones() {
    const socket = new SockJS('http://localhost:5000/ws');
    clienteChat = Stomp.over(socket);
    clienteChat.connect({}, function(frame) {
        conectadoReacciones = true;
    });
}
```

### 2.3 Cliente: suscripción a canales del audio

`funciones.js` — `suscribirReaccionCancion(cancionId)`:

```javascript
function suscribirReaccionCancion(cancionId) {
    // Canal de reacciones
    suscripcionReacciones = clienteChat.subscribe(
        '/brokerDeReacciones/cancion/' + cancionId,
        function(mensaje) { mostrarReaccionRecibida(mensaje.body); }
    );
    // Canal de presencia (play/pause)
    suscripcionPresencia = clienteChat.subscribe(
        '/brokerDeReacciones/cancion/' + cancionId + '/presencia',
        function(mensaje) { mostrarNotificacionUsuario(mensaje.body); }
    );
}
```

Se llama desde `reproducirAudioSeleccionado()` **antes** de iniciar el streaming.

### 2.4 Cliente: envía reacción

`funciones.js` — `enviarReaccion(event)`:

```javascript
function enviarReaccion(event) {
    var mensaje = JSON.stringify({
        nickname: usuarioActual,
        cancionId: audioSeleccionado.archivo,
        reaccion: reaccion  // "like", "heart", "sad", "fun"
    });
    clienteChat.send('/app/enviarReaccion', {}, mensaje);
}
```

### 2.5 Cliente: notifica play/pause

`funciones.js` — `notificarEstadoReproduccion(estado)`:

```javascript
function notificarEstadoReproduccion(estado) {
    var accion = estado === 'reproduciendo' ? 'play' : 'stop';
    var mensaje = JSON.stringify({
        nickname: usuarioActual,
        cancionId: audioSeleccionado.archivo,
        accion: accion
    });
    clienteChat.send('/app/notificarEstado', {}, mensaje);
}
```

Se dispara automáticamente desde los listeners del `<audio>`:

```javascript
audio.addEventListener('play', function() {
    notificarEstadoReproduccion('reproduciendo');
});
audio.addEventListener('pause', function() {
    notificarEstadoReproduccion('pausa');
});
```

### 2.6 Servidor Reacciones: recibe y reenvía (broadcast)

**`ReaccionesController.java`**:

```java
@MessageMapping("/enviarReaccion")
public void processReaccion(ReaccionDTO reaccion) {
    String cancionId = reaccion.getCancionId();
    // Reenvía a TODOS los suscritos a /brokerDeReacciones/cancion/{id}
    simpMessagingTemplate.convertAndSend(
            "/brokerDeReacciones/cancion/" + cancionId, reaccion);
}

@MessageMapping("/notificarEstado")
public void processNotification(ReproduccionDTO evento) {
    String cancionId = evento.getCancionId();
    // Reenvía a TODOS los suscritos a /brokerDeReacciones/cancion/{id}/presencia
    simpMessagingTemplate.convertAndSend(
            "/brokerDeReacciones/cancion/" + cancionId + "/presencia", evento);
}
```

### 2.7 Cliente: muestra notificaciones recibidas

`funciones.js` — `mostrarNotificacionUsuario()`:

```javascript
function mostrarNotificacionUsuario(mensaje) {
    var datos = JSON.parse(mensaje);
    var notif = document.createElement("div");
    notif.className = "notificacion-usuario " +
        (datos.accion === 'play' ? 'notif-play' : 'notif-pause');

    var texto = datos.accion === 'play'
        ? '<strong>' + datos.nickname + '</strong> comenzo a reproducir...'
        : '<strong>' + datos.nickname + '</strong> pauso la reproduccion...';

    notif.innerHTML = icono + ' ' + texto;
    contenedor.appendChild(notif);
    animarEntradaNotificacion(notif);  // ← animación slide-in
    setTimeout(() => notif.remove(), 6000);
}
```

---

## 3. FLUJO DE METADATOS Y AUDIOS (REST)

**`ServidorAudiosMetadatos/servidor.go`** — Servidor HTTP en puerto `8090`:

### 3.1 Ver tipos de audio

```
GET http://localhost:8090/tipos
```

Respuesta:
```json
[{"id":1,"nombre":"Musica"}, {"id":2,"nombre":"Podcast"},
 {"id":3,"nombre":"Audiolibro"}, {"id":4,"nombre":"Ruido Blanco"}]
```

Llamado desde `funciones.js` → `cargarTiposAudio()` al cargar la página.

### 3.2 Ver lista de audios por tipo

```
GET http://localhost:8090/audios?tipo=1
```

Respuesta:
```json
[{"id":1,"titulo":"Una de esas canciones","archivo":"cancion1","formato":"mp3","tipoId":1},...]
```

### 3.3 Ver detalles de un audio

```
GET http://localhost:8090/audio/1
```

Respuesta incluye metadatos específicos según el tipo (artista, álbum, género para música; narrador, editorial, ISBN para audiolibros; etc.).

### 3.4 Registrar nuevo audio (Admin)

**`adminweb/admin.js`** — `registrarAudio()`:

```javascript
function registrarAudio() {
    var metadata = {
        titulo: ..., autor: ..., album: ..., genero: ...,
        tipoId: ..., tipoNombre: ..., /* + campos específicos según tipo */
    };
    var formData = new FormData();
    formData.append("metadata", JSON.stringify(metadata));
    formData.append("audio", archivoInput.files[0]);  // ← archivo MP3

    fetch('http://localhost:8090/registrar', {
        method: 'POST',
        body: formData
    });
}
```

El servidor (`servidor.go:manejarRegistrar`) recibe el multipart, guarda el archivo en `audio_files/` y registra los metadatos en memoria.

---

## 4. ANIMACIONES

**`clienteweb/animaciones.js`**:

```javascript
function handleReactionClick(icon) {
    icon.style.transform = 'scale(1.4)';
    icon.style.boxShadow = '0 0 20px gold';
    setTimeout(() => { icon.style.transform = ''; icon.style.boxShadow = ''; }, 300);
}

function animarEntradaNotificacion(elemento) {
    elemento.style.transform = 'translateX(100px)';
    elemento.style.opacity = '0';
    setTimeout(() => {
        elemento.style.transition = 'all 0.5s ease';
        elemento.style.transform = 'translateX(0)';
        elemento.style.opacity = '1';
    }, 50);
}
```

**`clienteweb/estilos.css`** incluye:
- `.reaccion-burbuja` — burbujas flotantes con fade-in para reacciones
- `.notificacion-usuario` — notificaciones con efecto slide-in desde la derecha
- Transiciones CSS para hover en botones de tipo de audio
- Efectos de escala y sombra en reacciones

---

## 5. MAPA DE ARCHIVOS Y UBICACIÓN DEL CÓDIGO

| Archivo | Propósito |
|---------|-----------|
| `clienteweb/index.html` | Interfaz gráfica (tipos, lista, detalles, reproductor, panel reacciones) |
| `clienteweb/funciones.js` | Conexión STOMP/WebSocket, suscripción a canales, envío de reacciones y estado, carga de metadatos vía REST |
| `clienteweb/cliente_receptor_fragmentos.js` | Cliente gRPC-Web que hace streaming via Envoy, crea Blob URL y reproduce |
| `clienteweb/webpack.config.js` | Compila `cliente_receptor_fragmentos.js` → `bundle.js` |
| `clienteweb/animaciones.js` | Efectos visuales (glow en reacciones, slide-in en notificaciones) |
| `clienteweb/estilos.css` | Estilos de toda la interfaz |
| `clienteweb/servicios.proto` | Definición del contrato gRPC |
| `clienteweb/stubs_generados/servicios_pb.js` | Clases protobuf generadas para JS |
| `clienteweb/stubs_generados/servicios_grpc_web_pb.js` | Stub cliente gRPC-Web generado |
| `ProxyEnvoy/envoyConfig2.yaml` | Configuración de Envoy (puerto 8080, CORS, filtro grpc_web) |
| `ServidorStreaming/main/servidor.go` | Servidor gRPC en Go (puerto 50051) |
| `ServidorStreaming/capaControladores/controladorEnvioAudio.go` | Controlador que recibe petición y envía streaming |
| `ServidorStreaming/capaFachadaServices/audioInt.go` | Fachada que lee archivo en fragmentos de 64KB |
| `ServidorStreaming/capalogger/logger.go` | Logger singleton (registra solicitudes) |
| `ServidorReacciones/.../WebSocketConfig.java` | Configura STOMP broker (topic `/brokerDeReacciones`, endpoint `/ws`) |
| `ServidorReacciones/.../ReaccionesController.java` | Controlador que reenvía reacciones y notificaciones a suscriptores |
| `ServidorAudiosMetadatos/servidor.go` | API REST (tipos, audios, detalle, registro) |
| `ServidorAudiosMetadatos/repositorioAudios.go` | Datos en memoria (8 audios precargados) |
| `adminweb/index.html` | Interfaz de administrador para registrar audios |
| `adminweb/admin.js` | Lógica del formulario de registro y envío multipart |

---

## 6. PUERTOS Y ORDEN DE EJECUCIÓN

| Servicio | Puerto | Comando |
|----------|--------|---------|
| ServidorStreaming | 50051 | `go run main/servidor.go` (dentro de `ServidorStreaming/`) |
| Proxy Envoy | 8080 | `./envoyServer -c envoyConfig2.yaml --disable-hot-restart` |
| ServidorAudiosMetadatos | 8090 | `go run servidor.go` (dentro de `ServidorAudiosMetadatos/`) |
| ServidorReacciones | 5000 | `mvn spring-boot:run` o ejecutar JAR |

**Orden de arranque recomendado:**
1. Servidor gRPC de Streaming (`:50051`)
2. Proxy Envoy (`:8080`)
3. Servidor de Audios y Metadatos (`:8090`)
4. Servidor de Reacciones (`:5000`)
5. Abrir `clienteweb/index.html` en el navegador

---

## 7. RÚBRICA DE CALIFICACIÓN — DÓNDE SE IMPLEMENTA CADA CRITERIO

| # | Criterio (0.5 c/u) | Dónde se implementa |
|---|-------------------|---------------------|
| 1 | Nickname al iniciar reproducción | `funciones.js:210` — listener `play` → `notificarEstadoReproduccion('reproduciendo')` → servidor reenvía a presencia → `mostrarNotificacionUsuario()` |
| 2 | Nickname al pausar reproducción | `funciones.js:215` — listener `pause` → `notificarEstadoReproduccion('pausa')` → mismo flujo |
| 3 | Reacciones de usuarios | `funciones.js:72` — `enviarReaccion()` → servidor reenvía a topic → `mostrarReaccionRecibida()` |
| 4 | Ver tipos de audio (cliente y admin) | `funciones.js:230` — `cargarTiposAudio()` fetch a `GET /tipos`; admin: formulario con `<select>` de tipos |
| 5 | Ver metadatos de audio actual y nuevo | `funciones.js:287` — `seleccionarAudio()` fetch a `GET /audio/{id}`, `mostrarDetalles()` renderiza metadatos específicos |
| 6 | Reproducir audio actual y nuevo | `funciones.js:347` — `reproducirAudioSeleccionado()` → `window.iniciar_streaming_cancion()` gRPC-Web via Envoy |
| 7 | Registrar metadatos y audio | `adminweb/admin.js` — `registrarAudio()` POST multipart a `GET /registrar` con JSON metadata + archivo |
| 8 | Animaciones para nicknames y reacciones | `animaciones.js` — `handleReactionClick()` (glow), `animarEntradaNotificacion()` (slide-in); CSS con transiciones |
| 9 | Explicación del código fuente | Este archivo (`explicacion.md`) |
| 10 | Sustentación | Presentación oral del proyecto |

---

## 8. NOTAS IMPORTANTES

- **Envoy es necesario** porque el navegador solo soporta HTTP/1.1 y gRPC requiere HTTP/2 con streaming. Envoy traduce entre ambos protocolos.
- El archivo `cliente_receptor_fragmentos.js` se compila con **webpack** porque usa `require()` para importar los stubs. Después de modificar este archivo, ejecutar `npx webpack` para regenerar `bundle.js`.
- Los archivos de audio se almacenan en `audio_files/` y se sirven desde el ServidorStreaming.
- Las reacciones y notificaciones de presencia se manejan mediante **topics STOMP** del servidor de reacciones, que actúa como broker. No hay almacenamiento persistente de reacciones.
