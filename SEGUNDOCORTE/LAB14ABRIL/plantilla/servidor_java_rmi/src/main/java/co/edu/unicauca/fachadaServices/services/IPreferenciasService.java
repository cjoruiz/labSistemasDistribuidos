
package co.edu.unicauca.fachadaServices.services;

import java.rmi.RemoteException;

import co.edu.unicauca.capaDeControladores.CallBackInt;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;


public interface IPreferenciasService {

    public PreferenciasDTORespuesta getReferencias(Integer id) throws RemoteException;
	public Boolean registrarReferenciaAdministrador(CallBackInt objRemotoCliente);
}


