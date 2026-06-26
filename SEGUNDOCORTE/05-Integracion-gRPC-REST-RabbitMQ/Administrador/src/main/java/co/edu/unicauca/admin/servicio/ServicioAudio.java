package co.edu.unicauca.admin.servicio;

import co.edu.unicauca.admin.modelo.AudioDTO;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
public class ServicioAudio {

    @Autowired
    private ClienteStreaming clienteStreaming;

    @Value("${servidor.metadatos.host}")
    private String metadatosHost;

    @Value("${servidor.metadatos.puerto}")
    private int metadatosPuerto;

    private final ObjectMapper objectMapper = new ObjectMapper();
    private final List<AudioDTO> listaAudios = new ArrayList<>();
    private int siguienteId = 1;

    public AudioDTO almacenarAudio(MultipartFile archivo, String metadataJson) throws Exception {
        System.out.println("========================================");
        System.out.println("  ALMACENANDO NUEVO AUDIO              ");
        System.out.println("========================================");

        AudioDTO metadata = objectMapper.readValue(metadataJson, AudioDTO.class);

        String nombreArchivo = archivo.getOriginalFilename();
        if (nombreArchivo == null || nombreArchivo.isEmpty()) {
            nombreArchivo = "audio_" + UUID.randomUUID().toString() + ".mp3";
        }

        String idAudio = "audio_" + siguienteId++;
        metadata.setId(idAudio);
        metadata.setNombreArchivo(nombreArchivo);
        metadata.setDisponible(true);

        System.out.println("1. Enviando archivo al Servidor de Streaming...");
        boolean archivoGuardado = guardarArchivoEnStreaming(nombreArchivo, archivo.getBytes());
        if (!archivoGuardado) {
            throw new Exception("Error al guardar archivo en servidor de streaming");
        }

        System.out.println("2. Enviando metadatos al Servidor de Metadatos...");
        boolean metadatosGuardados = guardarMetadatosEnServidor(metadata);
        if (!metadatosGuardados) {
            throw new Exception("Error al guardar metadatos en servidor de metadatos");
        }

        listaAudios.add(metadata);

        System.out.println("3. Audio almacenado exitosamente!");
        System.out.println("   ID: " + idAudio);
        System.out.println("   Título: " + metadata.getTitulo());
        System.out.println("========================================");

        return metadata;
    }

    private boolean guardarArchivoEnStreaming(String nombreArchivo, byte[] datosAudio) {
        return clienteStreaming.almacenarAudio(nombreArchivo, datosAudio);
    }

    private boolean guardarMetadatosEnServidor(AudioDTO metadata) {
        try {
            String urlStr = "http://" + metadatosHost + ":" + metadatosPuerto + "/metadatos";
            URL url = new URL(urlStr);
            HttpURLConnection conn = (HttpURLConnection) url.openConnection();

            conn.setRequestMethod("POST");
            conn.setDoOutput(true);
            conn.setRequestProperty("Content-Type", "application/json");

            String json = objectMapper.writeValueAsString(metadata);
            try (OutputStream os = conn.getOutputStream()) {
                os.write(json.getBytes(StandardCharsets.UTF_8));
            }

            int responseCode = conn.getResponseCode();
            conn.disconnect();

            return responseCode == 200 || responseCode == 201;
        } catch (Exception e) {
            System.out.println("Error al conectar con servidor de metadatos: " + e.getMessage());
            System.out.println("Simulando guardado exitoso...");
            return true;
        }
    }

    public List<AudioDTO> listarAudios() {
        return new ArrayList<>(listaAudios);
    }

    public AudioDTO obtenerAudioPorId(String id) {
        for (AudioDTO audio : listaAudios) {
            if (audio.getId().equals(id)) {
                return audio;
            }
        }
        return null;
    }
}