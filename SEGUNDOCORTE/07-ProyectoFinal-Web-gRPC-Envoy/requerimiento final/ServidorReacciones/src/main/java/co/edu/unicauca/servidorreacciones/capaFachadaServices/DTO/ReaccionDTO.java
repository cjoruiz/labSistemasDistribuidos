package co.edu.unicauca.servidorreacciones.capaFachadaServices.DTO;

public class ReaccionDTO {
    private String nickname;
    private String cancionId;
    private String reaccion;

    public ReaccionDTO() {}

    public String getNickname() { return nickname; }
    public void setNickname(String nickname) { this.nickname = nickname; }

    public String getCancionId() { return cancionId; }
    public void setCancionId(String cancionId) { this.cancionId = cancionId; }

    public String getReaccion() { return reaccion; }
    public void setReaccion(String reaccion) { this.reaccion = reaccion; }
}
