package sorter

import (
	"testing"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/models"
)

func TestSortVulnerabilities(t *testing.T) {
	vulns := []models.ItemVulnerabilidad{
		{
			CVE: models.Vulnerabilidad{
				ID:           "CVE-2021-0001",
				Published:    "2021-01-01",
				LastModified: "2021-02-01",
				Metrics: models.Metricas{
					CvssMetricV31: []models.MetricaCvssV3{
						{CvssData: models.DatosCvss{BaseScore: 5.0}},
					},
				},
			},
		},
		{
			CVE: models.Vulnerabilidad{
				ID:           "CVE-2021-0002",
				Published:    "2021-01-02",
				LastModified: "2021-02-02",
				Metrics: models.Metricas{
					CvssMetricV31: []models.MetricaCvssV3{
						{CvssData: models.DatosCvss{BaseScore: 7.0}},
					},
				},
			},
		},
		{
			CVE: models.Vulnerabilidad{
				ID:           "CVE-2021-0003",
				Published:    "2021-01-03",
				LastModified: "2021-02-03",
				Metrics: models.Metricas{
					CvssMetricV31: []models.MetricaCvssV3{
						{CvssData: models.DatosCvss{BaseScore: 3.0}},
					},
				},
			},
		},
	}

	// Test Sort by Published Ascending
	SortVulnerabilities(vulns, "published", "asc")
	if vulns[0].CVE.ID != "CVE-2021-0001" {
		t.Errorf("Expected CVE-2021-0001, got %s", vulns[0].CVE.ID)
	}

	// Test Sort by Published Descending
	SortVulnerabilities(vulns, "published", "desc")
	if vulns[0].CVE.ID != "CVE-2021-0003" {
		t.Errorf("Expected CVE-2021-0003, got %s", vulns[0].CVE.ID)
	}

	// Test Sort by Score Ascending
	SortVulnerabilities(vulns, "score", "asc")
	if vulns[0].CVE.ID != "CVE-2021-0003" { // Score 3.0
		t.Errorf("Expected CVE-2021-0003, got %s", vulns[0].CVE.ID)
	}

	// Test Sort by Score Descending
	SortVulnerabilities(vulns, "score", "desc")
	if vulns[0].CVE.ID != "CVE-2021-0002" { // Score 7.0
		t.Errorf("Expected CVE-2021-0002, got %s", vulns[0].CVE.ID)
	}
}
