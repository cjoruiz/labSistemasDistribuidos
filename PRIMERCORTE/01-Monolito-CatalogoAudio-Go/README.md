# Compilación y Ejecución en Go — Búsqueda de Metadatos de Audio

Aplicación monolítica de consola en Go que permite buscar metadatos de audio en un catálogo precargado en memoria.

## Descripción

El programa mantiene un vector fijo de 5 objetos `MetadataAudio` con datos hardcodeados y ofrece un menú interactivo para buscar audios por título. Demuestra la organización de un proyecto Go con separación en paquetes (`modelos`, `servicios`) y aplicación de getters/setters con campos privados.

## Estructura

```
CompilacionYEjecucionGo/
├── go.mod                        # Definición del módulo Go
├── main.go                       # Punto de entrada y menú por consola
├── modelos/
│   ├── MetadataAudio.go          # Struct con getters/setters (titulo, duracion, tipo, disponible)
│   └── RespuestaMetadataAudioDTO.go  # DTO con código, mensaje y objeto audio
└── servicios/
    └── AudioServices.go          # Carga de datos y búsqueda lineal por título
```

## Compilación y Ejecución

```bash
cd CompilacionYEjecucionGo
go build -o ejecutable main.go
./ejecutable
```

## Uso

```
==== Menu ====
1. Buscar metadata de un audio
2. Salir
Opción:
```

Ingrese el título exacto del audio (ej. `Canción 1`, `Podcats 2`, `Ruido Blanco 3`, `Audiolibro 4`, `Meditación 5`) para ver sus metadatos.

## Catálogo precargado

| Título           | Duración | Tipo                | Disponible |
|------------------|----------|---------------------|------------|
| Canción 1        | 10       | Música              | Sí         |
| Podcats 2        | 20       | Podcats             | No         |
| Ruido Blanco 3   | 30       | Ruido Blanco        | Sí         |
| Audiolibro 4     | 40       | Audiolibros         | Sí         |
| Meditación 5     | 50       | Meditaciones guiadas| No         |

## Conceptos demostrados

- Organización de proyectos Go con múltiples paquetes
- Structs con campos privados y métodos getter/setter
- Patrón DTO para respuestas estandarizadas
- Entrada/salida por consola con `fmt` y `bufio`
- Búsqueda lineal en arreglos
