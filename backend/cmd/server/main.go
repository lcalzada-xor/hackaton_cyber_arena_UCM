package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/config"
	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/models"
	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/pkg/exploitdb"
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

	// Initialize ExploitDB Client
	// Note: In Docker, we might need to ensure the binary is available or adjust this.
	// For now, we follow the CLI pattern.
	edbClient, err := exploitdb.NewClient()
	if err != nil {
		log.Printf("Warning: Failed to initialize exploitdb: %v", err)
	} else {
		// Initial start
		if err := edbClient.StartServer(); err != nil {
			log.Printf("Warning: Failed to start exploitdb server: %v", err)
			edbClient = nil
		} else {
			// Schedule periodic updates
			go func() {
				// Ticker for 12 hours
				ticker := time.NewTicker(12 * time.Hour)
				defer ticker.Stop()

				for range ticker.C {
					log.Println("Starting periodic exploitdb fetch...")
					if err := edbClient.Fetch(); err != nil {
						log.Printf("Error fetching exploitdb updates: %v", err)
						continue
					}
					log.Println("Exploitdb fetch successful. Restarting server to apply changes...")

					// Restart server to load new data
					edbClient.StopServer()
					// Give it a moment to fully stop
					time.Sleep(2 * time.Second)

					if err := edbClient.StartServer(); err != nil {
						log.Printf("Error restarting exploitdb server after fetch: %v", err)
						// If restart fails, we might want to retry or just log it.
						// For now, we just log it. The client might be in a bad state.
					} else {
						log.Println("Exploitdb server restarted successfully.")
					}
				}
			}()

			// Ensure cleanup on shutdown (though main usually kills everything)
			defer edbClient.StopServer()
		}
	}

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

		// Enrich with Exploits
		if edbClient != nil {
			for i := range result.Vulnerabilities {
				cveID := result.Vulnerabilities[i].CVE.ID
				exploits, err := edbClient.Search(cveID)
				if err == nil && len(exploits) > 0 {
					var modelExploits []models.Exploit
					for _, e := range exploits {
						modelExploits = append(modelExploits, models.Exploit{
							ID:          e.ID,
							Name:        e.Name,
							Type:        e.Type,
							URL:         e.URL,
							Description: e.Description,
							Date:        e.Date,
							Author:      e.Author,
						})
					}
					result.Vulnerabilities[i].Exploits = modelExploits
				}
			}
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
