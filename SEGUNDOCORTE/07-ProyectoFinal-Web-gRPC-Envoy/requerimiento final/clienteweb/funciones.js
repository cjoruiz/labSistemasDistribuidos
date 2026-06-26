let clienteChat = null;
let audioSeleccionado = null;
let usuarioActual = '';
let conectadoReacciones = false;
let suscripcionReacciones = null;
let suscripcionPresencia = null;
let reenvioReaccionEnProgreso = false;
let iniciandoReproduccion = false;

function conectarReacciones() {
    usuarioActual = document.getElementById('nickname-input').value.trim() || 'Anonimo';

    const socket = new SockJS('http://localhost:5000/ws');
    clienteChat = Stomp.over(socket);

    clienteChat.connect({}, function(frame) {
        console.log("Conectado a servidor de reacciones: " + frame);
        conectadoReacciones = true;

        document.getElementById('btnConectar').disabled = true;
        document.getElementById('btnDesconectar').disabled = false;
        var status = document.getElementById('conexion-status');
        status.className = 'status-conectado';
        status.innerHTML = '<i class="fa-solid fa-circle"></i> Conectado como ' + usuarioActual;

        writeLog('Conectado como: ' + usuarioActual, 'success');
    }, function(error) {
        console.error('Error de conexion WebSocket:', error);
        writeLog('Error de conexion: ' + error, 'error');
    });
}

function desconectarReacciones() {
    if (suscripcionReacciones) { suscripcionReacciones.unsubscribe(); suscripcionReacciones = null; }
    if (suscripcionPresencia) { suscripcionPresencia.unsubscribe(); suscripcionPresencia = null; }

    if (clienteChat !== null) {
        if (audioSeleccionado) {
            notificarEstadoReproduccion('stop');
        }
        clienteChat.disconnect(function() {
            conectadoReacciones = false;
            document.getElementById('btnConectar').disabled = false;
            document.getElementById('btnDesconectar').disabled = true;
            var status = document.getElementById('conexion-status');
            status.className = 'status-desconectado';
            status.innerHTML = '<i class="fa-solid fa-circle"></i> Desconectado';
            writeLog('Desconectado de servidor de reacciones', 'error');
        });
    }
}

function suscribirReaccionCancion(cancionId) {
    if (!clienteChat || !clienteChat.connected) return;

    if (suscripcionReacciones) { suscripcionReacciones.unsubscribe(); suscripcionReacciones = null; }
    if (suscripcionPresencia) { suscripcionPresencia.unsubscribe(); suscripcionPresencia = null; }

    suscripcionReacciones = clienteChat.subscribe(
        '/brokerDeReacciones/cancion/' + cancionId,
        function(mensaje) { mostrarReaccionRecibida(mensaje.body); }
    );

    suscripcionPresencia = clienteChat.subscribe(
        '/brokerDeReacciones/cancion/' + cancionId + '/presencia',
        function(mensaje) { mostrarNotificacionUsuario(mensaje.body); }
    );

    writeLog('Suscrito a reacciones de: ' + cancionId, 'success');
}

function enviarReaccion(event) {
    if (!clienteChat || !clienteChat.connected) {
        alert('Primero conectate al servidor de reacciones.');
        return;
    }
    if (!audioSeleccionado) {
        alert('Selecciona y reproduce un audio primero.');
        return;
    }
    if (reenvioReaccionEnProgreso) return;
    reenvioReaccionEnProgreso = true;
    setTimeout(function() { reenvioReaccionEnProgreso = false; }, 500);

    var icon = event.currentTarget;
    handleReactionClick(icon);
    var reaccion = icon.getAttribute('data-reaccion');
    var mensaje = JSON.stringify({
        nickname: usuarioActual,
        cancionId: audioSeleccionado.archivo,
        reaccion: reaccion
    });
    clienteChat.send('/app/enviarReaccion', {}, mensaje);
}

