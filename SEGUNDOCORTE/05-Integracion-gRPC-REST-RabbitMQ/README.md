# Sistema de Administración de Audio Distribuido

Sistema distribuido de administración, streaming de audio y notificaciones por correo.

## Arquitectura del Sistema

```
┌──────────┐   POST/multipart   ┌────────────┐    gRPC      ┌──────────────────┐
│ Postman  │───────────────────▶│ ADMINISTRADOR│◀───────────▶│ Serv.DeStreaming │
│          │                    │   (Java)    │  50052      │     (Go)         │
└──────────┘                    └──────┬──────┘              └────────┬─────────┘
                                       │  gRPC                     │
                                       │  ┌─────────────────────────┘
                                       │  │                    gRPC (50052)
                                       ▼  ▼                               ▼
                            ┌──────────────────┐           ┌────────────┐
                            │Serv.DeAudios     │           │  Cliente   │
                            │   (Go)           │           │   (Go)     │
                            │Puerto 50051/8090 │           └────────────┘
                            └────────┬────────┘
                                     │
                                     ▼
                            ┌──────────┐
                            │ RabbitMQ │
                            └──────────┘
                                     │
                                     ▼
                            ┌──────────────┐
                            │Serv.Correos  │
                            │   (Java)     │
                            │Puerto 6000   │
                            └──────────────┘

         Callback ◀────────────────────────────────────────
         (cuando cliente reproduce un audio a todos los admins)
```

## Componentes del Sistema

### 1. Administrador (Java - Puerto 8080 o dinámico)
- Recibe audio y metadatos por POST de Postman
- Conecta con Servidor Streaming (Go) por **gRPC** (puerto 50052)
- Conecta con Servidor Metadatos (Go) por REST (puerto 8090)
- Recibe callback de reproducción por REST (varios admins soportados)
- Al iniciarse, registra automáticamente su URL de callback en el ServidorDeStreaming

### 2. Servidor de Streaming (Go - Puerto 50052 y 8091)
- **gRPC (50052)**: Streaming de audio al Cliente y almacenamiento desde Administrador
- **HTTP (8091)**: Almacenar nuevos audios (legacy, aún soportado)
- Envía callback a TODOS los administradores registrados cuando se reproduce un audio
- Mantiene lista de URLs de callback para soporte múltiple

### 3. Servidor de Metadatos (Go - Puerto 50051 y 8090)
- **gRPC (50051)**: Consulta de metadatos existente (para Cliente)
- **REST (8090)**: Recibir nuevos metadatos (desde Administrador)
- **RabbitMQ**: Publica notificaciones para servidor de correos

### 4. Servidor de Correos (Java - Puerto 6000)
- Consume mensajes de RabbitMQ
- Envía correo con metadatos y frase motivadora

### 5. Cliente Go
- Conecta al ServidorDeStreaming (puerto 50052) para reproducir audios
- Conecta al ServidorDeMetadatos (puerto 50051) para ver metadatos
- Requiere autenticación con nickname y contraseña (usuarios hardcodeados)

## Estructura del Proyecto

```
prueba/
├── Administrador/                    # Aplicación Java + gRPC
│   ├── pom.xml                       # Configuración Maven con plugins gRPC
│   ├── src/main/
│   │   ├── proto/audio_streaming.proto
│   │   ├── java/co/edu/unicauca/admin/
│   │   │   ├── AdministradorApplication.java   # Registro callback al iniciar
│   │   │   ├── controlador/AdminControlador.java
│   │   │   ├── modelo/AudioDTO.java, CallbackReproduccionDTO.java
│   │   │   └── servicio/
│   │   │       ├── ServicioAudio.java
│   │   │       └── ClienteStreaming.java       # Cliente gRPC real
│   │   └── resources/application.properties
│   └── target/administrador-1.0.0.jar
│
├── ServidorDeStreaming/              # Go + gRPC + arquitectura por capas
│   ├── main/servidor.go              # Main limpio, solo inicializa
│   ├── capaControladores/
│   │   ├── ControladorStreaming.go  # StreamAudio + RegistrarCallback (gRPC)
│   │   └── ControladorHTTP.go       # Endpoint REST /almacenar (HTTP)
│   ├── capaFachada/
│   │   └── fachadaStreaming.go       # EnviarCallback a múltiples URLs
│   ├── capaAccesoDatos/
│   │   └── repositorioAudio.go
│   ├── protos/audio_streaming.proto
│   └── audios/                       # Archivos MP3 almacenados
│
├── ServidorDeAudiosMetadatos/        # Go + gRPC/REST + arquitectura por capas
│   ├── main/servidor.go              # Main limpio
│   ├── capaControladores/
│   │   ├── ControladorAudios.go      # gRPC handlers
│   │   └── ControladorHTTP.go        # REST handlers
│   ├── capaFachada/
│   ├── capaAccesoDatos/
│   ├── modelos/
│   ├── protos/
│   └── componenteConexionCola/       # Lógica de RabbitMQ
│
├── Cliente/                          # Cliente Go con arquitectura por capas
│   ├── main/
│   │   └── cliente.go                # Main limpio, solo inicializa conexiones
│   ├── vistas/
│   │   ├── menu.go                   # Menús de navegación
│   │   └── autenticacion.go          # Validación de credenciales
│   ├── utilidades/
│   │   ├── utilidades.go            # Funciones UI y decodificación audio
│   │   └── reproductor.go            # Reproducción de audio
│   └── protos/
│
└── microservicioEnvioCorreos - plantilla/  # Java + Spring Boot
```

