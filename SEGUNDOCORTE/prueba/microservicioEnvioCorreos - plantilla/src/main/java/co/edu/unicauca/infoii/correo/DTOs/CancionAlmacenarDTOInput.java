package co.edu.unicauca.infoii.correo.DTOs;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class CancionAlmacenarDTOInput {
    @JsonProperty("id")
    private String id;
    @JsonProperty("titulo")
    private String titulo;
    @JsonProperty("artista")
    private String artista;
    @JsonProperty("genero")
    private String genero;
    @JsonProperty("fechaRegistro")
    private String fechaRegistro;
    @JsonProperty("fraseMotivadora")
    private String fraseMotivadora;
}
	