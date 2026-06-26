package co.edu.unicauca.servidorreacciones.capaControladores;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Controller;

import co.edu.unicauca.servidorreacciones.capaFachadaServices.DTO.ReaccionDTO;
import co.edu.unicauca.servidorreacciones.capaFachadaServices.DTO.ReproduccionDTO;

@Controller
public class ReaccionesController {

    @Autowired
    private SimpMessagingTemplate simpMessagingTemplate;

    @MessageMapping("/enviarReaccion")
    public void processReaccion(ReaccionDTO reaccion) {
        String cancionId = reaccion.getCancionId();
        if (cancionId == null || cancionId.isBlank()) return;
        simpMessagingTemplate.convertAndSend(
                "/brokerDeReacciones/cancion/" + cancionId, reaccion);
    }

    @MessageMapping("/notificarEstado")
    public void processNotification(ReproduccionDTO evento) {
        String cancionId = evento.getCancionId();
        if (cancionId == null || cancionId.isBlank()) return;
        simpMessagingTemplate.convertAndSend(
                "/brokerDeReacciones/cancion/" + cancionId + "/presencia", evento);
    }
}
