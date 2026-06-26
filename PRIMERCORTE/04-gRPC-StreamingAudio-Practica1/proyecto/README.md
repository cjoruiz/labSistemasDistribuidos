# Sistema de Audio Distribuido — gRPC (Go)

**Universidad del Cauca — Laboratorio de Sistemas Distribuidos**
Integrantes: Ortega Ruiz Cristian · Ordóñez Chacón Edwin

---

## Arquitectura

```
┌──────────────────────────────────────────────────────┐
│                      Cliente                         │
│  (menú consola, puerto 50052 ↔ 50051)                │
└─────────────┬───────────────────┬────────────────────┘
    gRPC       │  Interface 1      │  Interface 2
  metadatos   ▼                   ▼  streaming
┌─────────────────┐    ┌──────────────────────┐
│ ServidorDeAudios│    │ ServidorDeStreaming   │
│  puerto :50052  │    │  puerto :50051        │
└─────────────────┘    └──────────────────────┘
```

---

## Requisitos previos

```bash
# Go 1.21+
go version

# Instalar protoc (Protocol Buffers compiler)
sudo apt install -y protobuf-compiler

# Instalar plugins Go para protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## Pasos para compilar y ejecutar

### 1. Generar archivos protobuf

```bash
chmod +x generar_protos.sh
./generar_protos.sh
```

Esto genera los archivos `*.pb.go` y `*_grpc.pb.go` en:
- `ServidorDeAudios/serviciosMetadatos/`
- `ServidorDeStreaming/serviciosStreaming/`

### 2. Agregar archivos mp3 de prueba (ServidorDeStreaming)

El ServidorDeStreaming busca los archivos en la ruta relativa que entrega el ServidorDeAudios:

```bash
mkdir -p ServidorDeStreaming/audios
# Copiar 8 archivos mp3 con los siguientes nombres:
# musica1.mp3, musica2.mp3
# podcast1.mp3, podcast2.mp3
# audiolibro1.mp3, audiolibro2.mp3
# ruidoblanco1.mp3, ruidoblanco2.mp3
```

> Puede usar cualquier mp3 corto para pruebas; renómbrelo según la lista anterior.

### 3. Compilar

```bash
chmod +x compilar.sh
./compilar.sh
```

### 4. Ejecutar (3 terminales separadas)

```bash
# Terminal 1 — ServidorDeAudios
./bin/ServidorDeAudios

# Terminal 2 — ServidorDeStreaming (desde su directorio para que encuentre ./audios/)
cd ServidorDeStreaming && ../bin/ServidorDeStreaming

# Terminal 3 — Cliente
./bin/Cliente
```

---

## Estructura del proyecto

```
.
├── Cliente/
│   ├── capaControladores/
│   │   ├── ControladorMetadatos.go   # RPC → ServidorDeAudios
│   │   └── ControladorStreaming.go   # RPC → ServidorDeStreaming
│   ├── utilidades/
│   │   ├── Consola.go                # Lectura de entrada y formato
│   │   └── Reproductor.go            # Decodificación y reproducción mp3
│   ├── vistas/
│   │   ├── MenuPrincipal.go          # Menú raíz
│   │   ├── MenuTipos.go              # Lista de tipos de audio
│   │   ├── MenuAudios.go             # Lista de audios por tipo
│   │   ├── MenuDetalles.go           # Metadatos del audio seleccionado
│   │   └── MenuReproduccion.go       # Pantalla de reproducción streaming
│   ├── main/cliente.go
│   └── go.mod
│
├── ServidorDeAudios/
│   ├── capaAccesoDatos/
│   │   └── RepositorioAudios.go      # Datos en memoria (tipos + 8 audios)
│   ├── capaFachada/
│   │   └── FachadaMetadatos.go       # Delegación hacia acceso a datos
│   ├── capaControladores/
│   │   └── ControladorMetadatos.go   # Implementación de los 3 RPCs
│   ├── modelos/
│   │   ├── TipoAudio.go
│   │   ├── AudioResumen.go
│   │   ├── Metadatos.go              # Música, Podcast, Audiolibro, RuidoBlanco
│   │   └── RegistroAudio.go
│   ├── serviciosMetadatos/           # Generado por protoc
│   ├── main/servidor.go
│   ├── metadatos.proto
│   └── go.mod
│
├── ServidorDeStreaming/
│   ├── capaAccesoDatos/
│   │   └── RepositorioAudios.go      # Abre archivos mp3 del disco
│   ├── capaFachada/
│   │   └── FachadaStreaming.go       # Lee el archivo y envía chunks
│   ├── capaControladores/
│   │   └── ControladorStreaming.go   # Implementación del RPC ReproducirAudio
│   ├── serviciosStreaming/           # Generado por protoc
│   ├── audios/                       # Archivos mp3 (agregar manualmente)
│   ├── main/servidor.go
│   ├── streaming.proto
│   └── go.mod
│
├── generar_protos.sh
├── compilar.sh
└── README.md
```

---

## Patrones de diseño aplicados

| Patrón | Aplicación |
|--------|-----------|
| **Capas** | capaAccesoDatos → capaFachada → capaControladores → vistas |
| **MVC** | vistas (View), capaControladores (Controller), modelos (Model) |
| **DTO** | Mensajes protobuf actúan como DTOs entre cliente y servidores |

---

## Procedimientos remotos implementados

### ServidorDeAudios (puerto 50052)
| RPC | Descripción |
|-----|-------------|
| `ObtenerTiposDeAudio` | Retorna los 4 tipos: Música, Podcasts, Audiolibros, Ruido Blanco |
| `ObtenerAudiosPorTipo` | Retorna la lista de audios filtrada por tipo |
| `ObtenerDetallesAudio` | Retorna los metadatos completos de un audio por su id |

### ServidorDeStreaming (puerto 50051)
| RPC | Descripción |
|-----|-------------|
| `ReproducirAudio` | Transmite el archivo mp3 en fragmentos de 32 KB (server-side streaming) |
