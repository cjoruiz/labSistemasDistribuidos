# LAB14ABRIL - Sistema de Preferencias con Callbacks vía RMI

## Descripción
Extensión del laboratorio LAB7ABRIL que incorpora un **módulo administrador** con notificaciones push (callbacks) mediante RMI. Cuando un cliente consulta preferencias, el servidor notifica automáticamente a todos los administradores registrados.

## Arquitectura
```
Cliente (RMI) <--> Servidor RMI <--> json-server
                        ^
                        | (callback RMI)
                        v
                 Administrador
```

## Componentes

### servidor_java_rmi
- Objeto remoto `ControladorPreferenciasUsuarios` (puerto `2020`)
- Soporta registro de administradores via `registrarReferenciaAdministrador()`
- Después de calcular preferencias, notifica a todos los administradores registrados
- Consulta canciones/reproducciones desde json-server

### cliente_java_rmi
- Menú en consola para consultar preferencias por ID de usuario
- Muestra géneros y artistas con conteo de reproducciones

### administrador_java_rmi
- Se registra como callback en el servidor al iniciar
- Recibe notificaciones automáticas cuando cualquier cliente consulta preferencias
- Muestra las preferencias calculadas en consola

## Tecnologías
- Java 17, Maven, RMI, OpenFeign, Jackson, Lombok
- Patrón **Callback** vía objetos remotos RMI

## Ejecución
1. Iniciar json-server: `json-server --watch db.json --port 3000`
2. Iniciar servidor RMI
3. Iniciar administrador (recibirá notificaciones)
4. Iniciar cliente(s) para consultar preferencias