## Requisitos Previos

- **Go 1.24+** para servidores de Streaming y Metadatos
- **Java 17+** y **Maven** para Administrador y Servidor de Correos
- **RabbitMQ** ejecutándose en localhost:5672 (admin:1234)
- **protoc** para compilar archivos proto (ubicación: /usr/bin/protoc)

## Compilación

### 1. Compilar Servidor de Metadatos (Go)
```bash
cd ServidorDeAudiosMetadatos
go build -o servidor main/servidor.go
```

### 2. Compilar Servidor de Streaming (Go)
```bash
cd ServidorDeStreaming
go build -o servidor main/servidor.go
```

### 3. Compilar Cliente (Go)
```bash
cd Cliente
go build -o cliente main/cliente.go
```

### 4. Compilar Administrador (Java) - Incluye generación de stubs gRPC
```bash
cd Administrador
mvn clean package
```

### 5. Compilar Servidor de Correos (Java)
```bash
cd "microservicioEnvioCorreos - plantilla"
mvn clean package
```

## Ejecución

### Paso 1: Iniciar RabbitMQ
```bash
rabbitmq-server
```

### Paso 2: Iniciar Servidor de Metadatos
```bash
cd ServidorDeAudiosMetadatos
./servidor
```

### Paso 3: Iniciar Servidor de Streaming
```bash
cd ServidorDeStreaming
./servidor
```

### Paso 4: Iniciar Servidor de Correos
```bash
cd "microservicioEnvioCorreos - plantilla"
mvn spring-boot:run
```

### Paso 5: Iniciar Administrador
```bash
cd Administrador
java -jar target/administrador-1.0.0.jar
```

### Ejecutar Múltiples Administradores

Para ejecutar varios administradores en diferentes puertos:
```bash
# Administrador 1 (puerto 8080 por defecto)
java -jar target/administrador-1.0.0.jar

# Administrador 2 (puerto dinámico)
java -jar target/administrador-1.0.0.jar --server.port=0

# Administrador 3 (puerto específico)
java -jar target/administrador-1.0.0.jar --server.port=8081
```

## Endpoints del Administrador (Puerto 8080 o dinámico)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/test` | Verificar que el servidor funciona |
| POST | `/api/audios` | Almacenar nuevo audio (multipart) |
| GET | `/api/audios` | Listar todos los audios |
| GET | `/api/audios/{id}` | Ver detalle de un audio |
| POST | `/api/callback/reproducir` | Recibir callback de reproducción |

## Ejemplo de Uso con Postman

### Almacenar un nuevo audio:

**URL:** `POST http://localhost:8080/api/audios`

**Body (form-data):**
- Key: `archivo` - Value: [Seleccionar archivo MP3]
- Key: `metadata` - Value:
```json
{
  "titulo": "Nueva Canción",
  "autor": "Artista Ejemplo",
  "album": "Álbum Demo",
  "genero": "Pop",
  "duracion": 180,
  "tipoId": 1,
  "tipoNombre": "Música",
  "fechaLanzamiento": "2026",
  "disponible": true,
  "selloDiscografico": "EMI"
}
```

**IMPORTANTE**: El campo `tipoNombre` debe ser exacto:
- "Música" (con tilde), "Podcast", "Audiolibro", "Ruido Blanco"

### Listar audios:

**URL:** `GET http://localhost:8080/api/audios`

## Flujo de Callback (Notificaciones)

### Cuándo se envía el callback
Cuando el Cliente reproduce un audio, el ServidorDeStreaming envía un callback a **todos** los administradores registrados con:
- **ID del audio**: nombre del archivo reproducido
- **Fecha/Hora**: momento de la reproducción (formato: YYYY-MM-DD HH:MM:SS)

### Registro de callback
1. Al iniciar, cada Administrador registra su URL de callback en el ServidorDeStreaming via gRPC (puerto 50052)
2. El ServidorDeStreaming guarda las URLs en una lista (soporta múltiples administradores)
3. Cuando se reproduce un audio, se envía el callback a todas las URLs registradas

### Ejemplo de notificación en consola del Administrador:
```
═══════════════════════════════════════════════════════════
           NUEVA REPRODUCCIÓN DETECTADA                   
═══════════════════════════════════════════════════════════
  ID Audio:       nombre_archivo.mp3
  Título:         nombre_archivo.mp3
  Fecha/Hora:     2026-05-19 01:25:02
═══════════════════════════════════════════════════════════
```

## Flujo de Operaciones

