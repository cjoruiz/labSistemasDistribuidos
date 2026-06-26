# AGENTS.md - Sistema de Administración de Audio Distribuido

## Estructura del Proyecto

```
prueba/
├── Administrador/                    # Java + Spring Boot + gRPC
├── ServidorDeStreaming/             # Go + gRPC (Puerto 50052/8091)
├── ServidorDeAudiosMetadatos/       # Go + gRPC/REST (Puerto 50051/8090)
├── Cliente/                          # Cliente Go existente
└── microservicioEnvioCorreos - plantilla/  # Java + Spring Boot (Puerto 6000)
```

## Compilación

### Servidores Go
```bash
cd ServidorDeStreaming && go build -o servidor main/servidor.go
cd ServidorDeAudiosMetadatos && go build -o servidor main/servidor.go
cd Cliente && go build -o cliente main/cliente.go
```

### Administrador Java (incluye generación de stubs gRPC)
```bash
cd Administrador && mvn clean package
```

### Servidor de Correos Java
```bash
cd "microservicioEnvioCorreos - plantilla" && mvn clean package
```

## Puertos

| Servicio | Puerto gRPC | Puerto REST/HTTP |
|----------|-------------|------------------|
| ServidorDeStreaming | 50052 | 8091 |
| ServidorDeAudiosMetadatos | 50051 | 8090 |
| Administrador | - | 8080 (dinámico con --server.port=0) |
| Servidor de Correos | - | 6000 |

## Dependencias Externas

- **RabbitMQ**: debe estar ejecutándose en `localhost:5672` (admin:1234)
- **protoc**: instalado en `/usr/bin/protoc` para generar código Java desde archivos .proto
- **protoc-gen-grpc-java**: descargado automáticamente por Maven

## Ejecución de Servicios

### Orden de inicio recomendado

1. **RabbitMQ** (si no está corriendo)
2. **ServidorDeAudiosMetadatos** (Go)
   ```bash
   cd ServidorDeAudiosMetadatos && ./servidor
   ```

3. **ServidorDeStreaming** (Go)
   ```bash
   cd ServidorDeStreaming && ./servidor
   ```

4. **Servidor de Correos** (Java)
   ```bash
   cd "microservicioEnvioCorreos - plantilla" && mvn spring-boot:run
   ```

5. **Administrador** (Java)
   ```bash
   cd Administrador && java -jar target/administrador-1.0.0.jar
   ```

### Ejecutar múltiples Administradores

Para ejecutar varios administradores en diferentes puertos:
```bash
# Administrador 1 (puerto 8080 por defecto)
cd Administrador && java -jar target/administrador-1.0.0.jar

# Administrador 2 (puerto dinámico)
cd Administrador && java -jar target/administrador-1.0.0.jar --server.port=0
```

## Arquitectura del Sistema

```
┌──────────┐   POST/multipart   ┌────────────┐    gRPC      ┌──────────────────┐
│ Postman  │───────────────────▶│ ADMINISTRADOR│◀───────────▶│ Serv.DeStreaming │
│          │                    │   (Java)    │  50052      │     (Go)         │
└──────────┘                    └──────┬──────┘              └────────┬─────────┘
                                       │  REST                      │
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
         (cuando cliente reproduce un audio)
```

## Flujo de Callback (Notificaciones)

### Cuándo se envía el callback
Cuando el Cliente reproduce un audio, el ServidorDeStreaming envía un callback a **todos** los administradores registrados con:
- **ID del audio**: nombre del archivo reproducido
- **Fecha/Hora**: momento de la reproducción (formato: YYYY-MM-DD HH:MM:SS)

### Registro de callback
1. Al iniciar, cada Administrador registra su URL de callback en el ServidorDeStreaming via gRPC
2. El ServidorDeStreaming guarda las URLs en una lista (soporta múltiples administradores)
3. Cuando se reproduce un audio, se envía el callback a todas las URLs registradas

### Endpoint del callback
- **URL**: `http://localhost:{puerto}/api/callback/reproducir`
- **Método**: POST
- **Body JSON**:
```json
{
  "idAudio": "nombre_archivo.mp3",
  "tituloAudio": "nombre_archivo.mp3",
  "fechaHoraReproduccion": "2026-05-19 01:25:02"
}
```

## Endpoints del Administrador (Puerto 8080 o dinámico)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/test` | Verificar que el servidor funciona |
| POST | `/api/audios` | Almacenar nuevo audio (multipart) |
| GET | `/api/audios` | Listar todos los audios |
| GET | `/api/audios/{id}` | Ver detalle de un audio |
| POST | `/api/callback/reproducir` | Receibir callback de reproducción |

## Almacenar Audio con Postman

1. POST a `http://localhost:8080/api/audios`
2. Body type: `multipart/form-data`
3. Agregar dos campos:
   - `archivo`: File (seleccionar archivo de audio)
   - `metadata`: Text (JSON con los metadatos)

