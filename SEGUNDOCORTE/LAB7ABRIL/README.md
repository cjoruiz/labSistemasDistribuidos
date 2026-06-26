# LAB7ABRIL - Sistema de Preferencias de Usuarios con RMI

## Descripción
Sistema distribuido que calcula las preferencias musicales de usuarios (géneros y artistas más escuchados) basado en su historial de reproducciones. Implementa una arquitectura cliente-servidor usando **Java RMI**.

## Arquitectura
```
Cliente (RMI) <--> Servidor RMI <--> json-server (API REST mock)
```

## Componentes

### servidor_java_rmi
- Expone un objeto remoto `ControladorPreferenciasUsuarios` en el puerto `2020`
- Consulta canciones y reproducciones desde un `json-server` externo (puerto `3000`)
- Calcula preferencias usando el componente `CalculadorPreferencias`
- Cliente HTTP Feign para comunicación REST con json-server

### cliente_java_rmi
- Se conecta al servidor RMI y presenta un menú en consola
- Opción 1: Consultar preferencias por ID de usuario
- Opción 2: Salir

## Tecnologías
- Java 17, Maven, RMI, OpenFeign, Jackson, Lombok
- json-server (base de datos mock en `db.json`)

## Ejecución
1. Iniciar json-server: `json-server --watch db.json --port 3000`
2. Iniciar servidor RMI: ejecutar `main` de `servidor_java_rmi`
3. Iniciar cliente: ejecutar `main` de `cliente_java_rmi`
