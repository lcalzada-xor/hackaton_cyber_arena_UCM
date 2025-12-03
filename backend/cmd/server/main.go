package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/config"
	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/pkg/nvd"
)

func main() {
	// Load Configuration
	cfg, err := config.CargarConfiguracion()
	if err != nil {
		log.Printf("Warning: Failed to load configuration: %v", err)
	}

	// Initialize NVD Client
	client := nvd.NuevoCliente(cfg.APIKey)

	http.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS - Restrict to Frontend Origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		query := r.URL.Query()

		// Map query params to NVD Search Parameters
		params := nvd.ParametrosBusqueda{
			KeywordSearch:    query.Get("keyword"),
			CvssV3Severity:   query.Get("severity"),
			PubStartDate:     query.Get("startDate"),
			PubEndDate:       query.Get("endDate"),
			CpeName:          query.Get("cpe"),
			CweId:            query.Get("cwe"),
			CvssV2Severity:   query.Get("cvssV2Severity"),
			LastModStartDate: query.Get("modStartDate"),
			LastModEndDate:   query.Get("modEndDate"),
			SourceIdentifier: query.Get("source"),
		}

		// Handle Limit / ResultsPerPage
		if limitStr := query.Get("limit"); limitStr != "" {
			if limit, err := strconv.Atoi(limitStr); err == nil {
				params.ResultsPerPage = limit
			}
		} else {
			params.ResultsPerPage = cfg.DefaultLimit
		}

		// Execute Search
		result, err := client.BuscarCVEs(params)
		if err != nil {
			log.Printf("Error searching CVEs: %v", err)
			http.Error(w, "Error searching CVEs", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	})

	port := "8081"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Helper functions removed as they are no longer needed
