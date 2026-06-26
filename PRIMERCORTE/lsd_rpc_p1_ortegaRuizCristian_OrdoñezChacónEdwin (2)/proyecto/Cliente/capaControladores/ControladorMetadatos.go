// Package capacontroladores del cliente encapsula las llamadas RPC al ServidorDeAudios.
// Traduce las respuestas protobuf a estructuras de dominio del cliente.
package capacontroladores

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	pbMeta "servidor.audios.local/grpc-servidor-audios/serviciosMetadatos"
)

// ClienteMetadatos agrupa la conexión y el stub gRPC al ServidorDeAudios.
type ClienteMetadatos struct {
	conexion *grpc.ClientConn
	stub     pbMeta.ServicioMetadatosClient
}

// NuevoClienteMetadatos crea y devuelve un ClienteMetadatos conectado a la dirección indicada.
func NuevoClienteMetadatos(direccion string) (*ClienteMetadatos, error) {
	conn, err := grpc.Dial(direccion, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al ServidorDeAudios: %w", err)
	}
	return &ClienteMetadatos{
		conexion: conn,
		stub:     pbMeta.NewServicioMetadatosClient(conn),
	}, nil
}

// Cerrar libera la conexión gRPC.
func (c *ClienteMetadatos) Cerrar() {
	c.conexion.Close()
}

// contextoConTimeout crea un contexto con timeout estándar para las llamadas RPC.
func contextoConTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

// ObtenerTiposDeAudio invoca el RPC ObtenerTiposDeAudio y retorna la lista de tipos.
func (c *ClienteMetadatos) ObtenerTiposDeAudio() ([]*pbMeta.TipoAudio, error) {
	ctx, cancel := contextoConTimeout()
	defer cancel()
	resp, err := c.stub.ObtenerTiposDeAudio(ctx, &pbMeta.SolicitudVacia{})
	if err != nil {
		return nil, fmt.Errorf("RPC ObtenerTiposDeAudio falló: %w", err)
	}
	return resp.GetTipos(), nil
}

// ObtenerAudiosPorTipo invoca el RPC ObtenerAudiosPorTipo y retorna la lista de audios.
func (c *ClienteMetadatos) ObtenerAudiosPorTipo(idTipo int32) ([]*pbMeta.AudioResumen, error) {
	ctx, cancel := contextoConTimeout()
	defer cancel()
	resp, err := c.stub.ObtenerAudiosPorTipo(ctx, &pbMeta.SolicitudTipo{IdTipo: idTipo})
	if err != nil {
		return nil, fmt.Errorf("RPC ObtenerAudiosPorTipo falló: %w", err)
	}
	return resp.GetAudios(), nil
}

// ObtenerDetallesAudio invoca el RPC ObtenerDetallesAudio y retorna los metadatos del audio.
func (c *ClienteMetadatos) ObtenerDetallesAudio(idAudio int32) (*pbMeta.RespuestaDetallesAudio, error) {
	ctx, cancel := contextoConTimeout()
	defer cancel()
	resp, err := c.stub.ObtenerDetallesAudio(ctx, &pbMeta.SolicitudAudio{IdAudio: idAudio})
	if err != nil {
		return nil, fmt.Errorf("RPC ObtenerDetallesAudio falló: %w", err)
	}
	return resp, nil
}
