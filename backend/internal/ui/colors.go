package ui

import (
	"regexp"
)

// Códigos de Color ANSI para la Paleta Naranja
const (
	ColorReset       = "\033[0m"
	ColorOrange      = "\033[38;5;208m"               // Naranja Brillante
	ColorLightOrange = "\033[38;5;215m"               // Naranja Suave
	ColorDarkOrange  = "\033[38;5;166m"               // Naranja Más Oscuro
	ColorHighlight   = "\033[48;5;208m\033[38;5;232m" // Negro sobre fondo Naranja
	ColorGray        = "\033[38;5;240m"
	ColorRed         = "\033[31m"
)

// ResaltarPalabraClave resalta la palabra clave en el texto usando códigos de escape ANSI.
func ResaltarPalabraClave(texto, palabraClave string) string {
	if palabraClave == "" {
		return texto
	}
	// Case-insensitive replacement
	re := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(palabraClave))
	return re.ReplaceAllStringFunc(texto, func(match string) string {
		return ColorHighlight + match + ColorReset
	})
}
