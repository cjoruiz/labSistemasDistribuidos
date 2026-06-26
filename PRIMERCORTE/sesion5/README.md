# Sesión 5 — gRPC: Búsqueda de Metadatos de Audio

Aplicación cliente-servidor en Go que implementa un servicio de consulta de metadatos de audio mediante **gRPC** con **Protocol Buffers**.

## Arquitectura

```
┌──────────────┐    gRPC (localhost:50053)     ┌──────────────┐
│   Cliente    │ ──── BuscarAudio(request) ───→ │   Servidor   │
│  (consola)   │ ←─── RespuestaMetadataAudioDTO │  (gRPC)      │
└──────────────┘                                └──────────────┘
```

## Estructura

```
sesion5/
└── plantilla/
    ├── cliente/
    │   ├── vistas/
    │   │   └── cliente.go          # Punto de entrada: Llamada RPC BuscarAudio
    │   ├── go.mod                  # Módulo del cliente
    │   └── go.sum
    └── servidor/
        ├── vistas/
        │   └── servidor.go         # Punto de entrada: servidor gRPC puerto 50053
        ├── modelos/
        │   ├── MetadataAudio.go           # Struct con getters/setters
        │   └── RespuestaMetadataAudioDTO.go  # DTO de respuesta
        ├── servicios/
        │   └── AudioServices.go    # Carga de datos y búsqueda lineal
        ├── serviciosCancion/
        │   ├── servicios.proto           # Definición del servicio protobuf
        │   ├── servicios.pb.go           # Mensajes generados por protoc
        │   └── servicios_grpc.pb.go      # Stubs gRPC generados por protoc
        ├── go.mod
        └── go.sum
```

## Servicio gRPC Definido

```protobuf
service serviciosCanciones {
  rpc buscarAudio(PeticionDTO) returns (RespuestaMetadataAudioDTO);
}

message PeticionDTO {
  string titulo = 1;
}

message RespuestaMetadataAudioDTO {
  string mensaje = 1;
  int32 codigo = 2;
  MetadataAudio ObjAudio = 3;
}
```

## Compilación y Ejecución

```bash
# Terminal 1 — Servidor
cd sesion5/plantilla/servidor
go run vistas/servidor.go

# Terminal 2 — Cliente
cd sesion5/plantilla/cliente
go run vistas/cliente.go
```

**Nota:** Si se modificó el `.proto`, regenerar los stubs:
```bash
protoc --go_out=. --go-grpc_out=. servicios.proto
```

## Uso

```
Ingrese el título del audio a buscar: Canción 1

Mensaje: Mëtadata de audio encontrada
Codigo: 200
Audio: Canción 1, Duracion: 10, Tipo: Música, Disponible: true
```

## Catálogo de audios

| Título         | Duración | Tipo                | Disponible |
|----------------|----------|---------------------|------------|
| Canción 1      | 10       | Música              | Sí         |
| Podcats 2      | 20       | Podcats             | No         |
| Ruido Blanco 3 | 30       | Ruido Blanco        | Sí         |
| Audiolibro 4   | 40       | Audiolibros         | Sí         |
| Meditación 5   | 50       | Meditaciones guiadas| No         |

## Conceptos demostrados

- **gRPC** con comunicación remota basada en Protocol Buffers
- **Unary RPC**: llamada solicitud-respuesta simple
- **Generación de código** con `protoc` para Go
- **Contextos** con timeout en llamadas gRPC
- **Arquitectura modular**: modelos, servicios y vistas separados
