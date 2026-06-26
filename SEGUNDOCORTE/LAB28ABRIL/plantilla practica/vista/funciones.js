let clienteChat = null;

function setConectado(conectado) {
    document.getElementById('btnConectar').disabled = conectado;
    document.getElementById('btnDesconectar').disabled = !conectado;
}

function conectar() {
    const socket = new SockJS('http://localhost:5000/ws');
    clienteChat = Stomp.over(socket);
    
    clienteChat.connect({}, onConnected);
}

function onConnected(frame) {
    console.log("Conectado: " + frame);
    // Suscripción al canal de reacciones
    clienteChat.subscribe('/brokerDeReacciones/reaccionesPorCancion', recibirReaccion);
    setConectado(true);
}

function recibirReaccion(reaccion) {
    console.log("Reacción reenviada: " + reaccion.body);
    const contenedor = document.getElementById("mensajes");

    //crear burbuja 
    const burbuja = document.createElement("div");
    burbuja.classList.add("reaccion-burbuja");

    let icono ="";
    switch (reaccion.body) {
        case "like":
            icono = '<i class="fa-solid fa-thumbs-up icon-like"></i>';
            break;
        case "heart":
            icono = '<i class="fa-solid fa-heart icon-heart"></i>';
            break;
        case "sad":
            icono = '<i class="fa-solid fa-face-sad-tear icon-sad"></i>';
            break;
        case "fun":
            icono = '<i class="fa-solid fa-face-laugh-squint icon-fun"></i>';
            break;
    }

    burbuja.innerHTML = icono;
    contenedor.appendChild(burbuja);
    // Eliminar la burbuja después de 3 segundos
    setTimeout(() => {
        burbuja.remove();
    }, 3000);
    
}

function enviarReaccionServidor(event) {
    const icon = event.target;
    handleReactionClick(icon);

    const reaccion = event.target.dataset.reaction;
    if (clienteChat && clienteChat.connected) {
        // Envía la reacción al servidor
        clienteChat.send("/app/enviarReaccion", {}, reaccion);
    } else {
        alert("No estás conectado.");
    }
}
function desconectar() {
    if (clienteChat !== null) {
        clienteChat.disconnect(() => {
            setConectado(false);
            console.log("Desconectado");
        });
    }
}