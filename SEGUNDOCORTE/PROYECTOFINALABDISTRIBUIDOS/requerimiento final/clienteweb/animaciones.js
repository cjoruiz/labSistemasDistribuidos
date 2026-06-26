function handleReactionClick(icon) {
    icon.classList.add("gold-glow");
    setTimeout(function() {
        icon.classList.remove("gold-glow");
    }, 300);
}

function animarEntradaNotificacion(elemento) {
    elemento.style.opacity = '0';
    elemento.style.transform = 'translateX(-20px)';
    setTimeout(function() {
        elemento.style.transition = 'all 0.3s ease-out';
        elemento.style.opacity = '1';
        elemento.style.transform = 'translateX(0)';
    }, 50);
}

document.addEventListener('DOMContentLoaded', function() {
    var icons = document.querySelectorAll('.reaction');
    icons.forEach(function(icon) {
        icon.addEventListener('click', function() {
            icon.classList.add('active');
            setTimeout(function() { icon.classList.remove('active'); }, 300);
        });
    });
});
