package vistas

var usuariosRegistrados = map[string]string{
	"admin":  "admin123",
	"juan":   "juan123",
	"maria":  "maria123",
	"carlos": "carlos123",
}

func validarCredenciales(nickname, contrasena string) bool {
	if contrasenaGuardada, existe := usuariosRegistrados[nickname]; existe {
		return contrasenaGuardada == contrasena
	}
	return false
}