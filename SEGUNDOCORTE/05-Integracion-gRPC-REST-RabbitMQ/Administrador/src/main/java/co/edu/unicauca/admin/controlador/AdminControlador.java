package co.edu.unicauca.admin.controlador;

import co.edu.unicauca.admin.modelo.AudioDTO;
import co.edu.unicauca.admin.modelo.AudioRespuestaDTO;
import co.edu.unicauca.admin.modelo.CallbackReproduccionDTO;
import co.edu.unicauca.admin.servicio.ServicioAudio;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.List;

@RestController
@RequestMapping("/api")
public class AdminControlador {

    @Autowired
    private ServicioAudio servicioAudio;

    @PostMapping("/audios")
    public ResponseEntity<AudioRespuestaDTO> almacenarAudio(
            @RequestParam("archivo") MultipartFile archivo,
            @RequestParam("metadata") String metadataJson) {

        try {
            AudioDTO audio = servicioAudio.almacenarAudio(archivo, metadataJson);
            AudioRespuestaDTO respuesta = new AudioRespuestaDTO(
                201,
                "Audio almacenado exitosamente",
                audio
            );
            return ResponseEntity.status(HttpStatus.CREATED).body(respuesta);
        } catch (Exception e) {
            AudioRespuestaDTO respuesta = new AudioRespuestaDTO(
                500,
                "Error al almacenar audio: " + e.getMessage(),
                null
            );
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(respuesta);
        }
    }

    @GetMapping("/audios")
    public ResponseEntity<List<AudioDTO>> listarAudios() {
        List<AudioDTO> audios = servicioAudio.listarAudios();
        return ResponseEntity.ok(audios);
    }

    @GetMapping("/audios/{id}")
    public ResponseEntity<AudioDTO> obtenerAudio(@PathVariable String id) {
        AudioDTO audio = servicioAudio.obtenerAudioPorId(id);
        if (audio != null) {
            return ResponseEntity.ok(audio);
        }
        return ResponseEntity.notFound().build();
    }

    @PostMapping("/callback/reproducir")
    public ResponseEntity<String> recibirCallbackReproduccion(@RequestBody CallbackReproduccionDTO callback) {
        System.out.println(callback.toString());
        return ResponseEntity.ok("Callback recibido");
    }

    @GetMapping("/test")
    public ResponseEntity<String> test() {
        return ResponseEntity.ok("Administrador funcionando en puerto 8080");
    }
}