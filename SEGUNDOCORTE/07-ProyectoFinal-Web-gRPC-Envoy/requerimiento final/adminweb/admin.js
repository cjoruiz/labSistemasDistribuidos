function cambiarFormularioMetadatos() {
  var tipo = document.getElementById('admin-tipo').value;
  var div = document.getElementById('metadatos-especificos');
  var html = '';
  
  document.getElementById('admin-titulo').value = '';
  
  switch(tipo) {
    case '1':
      html = '<div class="form-group"><label>Artista Principal:</label><input type="text" id="admin-artista" placeholder="Artista o banda"></div>'
        + '<div class="form-group"><label>Album:</label><input type="text" id="admin-album" placeholder="Nombre del album"></div>'
        + '<div class="form-group"><label>Genero Musical:</label><input type="text" id="admin-genero" placeholder="Pop, Rock, Jazz..."></div>';
      break;
    case '2':
      html = '<div class="form-group"><label>Nombre del Podcast:</label><input type="text" id="admin-nombre-podcast" placeholder="Nombre del programa"></div>'
        + '<div class="form-group"><label>Titulo del Episodio:</label><input type="text" id="admin-titulo-episodio" placeholder="Titulo del episodio"></div>'
        + '<div class="form-group"><label>Anfitrion (Host):</label><input type="text" id="admin-host" placeholder="Nombre del host"></div>';
      break;
    case '3':
      html = '<div class="form-group"><label>Titulo del Libro:</label><input type="text" id="admin-titulo-libro" placeholder="Titulo de la obra"></div>'
        + '<div class="form-group"><label>Autor:</label><input type="text" id="admin-autor-libro" placeholder="Autor del texto"></div>'
        + '<div class="form-group"><label>Narrador:</label><input type="text" id="admin-narrador" placeholder="Narrador del audiolibro"></div>';
      break;
    case '4':
      html = '<div class="form-group"><label>Tipo de Sonido:</label><input type="text" id="admin-tipo-sonido" placeholder="Blanco, Marron, Rosa"></div>'
        + '<div class="form-group"><label>Fuente del Audio:</label><input type="text" id="admin-fuente" placeholder="Lluvia, Ventilador, Bosque..."></div>'
        + '<div class="form-group"><label>Uso Sugerido:</label><input type="text" id="admin-uso" placeholder="Dormir, Concentracion, Meditacion"></div>';
      break;
  }
  div.innerHTML = html;
}

function registrarAudio() {
  var tipo = document.getElementById('admin-tipo').value;
  var titulo = document.getElementById('admin-titulo').value.trim();
  var fileInput = document.getElementById('admin-audio-file');
  var file = fileInput.files[0];

  if (!titulo) { alert('Ingresa un titulo'); return; }
  if (!file) { alert('Selecciona un archivo de audio'); return; }

  var tipoNombre = '';
  var metadata = { titulo: titulo, tipoId: parseInt(tipo) };

  switch(tipo) {
    case '1':
      tipoNombre = 'Musica';
      metadata.autor = document.getElementById('admin-artista').value || '';
      metadata.album = document.getElementById('admin-album').value || '';
      metadata.genero = document.getElementById('admin-genero').value || '';
      break;
    case '2':
      tipoNombre = 'Podcast';
      metadata.nombrePodcast = document.getElementById('admin-nombre-podcast').value || '';
      metadata.autor = document.getElementById('admin-titulo-episodio').value || '';
      metadata.genero = document.getElementById('admin-host').value || '';
      break;
    case '3':
      tipoNombre = 'Audiolibro';
      metadata.album = document.getElementById('admin-titulo-libro').value || '';
      metadata.autor = document.getElementById('admin-autor-libro').value || '';
      metadata.narrador = document.getElementById('admin-narrador').value || '';
      break;
    case '4':
      tipoNombre = 'Ruido Blanco';
      metadata.tipoSonido = document.getElementById('admin-tipo-sonido').value || '';
      metadata.fuenteAudio = document.getElementById('admin-fuente').value || '';
      metadata.usoSugerido = document.getElementById('admin-uso').value || '';
      break;
  }
  metadata.tipoNombre = tipoNombre;

  var formData = new FormData();
  formData.append('audio', file);
  formData.append('metadata', JSON.stringify(metadata));

  var msgDiv = document.getElementById('admin-mensaje');
  msgDiv.innerHTML = '<i class="fa-solid fa-spinner fa-spin"></i> Registrando audio...';
  msgDiv.className = 'admin-msg info';

  fetch('http://localhost:8090/registrar', {
    method: 'POST',
    body: formData
  })
  .then(function(res) { return res.json(); })
  .then(function(data) {
    if (data.success) {
      msgDiv.innerHTML = '<i class="fa-solid fa-check-circle"></i> Audio registrado exitosamente!';
      msgDiv.className = 'admin-msg success';
      writeLog('Audio registrado: ' + titulo, 'success');
    } else {
      msgDiv.innerHTML = '<i class="fa-solid fa-times-circle"></i> ' + (data.message || 'Error');
      msgDiv.className = 'admin-msg error';
    }
  })
  .catch(function(err) {
    msgDiv.innerHTML = '<i class="fa-solid fa-times-circle"></i> Error de conexion: ' + err.message;
    msgDiv.className = 'admin-msg error';
  });
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
