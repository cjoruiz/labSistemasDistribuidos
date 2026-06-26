package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.stereotype.Service;

import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.commons.Simulacion;

import org.springframework.amqp.rabbit.annotation.RabbitListener;

@Service
public class MessageConsumer {

    @RabbitListener(queues = "notificaciones_audios")
    public void receiveMessage(CancionAlmacenarDTOInput objCancion) {
        System.out.println("═══════════════════════════════════════════════════════════");
        System.out.println("              NUEVO AUDIO REGISTRADO                         ");
        System.out.println("═══════════════════════════════════════════════════════════");
        System.out.println("Datos de la canción recibidos");
        System.out.println("Enviando correo electrónico...");
        
        Simulacion.simular(3000, "Enviando correo...");
        
        System.out.println("");
        System.out.println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
        System.out.println("                    CORREO ENVIADO                          ");
        System.out.println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━");
        System.out.println("De:       Sistema de Audio <noreply@sistemaaudio.com>");
        System.out.println("Para:     usuarios@sistemaaudio.com");
        System.out.println("Asunto:   Nuevo audio disponible en la plataforma");
        System.out.println("");
        System.out.println("────────────────────────────────────────────────────────────");
        System.out.println("                    CONTENIDO DEL CORREO                   ");
        System.out.println("────────────────────────────────────────────────────────────");
        System.out.println("Estimado usuario,");
        System.out.println("");
        System.out.println("Se ha registrado un nuevo audio en nuestra plataforma:");
        System.out.println("");
        System.out.println("   ID del Audio:      " + objCancion.getId());
        System.out.println("   Título:           " + objCancion.getTitulo());
        System.out.println("   Artista:          " + objCancion.getArtista());
        System.out.println("   Género:           " + objCancion.getGenero());
        System.out.println("   Fecha de Registro: " + objCancion.getFechaRegistro());
        System.out.println("");
        System.out.println("────────────────────────────────────────────────────────────");
        System.out.println("                    FRASE MOTIVADORA                        ");
        System.out.println("────────────────────────────────────────────────────────────");
        System.out.println("  _ " + objCancion.getFraseMotivadora());
        System.out.println("");
        System.out.println("────────────────────────────────────────────────────────────");
        System.out.println("Gracias por usar nuestro sistema de audio.");
        System.out.println("═══════════════════════════════════════════════════════════");
        System.out.println("");
    }
}
    