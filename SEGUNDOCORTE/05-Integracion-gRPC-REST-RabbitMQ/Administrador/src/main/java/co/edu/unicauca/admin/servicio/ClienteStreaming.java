package co.edu.unicauca.admin.servicio;

import co.edu.unicauca.admin.proto.AudioFileRequest;
import co.edu.unicauca.admin.proto.AudioFileResponse;
import co.edu.unicauca.admin.proto.CallbackRegistroRequest;
import co.edu.unicauca.admin.proto.CallbackRegistroResponse;
import co.edu.unicauca.admin.proto.StreamingServiceGrpc;
import com.google.protobuf.ByteString;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import org.springframework.stereotype.Service;

import java.util.concurrent.TimeUnit;

@Service
public class ClienteStreaming {

    private static final String HOST = "localhost";
    private static final int PUERTO = 50052;

    public boolean almacenarAudio(String filename, byte[] datosAudio) {
        System.out.println("Enviando audio al servidor de streaming via gRPC...");
        
        ManagedChannel channel = null;
        try {
            channel = ManagedChannelBuilder.forAddress(HOST, PUERTO)
                    .usePlaintext()
                    .build();

            StreamingServiceGrpc.StreamingServiceBlockingStub stub = 
                    StreamingServiceGrpc.newBlockingStub(channel);

            AudioFileRequest request = AudioFileRequest.newBuilder()
                    .setFilename(filename)
                    .setData(ByteString.copyFrom(datosAudio))
                    .build();

            AudioFileResponse response = stub.almacenarAudio(request);
            
            if (response.getSuccess()) {
                System.out.println("Audio almacenado exitosamente: " + response.getMessage());
                return true;
            } else {
                System.out.println("Error al almacenar audio: " + response.getMessage());
                return false;
            }

        } catch (Exception e) {
            System.out.println("Error en cliente streaming: " + e.getMessage());
            return false;
        } finally {
            if (channel != null) {
                try {
                    channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            }
        }
    }

    public boolean registrarCallback(String urlCallback) {
        System.out.println("Registrando callback en servidor de streaming: " + urlCallback);
        
        ManagedChannel channel = null;
        try {
            channel = ManagedChannelBuilder.forAddress(HOST, PUERTO)
                    .usePlaintext()
                    .build();

            StreamingServiceGrpc.StreamingServiceBlockingStub stub = 
                    StreamingServiceGrpc.newBlockingStub(channel);

            CallbackRegistroRequest request = CallbackRegistroRequest.newBuilder()
                    .setCallbackUrl(urlCallback)
                    .build();

            CallbackRegistroResponse response = stub.registrarCallback(request);
            
            System.out.println("Callback registrado exitosamente: " + response.getSuccess());
            return response.getSuccess();

        } catch (Exception e) {
            System.out.println("Error al registrar callback: " + e.getMessage());
            return false;
        } finally {
            if (channel != null) {
                try {
                    channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
            }
        }
    }
}