function mostrarReaccionRecibida(mensaje) {
    try {
        var datos = JSON.parse(mensaje);
        var contenedor = document.getElementById("mensajes-reacciones");
        var placeholder = contenedor.querySelector('.placeholder');
        if (placeholder) placeholder.remove();

        var burbuja = document.createElement("div");
        burbuja.classList.add("reaccion-burbuja");

        var iconoHtml = '';
        switch (datos.reaccion) {
            case "like": iconoHtml = '<i class="fa-solid fa-thumbs-up icon-like"></i>'; break;
            case "heart": iconoHtml = '<i class="fa-solid fa-heart icon-heart"></i>'; break;
            case "sad": iconoHtml = '<i class="fa-solid fa-face-sad-tear icon-sad"></i>'; break;
            case "fun": iconoHtml = '<i class="fa-solid fa-face-laugh-squint icon-fun"></i>'; break;
            default: iconoHtml = '<i class="fa-solid fa-star"></i>';
        }

        var nombreSpan = document.createElement("span");
        nombreSpan.className = "reaccion-usuario";
        nombreSpan.textContent = datos.nickname || 'Alguien';

        burbuja.innerHTML = iconoHtml;
        burbuja.appendChild(nombreSpan);
        contenedor.appendChild(burbuja);

        setTimeout(function() { burbuja.remove(); }, 3000);
    } catch(e) {
        var contenedor2 = document.getElementById("mensajes-reacciones");
        var burbuja2 = document.createElement("div");
        burbuja2.classList.add("reaccion-burbuja");
        burbuja2.textContent = mensaje;
        contenedor2.appendChild(burbuja2);
        setTimeout(function() { burbuja2.remove(); }, 3000);
    }
}

function mostrarNotificacionUsuario(mensaje) {
    try {
        var datos = JSON.parse(mensaje);
        var contenedor = document.getElementById("notificaciones-usuarios");
        var placeholder = contenedor.querySelector('.placeholder');
        if (placeholder) placeholder.remove();

        var notif = document.createElement("div");
        notif.className = "notificacion-usuario " + (datos.accion === 'play' ? 'notif-play' : 'notif-pause');

        var icono = datos.accion === 'play'
            ? '<i class="fa-solid fa-play" style="color:#2ecc71"></i>'
            : '<i class="fa-solid fa-pause" style="color:#e74c3c"></i>';

        var texto = datos.accion === 'play'
            ? '<strong>' + datos.nickname + '</strong> comenzo a reproducir <em>"' + datos.cancionId + '"</em>'
            : '<strong>' + datos.nickname + '</strong> pauso la reproduccion de <em>"' + datos.cancionId + '"</em>';

        notif.innerHTML = icono + ' ' + texto;
        contenedor.appendChild(notif);

        animarEntradaNotificacion(notif);

        if (contenedor.children.length > 20) {
            contenedor.removeChild(contenedor.firstChild);
        }

        setTimeout(function() {
            notif.style.opacity = '0';
            notif.style.transform = 'translateX(100px)';
            setTimeout(function() { if (notif.parentNode) notif.remove(); }, 500);
        }, 6000);
    } catch(e) {
        console.log("Notificacion recibida:", mensaje);
    }
}

function notificarEstadoReproduccion(estado) {
    if (!clienteChat || !clienteChat.connected || !audioSeleccionado) return;
    if (iniciandoReproduccion && estado === 'pausa') return;

    var accion = estado === 'reproduciendo' ? 'play' : 'stop';
    var mensaje = JSON.stringify({
        nickname: usuarioActual,
        cancionId: audioSeleccionado.archivo,
        accion: accion
    });
    clienteChat.send('/app/notificarEstado', {}, mensaje);
}

function toggleLog() {
    var content = document.getElementById('log-content');
    var icon = document.getElementById('log-toggle-icon');
    if (content.style.display === 'none') {
        content.style.display = 'block';
        icon.textContent = '▼';
    } else {
        content.style.display = 'none';
        icon.textContent = '▲';
    }
}

function writeLog(message, level) {
    var d = document.getElementById('log-content');
    if (!d) return;
    var p = document.createElement('div');
    p.className = level || '';
    var ts = new Date().toLocaleTimeString();
    p.textContent = '[' + ts + '] ' + message;
    d.appendChild(p);
    d.scrollTop = d.scrollHeight;
}

