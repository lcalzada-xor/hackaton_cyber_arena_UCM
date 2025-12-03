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
	// Load Config
	cfg, err := config.CargarConfiguracion()
	if err != nil {
		log.Printf("Warning: Failed to load config: %v", err)
	}

	apiKey := cfg.APIKey
	if envKey := os.Getenv("NVD_API_KEY"); envKey != "" {
		apiKey = envKey
	}

	client := nvd.NuevoCliente(apiKey)

	http.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		query := r.URL.Query()

		limit, _ := strconv.Atoi(query.Get("limit"))
		if limit == 0 {
			limit = 10
		}

		params := nvd.ParametrosBusqueda{
			KeywordSearch:    query.Get("keyword"),
			ResultsPerPage:   limit,
			CvssV3Severity:   query.Get("severity"),
			PubStartDate:     formatDate(query.Get("startDate"), true),
			PubEndDate:       formatDate(query.Get("endDate"), false),
			CpeName:          query.Get("cpe"),
			CweId:            query.Get("cwe"),
			CvssV2Severity:   query.Get("cvssV2Severity"),
			LastModStartDate: formatDate(query.Get("modStartDate"), true),
			LastModEndDate:   formatDate(query.Get("modEndDate"), false),
			SourceIdentifier: query.Get("source"),
		}

		results, err := client.BuscarCVEs(params)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error searching CVEs: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(results); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	})

	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func formatDate(dateStr string, isStart bool) string {
	if dateStr == "" {
		return ""
	}
	// Assuming input is YYYY-MM-DD
	if isStart {
		return dateStr + "T00:00:00.000"
	}
	return dateStr + "T23:59:59.999"
}
