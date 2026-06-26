# Sistema de Consulta y Streaming de Audio gRPC

Sistema distribuido de consulta de metadatos y streaming de audio desarrollado en Go utilizando gRPC para Linux.

## Descripción del Proyecto

Esta aplicación implementa un sistema de consulta de metadatos de audio y streaming en tiempo real, separados en dos servidores:

- **ServidorDeAudios** (Puerto 50051): Gestiona los metadatos de los audios (título, autor, álbum, género, duración, etc.)
- **ServidorDeStreaming** (Puerto 50052): Gestiona el streaming de los archivos de audio

### Características

- Interfaz de consola interactiva estilo Spotify
- Cuatro tipos de audio: Música, Podcast, Audiolibro y Ruido Blanco
- Streaming de audio en chunks mediante gRPC
- Arquitectura por capas (Acceso a Datos, Fachada, Controladores)
- Comunicación gRPC con logging de llamadas RPC
- Patrones de diseño: Capas, MVC, DTO
- Al menos 2 audios por cada tipo

## Estructura del Proyecto

```
final/
├── Cliente/                         # Aplicación cliente
│   ├── main/
│   │   └── cliente.go               # Punto de entrada del cliente
│   ├── vistas/
│   │   └── menu.go                   # Interfaz de usuario por consola (VISTAS)
│   ├── utilidades/
│   │   └── utilidades.go             # Funciones de reproducción de audio
│   └── go.mod                        # Dependencias del cliente
│
├── ServidorDeAudios/                 # Servidor de metadatos
│   ├── capaAccesoDatos/
│   │   └── repositorioAudios.go      # Datos de audios (MODELO)
│   ├── capaFachada/
│   │   └── fachadaServicios.go       # Lógica de negocio
│   ├── capaControladores/
│   │   └── ControladorAudios.go      # Controlador gRPC
│   ├── modelos/
│   │   └── MetadataAudio.go          # Estructuras de datos
│   ├── main/
│   │   └── servidor.go               # Punto de entrada
│   └── protos/
│       ├── audio_metadata.proto      # Definición de servicios
│       └── audio_metadata.pb.go      # Código generado
│
└── ServidorDeStreaming/              # Servidor de streaming
    ├── capaAccesoDatos/
    │   └── repositorioAudio.go       # Acceso a archivos de audio
    ├── capaFachada/
    │   └── fachadaStreaming.go       # Lógica de streaming
    ├── capaControladores/
    │   └── ControladorStreaming.go   # Controlador gRPC
    ├── main/
    │   └── servidor.go                # Punto de entrada
    ├── audios/                       # Archivos MP3 disponibles
    │   ├── musica1.mp3
    │   ├── musica2.mp3
    │   ├── podcast1.mp3
    │   ├── podcast2.mp3
    │   ├── audiolibro1.mp3
    │   ├── audiolibro2.mp3
    │   ├── ruido1.mp3
    │   └── ruido2.mp3
    └── protos/
        ├── audio_streaming.proto     # Definición de servicios
        └── audio_streaming.pb.go     # Código generado
```

## Requisitos Previos

- Go 1.24 o superior
- protoc (Protocol Buffers compiler)
- Bibliotecas Go:
  - google.golang.org/grpc
  - google.golang.org/protobuf
  - github.com/faiface/beep
  - github.com/golang/protobuf

## Compilación

### 1. Compilar Servidor de Metadatos

```bash
cd final/ServidorDeAudios
go build -o servidor main/servidor.go
```

### 2. Compilar Servidor de Streaming

```bash
cd final/ServidorDeStreaming
go build -o servidor main/servidor.go
```

### 3. Compilar Cliente

```bash
cd final/Cliente
go build -o cliente main/cliente.go
```

## Ejecución

### Paso 1: Iniciar el Servidor de Metadatos

```bash
cd final/ServidorDeAudios
./servidor
```

El servidor escuchará en el puerto 50051 y mostrará:
```
Servidor de Metadatos gRPC escuchando en puerto 50051...
Servicios disponibles:
  - GetAudioTypes
  - GetAudiosByType
  - GetAudioDetails
  - GetAudioFilename
  - BuscarAudio
```

### Paso 2: Iniciar el Servidor de Streaming

```bash
cd final/ServidorDeStreaming
./servidor
```