function attachAudioPlayerListeners() {
    var audio = document.getElementById('audio-player');
    if (!audio) return;
    audio.addEventListener('play', function() {
        iniciandoReproduccion = false;
        writeLog('Reproduccion iniciada (play).', 'success');
        notificarEstadoReproduccion('reproduciendo');
    });
    audio.addEventListener('pause', function() {
        writeLog('Reproduccion pausada (pause).', 'error');
        notificarEstadoReproduccion('pausa');
    });
    audio.addEventListener('ended', function() {
        writeLog('Reproduccion finalizada (ended).', 'error');
        notificarEstadoReproduccion('pausa');
    });
}

document.addEventListener('DOMContentLoaded', function() {
    attachAudioPlayerListeners();
    cargarTiposAudio();
});

function cargarTiposAudio() {
    var lista = document.getElementById('tipos-lista');
    lista.innerHTML = '<div class="loading">Cargando...</div>';

    fetch('http://localhost:8090/tipos')
        .then(function(res) { return res.json(); })
        .then(function(tipos) {
            lista.innerHTML = '';
            tipos.forEach(function(t) {
                var btn = document.createElement('button');
                btn.className = 'tipo-btn';
                var icono = '';
                switch(t.id) {
                    case 1: icono = '<i class="fa-solid fa-music"></i>'; break;
                    case 2: icono = '<i class="fa-solid fa-podcast"></i>'; break;
                    case 3: icono = '<i class="fa-solid fa-book"></i>'; break;
                    case 4: icono = '<i class="fa-solid fa-wind"></i>'; break;
                }
                btn.innerHTML = icono + ' ' + t.nombre;
                btn.onclick = function() { cargarAudiosPorTipo(t.id); };
                lista.appendChild(btn);
            });
            writeLog('Tipos de audio cargados: ' + tipos.length, 'success');
        })
        .catch(function(err) {
            lista.innerHTML = '<div class="error-msg">Error al cargar tipos. Asegurate que el servidor de metadatos este en puerto 8090.</div>';
            console.error('Error cargando tipos:', err);
        });
}

function cargarAudiosPorTipo(tipoId) {
    var lista = document.getElementById('audios-lista');
    lista.innerHTML = '<div class="loading">Cargando audios...</div>';

    fetch('http://localhost:8090/audios?tipo=' + tipoId)
        .then(function(res) { return res.json(); })
        .then(function(audios) {
            lista.innerHTML = '';
            if (audios.length === 0) {
                lista.innerHTML = '<div class="placeholder">No hay audios disponibles</div>';
                return;
            }
            audios.forEach(function(a) {
                var btn = document.createElement('button');
                btn.className = 'audio-btn';
                btn.innerHTML = '<i class="fa-solid fa-file-audio"></i> ' + a.titulo;
                btn.onclick = function() { seleccionarAudio(a, tipoId); };
                lista.appendChild(btn);
            });
            writeLog('Audios cargados para tipo ' + tipoId + ': ' + audios.length, 'success');
        })
        .catch(function(err) {
            lista.innerHTML = '<div class="error-msg">Error al cargar audios</div>';
            console.error('Error cargando audios:', err);
        });
}

function seleccionarAudio(audio, tipoId) {
    if (audioSeleccionado && clienteChat && clienteChat.connected) {
        notificarEstadoReproduccion('pausa');
    }

    if (suscripcionReacciones) { suscripcionReacciones.unsubscribe(); suscripcionReacciones = null; }
    if (suscripcionPresencia) { suscripcionPresencia.unsubscribe(); suscripcionPresencia = null; }

    audioSeleccionado = audio;

    fetch('http://localhost:8090/audio/' + audio.id)
        .then(function(res) { return res.json(); })
        .then(function(detalles) {
            mostrarDetalles(detalles);
            document.getElementById('btn-reproducir').disabled = false;
            writeLog('Audio seleccionado: ' + audio.titulo, 'success');
        })
        .catch(function(err) {
            mostrarDetallesDesdeResumen(audio, tipoId);
            document.getElementById('btn-reproducir').disabled = false;
        });
}