### 1. Almacenar Audio (Postman → Admin → Servidores)
1. Postman envía POST `/api/audios` con archivo y metadatos
2. Administrador reenvía archivo al Servidor Streaming (**gRPC** puerto 50052)
3. Administrador reenvía metadatos al Servidor Metadatos (REST puerto 8090)
4. Servidor Metadatos publica en cola RabbitMQ
5. Servidor de Correos consume y muestra correo con frase motivadora

### 2. Reproducción (Cliente → Streaming → Admin)
1. Cliente solicita reproducir audio al ServidorDeStreaming (gRPC)
2. Servidor Streaming envía audio por streaming
3. Servidor Streaming detecta reproducción y envía callback
4. Callback se envía a TODOS los administradores registrados
5. Cada Administrador muestra notificación con ID audio y fecha/hora

### 3. Notificación por Correo
1. Nuevo audio almacenado
2. Servidor Metadatos publica en cola `notificaciones_audios`
3. Servidor de Correos consume mensaje
4. Muestra correo con: ID, título, artista, género, fecha registro, frase motivadora

## Frases Motivadoras

El servidor de correos incluye frases como:
- "La música es el lenguaje de los pulmones."
- "Cada canción es una aventura musical."
- "El ritmo de la vida está en la música."
- "Deja que la melodía guíe tu alma."
- "La música conecta corazones."
- "Un nuevo audio, una nueva historia por contar."
- "Tu biblioteca musical crece con cada nota."

## Notas Importantes

1. **Orden de servicios**: El Administrador debe iniciarse DESPUÉS del ServidorDeStreaming para que el registro de callback funcione correctamente.

2. **Puerto dinámico del Administrador**: Si se usa `--server.port=0`, Spring Boot asigna un puerto libre automáticamente. El código detecta este puerto para registrar la URL de callback correcta.

3. **Callback a múltiples administradores**: El ServidorDeStreaming envía el callback a TODOS los administradores registrados, no solo a uno.

4. **Cliente Go**: Necesita ejecutarse en una terminal separada. Usa los servidores gRPC directamente (no el Administrador).

5. **RabbitMQ**: Necesario para la comunicación entre ServidorDeMetadatos y ServidorDeCorreos (notificaciones por correo).

6. **gRPC**: El Administrador usa gRPC para comunicarse con el ServidorDeStreaming (almacenar audio y registrar callback). Los stubs se generan automáticamente durante la compilación con Maven.

7. **Autenticación del Cliente**: Al iniciar, el Cliente solicita nickname y contraseña. Usuarios disponibles:
   - `admin` / `admin123`
   - `juan` / `juan123`
   - `maria` / `maria123`
   - `carlos` / `carlos123`

## Estructura de Capas (ServidorDeAudiosMetadatos)

El servidor sigue arquitectura por capas para mejor organización:
```
ServidorDeAudiosMetadatos/
├── main/
│   └── servidor.go              # Main limpio, solo inicializa
├── capaControladores/
│   ├── ControladorAudios.go     # gRPC handlers
│   └── ControladorHTTP.go       # REST handlers
├── capaAccesoDatos/
│   └── repositorioAudios.go     # Datos en memoria
├── capaFachada/
│   └── fachadaServicios.go      # Lógica de negocio
├── modelos/
│   └── MetadataAudio.go         # Estructura de datos
└── componenteConexionCola/
    └── publicadorRabbit.go      # Lógica de RabbitMQ separada
```

## Estructura de Capas (ServidorDeStreaming)

El servidor sigue arquitectura por capas para mejor organización:
```
ServidorDeStreaming/
├── main/
│   └── servidor.go              # Main limpio, solo inicializa
├── capaControladores/
│   ├── ControladorStreaming.go  # StreamAudio + RegistrarCallback (gRPC)
│   └── ControladorHTTP.go       # Endpoint REST /almacenar (HTTP)
├── capaFachada/
│   └── fachadaStreaming.go      # Lógica de negocio y callbacks
├── capaAccesoDatos/
│   └── repositorioAudio.go      # Acceso a archivos de audio
├── modelos/
│   └── (definidos en protos)
└── protos/
```

## Estructura de Capas (Cliente)

El Cliente está organizado en capas bien definidas:
```
Cliente/
├── main/
│   └── cliente.go                # Punto de entrada - solo inicializa conexiones
├── vistas/                       # Capa de presentación (UI)
│   ├── menu.go                   # Menús de navegación
│   └── autenticacion.go          # Validación de credenciales
├── utilidades/                   # Capa de utilidades y servicios
│   ├── utilidades.go             # Funciones UI y decodificación audio
│   └── reproductor.go            # Reproducción de audio
└── protos/                        # Stubs gRPC generados
```

## Issue Conocido

### Nuevo audio no aparece en el Cliente
**Problema**: El campo `tipoNombre` debe coincidir exactamente.
**Solución**: Al enviar audio desde Postman, usar:
- `"tipoNombre": "Música"` (con tilde, M mayúscula)
- `"tipoNombre": "Podcast"`
- `"tipoNombre": "Audiolibro"`
- `"tipoNombre": "Ruido Blanco"`

Si se envía "Musica" o "musica" no aparecerá en las listas del cliente.