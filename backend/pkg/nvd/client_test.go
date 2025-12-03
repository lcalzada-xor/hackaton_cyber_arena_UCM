package nvd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/models"
)

func TestSearchCVEs(t *testing.T) {
	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query params
		if r.URL.Query().Get("keywordSearch") != "log4j" {
			t.Errorf("Expected keywordSearch=log4j, got %s", r.URL.Query().Get("keywordSearch"))
		}

		// Mock Response
		resp := models.RespuestaNVD{
			TotalResults: 1,
			Vulnerabilities: []models.ItemVulnerabilidad{
				{
					CVE: models.Vulnerabilidad{
						ID: "CVE-2021-44228",
					},
				},
			},
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
	defer server.Close()

	// Create Client with Mock URL
	client := NuevoCliente("test-api-key")
	client.BaseURL = server.URL

	// Test Search
	params := ParametrosBusqueda{
		KeywordSearch: "log4j",
	}
	results, err := client.BuscarCVEs(params)
	if err != nil {
		t.Fatalf("SearchCVEs failed: %v", err)
	}

	if results.TotalResults != 1 {
		t.Errorf("Expected 1 result, got %d", results.TotalResults)
	}
	if len(results.Vulnerabilities) != 1 {
		t.Fatalf("Expected 1 vulnerability, got %d", len(results.Vulnerabilities))
	}
	if results.Vulnerabilities[0].CVE.ID != "CVE-2021-44228" {
		t.Errorf("Expected CVE-2021-44228, got %s", results.Vulnerabilities[0].CVE.ID)
	}
}

func TestSearchCVEs_Error(t *testing.T) {
	// Mock Server Error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NuevoCliente("test-api-key")
	client.BaseURL = server.URL

	params := ParametrosBusqueda{KeywordSearch: "fail"}
	_, err := client.BuscarCVEs(params)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestSearchCVEs_JSONError(t *testing.T) {
	// Mock Server with Invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("invalid json")); err != nil {
			return
		}
	}))
	defer server.Close()

	client := NuevoCliente("test-api-key")
	client.BaseURL = server.URL

	params := ParametrosBusqueda{KeywordSearch: "jsonfail"}
	_, err := client.BuscarCVEs(params)
	if err == nil {
		t.Error("Expected error on invalid JSON, got nil")
	}
}
