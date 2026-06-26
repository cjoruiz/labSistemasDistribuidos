// Package utilidades provee funciones auxiliares del cliente.
package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lectorEntrada = bufio.NewReader(os.Stdin)

// LeerOpcion lee una línea de la entrada estándar y la convierte a entero.
// Retorna -1 si la entrada no es un número válido.
func LeerOpcion() int {
	entrada, _ := lectorEntrada.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
	opcion, err := strconv.Atoi(entrada)
	if err != nil {
		return -1
	}
	return opcion
}

// ImprimirEncabezado imprime el título de la aplicación con un borde simple.
func ImprimirEncabezado(titulo string) {
	fmt.Println()
	fmt.Println("  ╔══════════════════════════════════════╗")
	fmt.Printf("  ║  %-36s║\n", titulo)
	fmt.Println("  ╚══════════════════════════════════════╝")
}

// ImprimirSeparador imprime una línea separadora.
func ImprimirSeparador() {
	fmt.Println("  ──────────────────────────────────────")
}
