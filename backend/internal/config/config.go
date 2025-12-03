package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Configuracion mantiene la configuración de la aplicación.
type Configuracion struct {
	APIKey       string `json:"api_key"`
	OutputFormat string `json:"output_format"`
	DefaultLimit int    `json:"default_limit"`
}

// CargarConfiguracion carga la configuración desde un archivo y variables de entorno.
// Prioridad: Vars Entorno > Archivo Config > Valores por Defecto
func CargarConfiguracion() (*Configuracion, error) {
	// Valores por Defecto
	cfg := &Configuracion{
		OutputFormat: "human",
		DefaultLimit: 10,
	}

	// Cargar desde archivo
	home, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(home, ".cve-search.json")
		if file, err := os.Open(configPath); err == nil {
			defer file.Close()
			if err := json.NewDecoder(file).Decode(cfg); err != nil {
				fmt.Fprintf(os.Stderr, "Advertencia: Fallo al decodificar archivo de configuración: %v\n", err)
			}
		}
	}

	// Cargar desde Vars Entorno
	if key := os.Getenv("NVD_API_KEY"); key != "" {
		cfg.APIKey = key
	}
	if fmt := os.Getenv("CVE_OUTPUT_FORMAT"); fmt != "" {
		cfg.OutputFormat = fmt
	}

	return cfg, nil
}
