package co.edu.unicauca.servidorreacciones.capaFachadaServices.DTO;

public class ReproduccionDTO {
    private String nickname;
    private String cancionId;
    private String accion; // "play" o "stop"

    public ReproduccionDTO() {}

    public String getNickname() { return nickname; }
    public void setNickname(String nickname) { this.nickname = nickname; }

    public String getCancionId() { return cancionId; }
    public void setCancionId(String cancionId) { this.cancionId = cancionId; }

    public String getAccion() { return accion; }
    public void setAccion(String accion) { this.accion = accion; }
}