El servidor escuchará en el puerto 50052 y mostrará:
```
Servidor de Streaming gRPC escuchando en puerto 50052...
Servicios disponibles:
  - StreamAudio
```

### Paso 3: Iniciar el Cliente

```bash
cd final/Cliente
./cliente
```

## Uso de la Aplicación

### Menú Principal

```
----------------------------------------
    APLICACIÓN DE AUDIO gRPC            
----------------------------------------
  1. Ver Tipos de Audio                 
  2. Salir                               
----------------------------------------
```

### Menú de Tipos de Audio

```
----------------------------------------
        TIPOS DE AUDIO                  
----------------------------------------

  1. Música
  2. Podcast
  3. Audiolibro
  4. Ruido Blanco
  5. Atrás
```

### Lista de Audios (Ejemplo: Música)

```
----------------------------------------
        AUDIOS TIPO: Música             
----------------------------------------

  1. Bohemian Rhapsody
  2. Stairway to Heaven
  3. Atrás
```

### Detalles del Audio

```
----------------------------------------
        DETALLES DEL AUDIO             
----------------------------------------

  Título:              Bohemian Rhapsody
  Autor:               Queen
  Álbum:               A Night at the Opera
  Género Musical:      Rock
  Sello Discográfico:  EMI
  Año de Lanzamiento: 1975
  Duración:            354 segundos

  Disponible:       Sí

  1. Reproducir
  2. Atrás
```

### Reproducción de Audio

```
----------------------------------------
        REPRODUCIENDO AUDIO            
----------------------------------------

Recibiendo y reproduciendo audio en vivo...

Presione Enter para volver al menú...
```

## Metadatos por Tipo de Audio

### Música
- Título, Autor, Álbum, Género Musical, Sello Discográfico, Año de Lanzamiento, Duración

### Podcast
- Título del Episodio, Nombre del Podcast, Anfitrión, Temporada/Episodio, Clasificación de Contenido, Notas del Show, Año

### Audiolibro
- Título del Libro, Autor, Narrador, Editorial, ISBN, Capítulo, Año

### Ruido Blanco
- Título, Tipo de Sonido, Fuente del Audio, Proveedor, Uso Sugerido, Duración del Bucle, Frecuencia Dominante

## Servicios gRPC

### ServidorDeAudios (Puerto 50051)

| Servicio | Descripción |
|----------|-------------|
| GetAudioTypes | Retorna la lista de tipos de audio disponibles |
| GetAudiosByType | Retorna la lista de audios de un tipo específico |
| GetAudioDetails | Retorna los detalles completos de un audio |
| GetAudioFilename | Retorna el nombre del archivo de audio |
| BuscarAudio | Busca un audio por título |

### ServidorDeStreaming (Puerto 50052)

| Servicio | Descripción |
|----------|-------------|
| StreamAudio | Envía el archivo de audio en chunks mediante streaming |

## Patrones de Diseño Implementados

### 1. Patrón Capas
- **Capa Acceso a Datos**: Repositorios que gestionan los datos
- **Capa Fachada**: Lógica de negocio y transformación de datos
- **Capa Controladores**: Implementación de los métodos RPC

### 2. Patrón MVC
- **Modelo**: estructuras en `modelos/` y `capaAccesoDatos/`
- **Vista**: módulo `vistas/` con menú interactivo
- **Controlador**: `capaControladores/` con lógica de presentación

### 3. Patrón DTO
- Messages en archivos proto que encapsulan datos para transporte via gRPC

## Logging de Llamadas RPC

Los servidores registran cada llamada RPC recibida utilizando `fmt.Printf`:

**ServidorDeAudios:**
```
RPC Call: GetAudioTypes
RPC Call: GetAudiosByType - Tipo ID: 1
RPC Call: GetAudioDetails - ID: musica_001
```

**ServidorDeStreaming:**
```
RPC Call: StreamAudio - Archivo: musica1.mp3
```

## Archivos Proto

Los archivos proto se encuentran en:
- `ServidorDeAudios/protos/audio_metadata.proto`
- `ServidorDeStreaming/protos/audio_streaming.proto`

## Problemas Comunes

1. **Error de conexión**: Asegúrese de que los servidores estén ejecutándose antes de iniciar el cliente
2. **Audio no encontrado**: Verifique que los archivos MP3 existan en `ServidorDeStreaming/audios/`
3. **Error de compilación**: Ejecute `go mod tidy` en cada directorio para descargar las dependencias