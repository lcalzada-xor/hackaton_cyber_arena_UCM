package nvd

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/models"
)

// Cache define la interfaz para cachear respuestas del NVD.
type Cache interface {
	Obtener(key string) (*models.RespuestaNVD, bool)
	Guardar(key string, value *models.RespuestaNVD) error
}

// CacheArchivo implementa una caché simple basada en archivos.
type CacheArchivo struct {
	Dir string
	TTL time.Duration
}

// NuevaCacheArchivo crea una nueva CacheArchivo.
func NuevaCacheArchivo(ttl time.Duration) (*CacheArchivo, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cacheDir := filepath.Join(home, ".cve-search", "cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return nil, err
	}
	return &CacheArchivo{Dir: cacheDir, TTL: ttl}, nil
}

func (c *CacheArchivo) getPath(key string) string {
	hash := sha256.Sum256([]byte(key))
	filename := hex.EncodeToString(hash[:]) + ".json"
	return filepath.Join(c.Dir, filename)
}

// Obtener recupera un valor de la caché.
func (c *CacheArchivo) Obtener(key string) (*models.RespuestaNVD, bool) {
	path := c.getPath(key)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, false
	}
	if time.Since(info.ModTime()) > c.TTL {
		_ = os.Remove(path) // Expirado
		return nil, false
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var resp models.RespuestaNVD
	if err := json.NewDecoder(file).Decode(&resp); err != nil {
		return nil, false
	}
	return &resp, true
}

// Guardar guarda un valor en la caché.
func (c *CacheArchivo) Guardar(key string, value *models.RespuestaNVD) error {
	path := c.getPath(key)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(value)
}
