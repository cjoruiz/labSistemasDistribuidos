package co.edu.unicauca.admin;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationContext;
import org.springframework.core.env.Environment;

import co.edu.unicauca.admin.servicio.ClienteStreaming;

@SpringBootApplication
public class AdministradorApplication {

    public static void main(String[] args) {
        ApplicationContext contexto = SpringApplication.run(AdministradorApplication.class, args);
        
        ClienteStreaming clienteStreaming = contexto.getBean(ClienteStreaming.class);
        
        // Obtener el puerto dinámico del servidor
        Environment env = contexto.getEnvironment();
        String puertoStr = env.getProperty("server.port", "8080");
        int puerto = 8080;
        try {
            puerto = Integer.parseInt(puertoStr);
            if (puerto == 0) {
                // Si es 0, buscar el puerto asignado dinámicamente
                String puertoAsignado = env.getProperty("local.server.port", "8080");
                puerto = Integer.parseInt(puertoAsignado);
            }
        } catch (NumberFormatException e) {
            puerto = 8080;
        }
        
        String urlCallback = "http://localhost:" + puerto + "/api/callback/reproducir";
        System.out.println("Registrando callback en servidor de streaming: " + urlCallback);
        boolean registrado = clienteStreaming.registrarCallback(urlCallback);
        
        System.out.println("===========================================");
        System.out.println("   ADMINISTRADOR DE AUDIO INICIADO        ");
        System.out.println("===========================================");
        System.out.println("Puerto REST: " + puerto);
        System.out.println("Esperando solicitudes de Postman...");
        System.out.println("Callback registrado: " + (registrado ? "SI" : "NO"));
    }
}