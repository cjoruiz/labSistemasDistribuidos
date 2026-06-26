package co.edu.unicauca.capaDeControladores;

import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;

public class CallBackImpl extends UnicastRemoteObject implements CallBackInt {
    
    public CallBackImpl() throws RemoteException {
        super();
    }

    public void notificar(PreferenciasDTORespuesta mensaje) throws RemoteException {
        mostrarMensaje(mensaje);
    }

    private void mostrarMensaje(PreferenciasDTORespuesta mensaje) {
        System.out.println("==Preferencias del usuario==");
        System.out.println("\nGeneros");
        mensaje.getPreferenciasGeneros().forEach(genero -> {
            System.out.println(genero.getNombreGenero()+
            "Cantidad de veces escuchado: " + genero.getNumeroPreferencias());
        });
        System.out.println("\nArtistas");
        mensaje.getPreferenciasArtistas().forEach(artista -> {
            System.out.println(artista.getNombreArtista()+
            "Cantidad de veces escuchado: " + artista.getNumeroPreferencias());
        });
    }
}
