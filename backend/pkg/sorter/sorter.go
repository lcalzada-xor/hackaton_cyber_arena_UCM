package sorter

import (
	"sort"

	"github.com/lcalzada-xor/hackaton_cyber_arena_UCM/internal/models"
)

// SortVulnerabilities sorts a slice of vulnerabilities based on the given criteria and direction.
// sortBy: "published", "modified", "score"
// direction: "asc", "desc" (default is "desc")
func SortVulnerabilities(vulnerabilities []models.ItemVulnerabilidad, sortBy, direction string) {
	if sortBy == "" {
		return
	}

	sort.Slice(vulnerabilities, func(i, j int) bool {
		a := vulnerabilities[i].CVE
		b := vulnerabilities[j].CVE

		var less bool
		switch sortBy {
		case "published":
			less = a.Published < b.Published
		case "modified":
			less = a.LastModified < b.LastModified
		case "score":
			// Helper to get max score
			obtenerPuntuacionMax := func(c models.Vulnerabilidad) float64 {
				if len(c.Metrics.CvssMetricV31) > 0 {
					return c.Metrics.CvssMetricV31[0].CvssData.BaseScore
				}
				if len(c.Metrics.CvssMetricV30) > 0 {
					return c.Metrics.CvssMetricV30[0].CvssData.BaseScore
				}
				if len(c.Metrics.CvssMetricV2) > 0 {
					return c.Metrics.CvssMetricV2[0].CvssData.BaseScore
				}
				return 0.0
			}
			less = obtenerPuntuacionMax(a) < obtenerPuntuacionMax(b)
		default:
			return false
		}

		if direction == "desc" {
			return !less
		}
		return less
	})
}
