
package co.edu.unicauca.fachadaServices.services;

import java.util.ArrayList;
import java.util.List;

import co.edu.unicauca.capaDeControladores.CallBackInt;
import co.edu.unicauca.fachadaServices.DTO.CancionDTOEntrada;
import co.edu.unicauca.fachadaServices.DTO.PreferenciasDTORespuesta;
import co.edu.unicauca.fachadaServices.DTO.ReproduccionesDTOEntrada;
import co.edu.unicauca.fachadaServices.services.componenteCalculaPreferencias.CalculadorPreferencias;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorCanciones.ComunicacionServidorCanciones;
import co.edu.unicauca.fachadaServices.services.componenteComunicacionServidorReproducciones.ComunicacionServidorReproducciones;

public class PreferenciasServiceImpl implements IPreferenciasService {
	
	private ComunicacionServidorCanciones comunicacionServidorCanciones;
	private ComunicacionServidorReproducciones comunicacionServidorReproducciones;
	private CalculadorPreferencias calculadorPreferencias;
	private List<CallBackInt> administradoresRegistrados;

	public PreferenciasServiceImpl() {
		this.comunicacionServidorCanciones = new ComunicacionServidorCanciones();
		this.comunicacionServidorReproducciones = new ComunicacionServidorReproducciones();
		this.calculadorPreferencias = new CalculadorPreferencias();
		this.administradoresRegistrados = new ArrayList<>();
	}

	@Override
	public PreferenciasDTORespuesta getReferencias(Integer id) {
		System.out.println("Obteniendo preferencias para el usuario con ID: " + id);
		List<CancionDTOEntrada> objCanciones= this.comunicacionServidorCanciones.obtenerCancionesRemotas();
		System.out.println("Canciones obtenidas del dervidor de canciones:");
		for(CancionDTOEntrada cancion : objCanciones) {
			System.out.println("Cancion obtenida: " + cancion.getTitulo());
			System.out.println("Genero: " + cancion.getGenero());
			System.out.println("Artista: " + cancion.getArtista());
		}

		List<ReproduccionesDTOEntrada> reproduccionesUsuario = 
		this.comunicacionServidorReproducciones.obtenerReproduccionesRemotas(id);
		System.out.println("Reproducciones obtenidas del servidor de reproducciones para el usuario: " + id);
		for(ReproduccionesDTOEntrada reproduccion : reproduccionesUsuario) {
			System.out.println(reproduccion.getIdUsuario() + " "  + reproduccion.getIdCancion());
		}

		 PreferenciasDTORespuesta preferencias = this.calculadorPreferencias.calcular(id, objCanciones, reproduccionesUsuario);
		 notificarAdministradores(preferencias);	
		return preferencias;
	}

	public Boolean registrarReferenciaAdministrador(CallBackInt objRemotoCliente) {
		if (objRemotoCliente != null && !administradoresRegistrados.contains(objRemotoCliente)) {
			administradoresRegistrados.add(objRemotoCliente);
			return true;
		}
		return false;
	}
	
	private void notificarAdministradores(PreferenciasDTORespuesta mensaje) {
		for (CallBackInt admin : administradoresRegistrados) {
			try {
				admin.notificar(mensaje);
			} catch (Exception e) {
				System.out.println("Error al notificar al administrador: " + e.getMessage());
			}
		}
	}

	
}