function mostrarDetalles(d) {
    var cont = document.getElementById('detalles-contenido');
    var html = '<div class="detalle-titulo">' + d.titulo + '</div>';
    html += '<table class="detalle-tabla">';

    if (d.tipo) html += '<tr><td>Tipo</td><td>' + d.tipo + '</td></tr>';

    if (d.tipo === 'Musica') {
        if (d.autor) html += '<tr><td>Artista</td><td>' + d.autor + '</td></tr>';
        if (d.album) html += '<tr><td>Album</td><td>' + d.album + '</td></tr>';
        if (d.genero) html += '<tr><td>Genero</td><td>' + d.genero + '</td></tr>';
        if (d.selloDiscografico) html += '<tr><td>Sello Discografico</td><td>' + d.selloDiscografico + '</td></tr>';
    } else if (d.tipo === 'Podcast') {
        if (d.nombrePodcast) html += '<tr><td>Podcast</td><td>' + d.nombrePodcast + '</td></tr>';
        if (d.autor) html += '<tr><td>Titulo Episodio</td><td>' + d.autor + '</td></tr>';
        if (d.genero) html += '<tr><td>Anfitrion</td><td>' + d.genero + '</td></tr>';
        if (d.numeroEpisodio) html += '<tr><td>Episodio</td><td>' + d.numeroEpisodio + '</td></tr>';
        if (d.notasShow) html += '<tr><td>Notas</td><td>' + d.notasShow + '</td></tr>';
    } else if (d.tipo === 'Audiolibro') {
        if (d.album) html += '<tr><td>Titulo del Libro</td><td>' + d.album + '</td></tr>';
        if (d.autor) html += '<tr><td>Autor</td><td>' + d.autor + '</td></tr>';
        if (d.narrador) html += '<tr><td>Narrador</td><td>' + d.narrador + '</td></tr>';
        if (d.editorial) html += '<tr><td>Editorial</td><td>' + d.editorial + '</td></tr>';
        if (d.isbn) html += '<tr><td>ISBN</td><td>' + d.isbn + '</td></tr>';
        if (d.capitulo) html += '<tr><td>Capitulo</td><td>' + d.capitulo + '</td></tr>';
    } else if (d.tipo === 'Ruido Blanco') {
        if (d.tipoSonido) html += '<tr><td>Tipo Sonido</td><td>' + d.tipoSonido + '</td></tr>';
        if (d.fuenteAudio) html += '<tr><td>Fuente</td><td>' + d.fuenteAudio + '</td></tr>';
        if (d.usoSugerido) html += '<tr><td>Uso Sugerido</td><td>' + d.usoSugerido + '</td></tr>';
    }

    if (d.duracion) html += '<tr><td>Duracion</td><td>' + d.duracion + 's</td></tr>';

    html += '</table>';
    cont.innerHTML = html;
}

function mostrarDetallesDesdeResumen(audio, tipoId) {
    var cont = document.getElementById('detalles-contenido');
    var nombres = ['', 'Musica', 'Podcast', 'Audiolibro', 'Ruido Blanco'];
    cont.innerHTML = '<div class="detalle-titulo">' + audio.titulo + '</div>'
        + '<table class="detalle-tabla">'
        + '<tr><td>Tipo</td><td>' + (nombres[tipoId] || 'Desconocido') + '</td></tr>'
        + '<tr><td>Archivo</td><td>' + audio.archivo + '.' + audio.formato + '</td></tr>'
        + '</table>';
}

function reproducirAudioSeleccionado() {
    if (!audioSeleccionado) return;

    iniciandoReproduccion = true;
    setTimeout(function() { iniciandoReproduccion = false; }, 5000);
    suscribirReaccionCancion(audioSeleccionado.archivo);

    var btn = document.getElementById('btn-reproducir');
    btn.innerHTML = '<i class="fa-solid fa-spinner fa-spin"></i> Cargando...';
    btn.disabled = true;

    if (typeof window.iniciar_streaming_cancion === 'function') {
        window.iniciar_streaming_cancion(audioSeleccionado.archivo, audioSeleccionado.formato);
    } else {
        alert('Error: funcion de streaming no disponible. Ejecuta npx webpack primero.');
    }

    setTimeout(function() {
        btn.innerHTML = '<i class="fa-solid fa-play"></i> Reproducir';
        btn.disabled = false;
    }, 1000);
}
