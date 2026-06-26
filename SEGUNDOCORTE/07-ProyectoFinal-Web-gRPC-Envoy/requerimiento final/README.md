# Sistema Distribuido de Audio

## Requisitos

- Go (ServidorStreaming, ServidorAudiosMetadatos)
- Java 17+ y Maven (ServidorReacciones)
- Envoy proxy
- Node.js y npm (solo para regenerar bundle.js)
- Navegador web moderno

## Puertos

| Servicio | Puerto |
|---|---|
| ServidorStreaming (gRPC) | 50051 |
| Proxy Envoy (gRPC-Web) | 8080 |
| ServidorAudiosMetadatos (REST) | 8090 |
| ServidorReacciones (STOMP/WS) | 5000 |

## Orden de inicio

### 1. ServidorStreaming

```bash
cd ServidorStreaming
go run main/servidor.go
```

### 2. ServidorAudiosMetadatos

```bash
cd ServidorAudiosMetadatos
go run .
```

### 3. ServidorReacciones

```bash
cd ServidorReacciones
mvn clean package -q
java -jar target/servidorreacciones-0.0.1-SNAPSHOT.jar
```

### 4. Proxy Envoy

```bash
cd ProxyEnvoy
envoy -c envoyConfig2.yaml
```

### 5. Cliente Web

Abrir `clienteweb/index.html` como archivo local (`file://`), **no** desde Live Server.

### 6. Administrador Web

Abrir `adminweb/index.html` como archivo local (`file://`).

## Uso

### Cliente
1. Escribe un nickname y presiona "Conectar Reacciones"
2. Selecciona un tipo de audio (Música, Podcast, Audiolibro, Ruido Blanco)
3. Selecciona un audio de la lista
4. Presiona "Reproducir" para escucharlo vía streaming
5. Envía reacciones (like, heart, sad, fun) — solo llegan a quienes escuchan el mismo audio
6. Al pausar/detener, otros usuarios ven tu estado

### Administrador
1. Selecciona el tipo de audio
2. Completa los metadatos específicos del tipo
3. Selecciona un archivo MP3
4. Presiona "Registrar Audio"

## Recompilar bundle.js (si se modifica `cliente_receptor_fragmentos.js`)

```bash
cd clienteweb
npx webpack
```
