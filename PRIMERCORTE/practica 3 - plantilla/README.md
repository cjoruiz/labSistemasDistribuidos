# Práctica 3 — Cliente-Servidor TCP para Búsqueda de Metadatos de Audio

Aplicación cliente-servidor en Go que implementa un sistema de consulta de metadatos de audio mediante **sockets TCP**, con manejo concurrente de peticiones mediante goroutines.

## Arquitectura

```
┌──────────────┐       TCP (localhost:9000)       ┌──────────────┐
│   Cliente    │ ─────── (request: título) ──────→ │   Servidor   │
│  (consola)   │ ←────── (response: JSON) ─────── │  (gorutinas) │
└──────────────┘                                   └──────────────┘
```

## Estructura

```
practica 3 - plantilla/
├── cliente/
│   ├── go.mod                   # Módulo del cliente (depende de modelos del servidor)
│   └── cliente.go               # Punto de entrada: envía título, recibe y muestra respuesta
└── servidor/
    ├── go.mod                   # Módulo del servidor
    ├── servidor.go              # Listener TCP + manejo concurrente de conexiones
    ├── ejecutableServidor       # Binario precompilado
    ├── modelos/
    │   ├── MetadataAudio.go           # Struct con getters/setters + Marshal/Unmarshal JSON
    │   └── RespuestaMetadataAudioDTO.go  # DTO de respuesta
    └── servicios/
        └── AudioServices.go     # Carga de datos y búsqueda lineal
```

## Comunicación

- **Protocolo:** TCP
- **Dirección:** `localhost:9000`
- **Request:** Título del audio (string plano con salto de línea)
- **Response:** JSON con `RespuestaMetadataAudioDTO` (código, mensaje, objeto audio)

## Compilación y Ejecución

```bash
# Terminal 1 — Servidor
cd servidor
go build -o servidor servidor.go
./servidor

# Terminal 2 — Cliente
cd cliente
go build -o cliente cliente.go
./cliente
```

## Uso del Cliente

```
Buscar metadata de un audio:
> Canción 1

 Título del audio: Canción 1
 Tamaño de audio: 10
 Tipo de audio: Música
 El audio está disponible: true
```

## Catálogo de audios

| Título         | Duración | Tipo        | Disponible |
|----------------|----------|-------------|------------|
| Canción 1      | 10       | Música      | Sí         |
| Podcasts 2     | 20       | Podcasts    | No         |
| Ruido Blanco 3 | 30       | Ruido Blanco| Sí         |
| Audiolibro 4   | 40       | Audiolibros | Sí         |
| Meditación 5   | 50       | Meditaciones| No         |

## Conceptos demostrados

- **Sockets TCP** con `net.Listen` / `net.Dial`
- **Concurrencia** con goroutines para atender múltiples clientes simultáneamente
- **Serialización JSON** para intercambio de datos estructurados
- **Módulos Go** con dependencias locales usando `replace` en `go.mod`
- **Simulación de procesamiento** con `time.Sleep` para evidenciar concurrencia
