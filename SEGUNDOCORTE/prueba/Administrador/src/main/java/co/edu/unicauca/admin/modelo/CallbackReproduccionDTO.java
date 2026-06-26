package co.edu.unicauca.admin.modelo;

public class CallbackReproduccionDTO {

    private String idAudio;
    private String tituloAudio;
    private String fechaHoraReproduccion;

    public CallbackReproduccionDTO() {
    }

    public CallbackReproduccionDTO(String idAudio, String tituloAudio, String fechaHoraReproduccion) {
        this.idAudio = idAudio;
        this.tituloAudio = tituloAudio;
        this.fechaHoraReproduccion = fechaHoraReproduccion;
    }

    public String getIdAudio() {
        return idAudio;
    }

    public void setIdAudio(String idAudio) {
        this.idAudio = idAudio;
    }

    public String getTituloAudio() {
        return tituloAudio;
    }

    public void setTituloAudio(String tituloAudio) {
        this.tituloAudio = tituloAudio;
    }

    public String getFechaHoraReproduccion() {
        return fechaHoraReproduccion;
    }

    public void setFechaHoraReproduccion(String fechaHoraReproduccion) {
        this.fechaHoraReproduccion = fechaHoraReproduccion;
    }

    @Override
    public String toString() {
        return "═══════════════════════════════════════════════════════════\n" +
               "           NUEVA REPRODUCCIÓN DETECTADA                   \n" +
               "═══════════════════════════════════════════════════════════\n" +
               "  ID Audio:       " + idAudio + "\n" +
               "  Título:         " + tituloAudio + "\n" +
               "  Fecha/Hora:     " + fechaHoraReproduccion + "\n" +
               "═══════════════════════════════════════════════════════════";
    }
}