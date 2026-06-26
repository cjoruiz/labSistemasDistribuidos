package co.edu.unicauca.infoii.correo.DTOs;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class CancionAlmacenarDTOInput {
    private Integer id;
    @JsonProperty("Titulo")
    private String titulo;
    @JsonProperty("Artista")
    private String artista;
    @JsonProperty("Genero")
    private String genero;
    private String correoUsuario;
}
	