### Ejemplo metadata JSON:
```json
{
  "titulo": "Mi canción",
  "autor": "Artista X",
  "album": "Álbum 1",
  "genero": "Rock",
  "duracion": 180,
  "tipoId": 1,
  "tipoNombre": "Música",
  "fechaLanzamiento": "2024-01-15",
  "disponible": true,
  "selloDiscografico": "EMI"
}
```

### IMPORTANTE: tipoNombre debe ser exacto
El campo `tipoNombre` debe coincidir exactamente con:
- "Música" (con tilde, M mayúscula)
- "Podcast"
- "Audiolibro"
- "Ruido Blanco"

Si envías "Musica" o "musica" no aparecerá en las listas del cliente.

### TipoId válido:
- 1 = Música
- 2 = Podcast
- 3 = Audiolibro
- 4 = Ruido Blanco

## Archivos Proto y gRPC

### Proto del Administrador
Ubicación: `Administrador/src/main/proto/audio_streaming.proto`

El Administrador usa este proto para comunicarse con el ServidorDeStreaming via gRPC. Los stubs se generan automáticamente con:
```bash
mvn protobuf:compile
mvn protobuf:compile-custom
```

### Proto del ServidorDeStreaming
Ubicación: `ServidorDeStreaming/protos/audio_streaming.proto`

Servicios definidos:
- `StreamAudio` - Streaming de audio al cliente
- `AlmacenarAudio` - Guardar audio desde el administrador
- `RegistrarCallback` - Registrar URL de callback para notificaciones

## Issues Conocidos y Soluciones

### Cliente no compila
**Problema**: go.mod referencia `ServidorDeAudios` pero el directorio se llama `ServidorDeAudiosMetadatos`.
**Solución**: Ya está resuelto en `Cliente/go.mod` línea 26:
```
replace ServidorDeAudios => ../ServidorDeAudiosMetadatos
```

### MaxUploadSizeExceededException
**Problema**: Spring Boot limita uploads a 1MB por defecto.
**Solución**: Ya está configurado en `Administrador/src/main/resources/application.properties`:
```properties
spring.servlet.multipart.max-file-size=50MB
spring.servlet.multipart.max-request-size=50MB
```

### Callback no llega a los administradores
**Problema**: El ServidorDeStreaming solo guardaba una URL de callback.
**Solución**: Modificado para mantener una lista de URLs y enviar callback a todos los administradores registrados.

### Administrador con puerto dinámico no registra callback correctamente
**Problema**: La URL de callback estaba hardcodeada a `localhost:8080`.
**Solución**: Modificado `AdministradorApplication.java` para obtener el puerto dinámicamente.

## Notas Importantes

1. **Orden de servicios**: El Administrador debe iniciarse DESPUÉS del ServidorDeStreaming para que el registro de callback funcione correctamente.

2. **Puerto dinámico del Administrador**: Si se usa `--server.port=0`, Spring Boot asigna un puerto libre automáticamente. El código detecta este puerto para registrar la URL de callback correcta.

3. **Callback a múltiples administradores**: El ServidorDeStreaming envía el callback a TODOS los administradores registrados, no solo a uno.

4. **Cliente Go**: Necesita ejecutarse en una terminal separada. Usa los servidores gRPC directamente (no el Administrador).

5. **RabbitMQ**: Necesario para la comunicación entre ServidorDeMetadatos y ServidorDeCorreos (notificaciones por correo).

6. **Cliente Go**: Al iniciar, solicita nickname y contraseña. Usuarios hardcodeados en `Cliente/vistas/menu.go`:
   - `admin` / `admin123`
   - `juan` / `juan123`
   - `maria` / `maria123`
   - `carlos` / `carlos123`

## Estructura de Capas (ServidorDeAudiosMetadatos)

El servidor fue refactorizado para seguir arquitectura por capas:
```
ServidorDeAudiosMetadatos/
├── main/
│   └── servidor.go              # Main limpio, solo inicializa
├── capaControladores/
│   ├── ControladorAudios.go    # gRPC handlers
│   └── ControladorHTTP.go     # REST handlers
├── capaAccesoDatos/
│   └── repositorioAudios.go    # Datos en memoria
├── capaFachada/
│   └── fachadaServicios.go     # Lógica de negocio
├── modelos/
│   └── MetadataAudio.go        # Estructura de datos
└── componenteConexionCola/     # NUEVO: Lógica de RabbitMQ separada
    └── publicadorRabbit.go
```

## Issues Conocidos y Soluciones (Actualizado)

### Nuevo audio no aparece en el Cliente
**Problema**: El campo `tipoNombre` debe coincidir exactamente ("Música", no "Musica").
**Solución**: Asegurar que el JSON enviado desde Postman tenga `"tipoNombre": "Música"` con tilde.