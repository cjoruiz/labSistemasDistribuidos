const { peticionDTO } = require('./stubs_generados/servicios_pb.js');
const { AudioServicePromiseClient } = require('./stubs_generados/servicios_grpc_web_pb.js');

const ENVOY_HOST = 'http://localhost:8080';
const client = new AudioServicePromiseClient(ENVOY_HOST);

let audioChunks = [];
const audioPlayer = document.getElementById('audio-player');
let currentStream = null;

if (audioPlayer) {
    audioPlayer.addEventListener('ended', () => {
        console.log('Reproduccion finalizada.');
    });
}

function iniciar_streaming_cancion(titulo, formato) {
    console.log("Iniciando llamada gRPC-Web a traves de Envoy...");

    if (currentStream) {
        currentStream.cancel();
        currentStream = null;
    }

    audioChunks = [];

    if (audioPlayer) {
        audioPlayer.pause();
        audioPlayer.removeAttribute('src');
    }

    const request = new peticionDTO();
    request.setTitulo(titulo);
    request.setFormato(formato);

    const stream = client.enviarCancionMedianteStream(request, {});
    currentStream = stream;

    let totalBytes = 0;
    let fragmentCount = 0;

    stream.on('data', (response) => {
        const data = response.getData_asU8();
        totalBytes += data.length;
        fragmentCount++;
        console.log("Fragmento #" + fragmentCount + ": " + data.length + " bytes (total: " + totalBytes + " bytes)");
        audioChunks.push(data);

        if (fragmentCount === 1) {
            if (audioPlayer) {
                audioPlayer.style.display = '';
            }
        }
    });

    stream.on('end', () => {
        console.log('Transmision finalizada. Total: ' + totalBytes + ' bytes en ' + fragmentCount + ' fragmentos.');
        const audioBlob = new Blob(audioChunks, { type: 'audio/mpeg' });
        const audioUrl = URL.createObjectURL(audioBlob);
        if (audioPlayer) {
            audioPlayer.src = audioUrl;
            audioPlayer.play()
                .then(() => {
                    console.log('Reproduccion iniciada. Duracion: ' + audioPlayer.duration + 's');
                })
                .catch(error => {
                    console.error('Error al reproducir:', error);
                });
        }
        audioChunks = [];
        currentStream = null;
    });

    stream.on('error', (err) => {
        console.error('ERROR en gRPC-Web:', err);
        if (typeof window !== 'undefined' && window.__writeAudioLog) {
            window.__writeAudioLog('Error en streaming: ' + err, 'error');
        }
        currentStream = null;
    });
}

function detener_streaming() {
    if (currentStream) {
        currentStream.cancel();
        currentStream = null;
    }
    if (audioPlayer) {
        audioPlayer.pause();
        audioPlayer.removeAttribute('src');
    }
    audioChunks = [];
}

if (typeof window !== 'undefined') {
    window.iniciar_streaming_cancion = iniciar_streaming_cancion;
    window.detener_streaming = detener_streaming;
    window.iniciarStreamGRPCImpl = iniciar_streaming_cancion;
    console.log('Exported funciones gRPC-Web a window');
}
