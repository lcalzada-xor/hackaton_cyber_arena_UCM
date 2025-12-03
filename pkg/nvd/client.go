package nvd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/models"
)

const BaseURL = "https://services.nvd.nist.gov/rest/json/cves/2.0"

// Cliente interactúa con la API del NVD.
type Cliente struct {
	httpClient *http.Client
	apiKey     string // Opcional: Clave API NVD para límites de tasa más altos
	BaseURL    string
	Cache      Cache
}

// NuevoCliente crea un nuevo cliente de la API del NVD.
func NuevoCliente(claveAPI string) *Cliente {
	// Inicializar caché con TTL de 1 hora por defecto
	cache, err := NuevaCacheArchivo(1 * time.Hour)
	if err != nil {
		// Si falla la caché, continuar sin ella
		cache = nil
	}

	return &Cliente{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		apiKey:  claveAPI,
		BaseURL: BaseURL,
		Cache:   cache,
	}
}

// ParametrosBusqueda define los parámetros para buscar CVEs.
type ParametrosBusqueda struct {
	KeywordSearch    string
	ResultsPerPage   int
	StartIndex       int
	CvssV3Severity   string
	PubStartDate     string
	PubEndDate       string
	CpeName          string
	CweId            string
	CvssV2Severity   string
	LastModStartDate string
	LastModEndDate   string
	SourceIdentifier string
}

// BuscarCVEs obtiene CVEs basados en los parámetros proporcionados.
func (c *Cliente) BuscarCVEs(parametros ParametrosBusqueda) (*models.RespuestaNVD, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	q := u.Query()
	if parametros.KeywordSearch != "" {
		q.Set("keywordSearch", parametros.KeywordSearch)
	}
	if parametros.ResultsPerPage > 0 {
		q.Set("resultsPerPage", fmt.Sprintf("%d", parametros.ResultsPerPage))
	}
	if parametros.StartIndex >= 0 {
		q.Set("startIndex", fmt.Sprintf("%d", parametros.StartIndex))
	}
	if parametros.CvssV3Severity != "" {
		q.Set("cvssV3Severity", parametros.CvssV3Severity)
	}
	if parametros.PubStartDate != "" {
		q.Set("pubStartDate", parametros.PubStartDate)
	}
	if parametros.PubEndDate != "" {
		q.Set("pubEndDate", parametros.PubEndDate)
	}
	if parametros.CpeName != "" {
		q.Set("cpeName", parametros.CpeName)
	}
	if parametros.CweId != "" {
		q.Set("cweId", parametros.CweId)
	}
	if parametros.CvssV2Severity != "" {
		q.Set("cvssV2Severity", parametros.CvssV2Severity)
	}
	if parametros.LastModStartDate != "" {
		q.Set("lastModStartDate", parametros.LastModStartDate)
	}
	if parametros.LastModEndDate != "" {
		q.Set("lastModEndDate", parametros.LastModEndDate)
	}
	if parametros.SourceIdentifier != "" {
		q.Set("sourceIdentifier", parametros.SourceIdentifier)
	}

	u.RawQuery = q.Encode()
	reqURL := u.String()

	// Check Cache
	if c.Cache != nil {
		if cached, found := c.Cache.Obtener(reqURL); found {
			return cached, nil
		}
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Set("apiKey", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("NVD API returned status: %s", resp.Status)
	}

	var result models.RespuestaNVD
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Save to Cache
	if c.Cache != nil {
		if err := c.Cache.Guardar(reqURL, &result); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to write to cache: %v\n", err)
		}
	}

	return &result, nil
}
