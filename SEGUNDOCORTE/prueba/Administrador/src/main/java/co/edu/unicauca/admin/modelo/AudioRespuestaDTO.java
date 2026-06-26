package co.edu.unicauca.admin.modelo;

public class AudioRespuestaDTO {

    private Integer codigo;
    private String mensaje;
    private AudioDTO audio;

    public AudioRespuestaDTO() {
    }

    public AudioRespuestaDTO(Integer codigo, String mensaje, AudioDTO audio) {
        this.codigo = codigo;
        this.mensaje = mensaje;
        this.audio = audio;
    }

    public Integer getCodigo() {
        return codigo;
    }

    public void setCodigo(Integer codigo) {
        this.codigo = codigo;
    }

    public String getMensaje() {
        return mensaje;
    }

    public void setMensaje(String mensaje) {
        this.mensaje = mensaje;
    }

    public AudioDTO getAudio() {
        return audio;
    }

    public void setAudio(AudioDTO audio) {
        this.audio = audio;
    }
}