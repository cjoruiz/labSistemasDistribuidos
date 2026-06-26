package co.edu.unicauca.capaDeControladores;

import java.rmi.Remote;
import java.rmi.RemoteException;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;


public interface CallBackInt extends Remote {
    public void notificar(PreferenciasDTORespuesta mensaje) throws RemoteException;    
}
