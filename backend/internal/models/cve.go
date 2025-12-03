package models

// RespuestaNVD representa la respuesta de nivel superior de la API del NVD.
type RespuestaNVD struct {
	ResultsPerPage  int                  `json:"resultsPerPage"`
	StartIndex      int                  `json:"startIndex"`
	TotalResults    int                  `json:"totalResults"`
	Format          string               `json:"format"`
	Version         string               `json:"version"`
	Timestamp       string               `json:"timestamp"`
	Vulnerabilities []ItemVulnerabilidad `json:"vulnerabilities"`
}

// ItemVulnerabilidad envuelve los detalles del CVE.
type ItemVulnerabilidad struct {
	CVE Vulnerabilidad `json:"cve"`
}

// Vulnerabilidad contiene los datos principales de Vulnerabilidades y Exposiciones Comunes.
type Vulnerabilidad struct {
	ID               string          `json:"id"`
	SourceIdentifier string          `json:"sourceIdentifier"`
	Published        string          `json:"published"` // Manteniendo como string por simplicidad, se puede parsear luego
	LastModified     string          `json:"lastModified"`
	VulnStatus       string          `json:"vulnStatus"`
	Descriptions     []Descripcion   `json:"descriptions"`
	Metrics          Metricas        `json:"metrics"`
	Weaknesses       []Debilidad     `json:"weaknesses"`
	Configurations   []Configuracion `json:"configurations"`
	References       []Referencia    `json:"references"`
}

// Descripcion representa una descripción en un idioma específico.
type Descripcion struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

// Metricas contiene puntuaciones CVSS.
type Metricas struct {
	CvssMetricV2  []MetricaCvssV2 `json:"cvssMetricV2,omitempty"`
	CvssMetricV30 []MetricaCvssV3 `json:"cvssMetricV30,omitempty"`
	CvssMetricV31 []MetricaCvssV3 `json:"cvssMetricV31,omitempty"`
}

// MetricaCvssV2 representa la puntuación CVSS v2.
type MetricaCvssV2 struct {
	Source                  string    `json:"source"`
	Type                    string    `json:"type"`
	CvssData                DatosCvss `json:"cvssData"`
	BaseSeverity            string    `json:"baseSeverity"`
	ExploitabilityScore     float64   `json:"exploitabilityScore"`
	ImpactScore             float64   `json:"impactScore"`
	AcInsufInfo             bool      `json:"acInsufInfo"`
	ObtainAllPrivilege      bool      `json:"obtainAllPrivilege"`
	ObtainUserPrivilege     bool      `json:"obtainUserPrivilege"`
	ObtainOtherPrivilege    bool      `json:"obtainOtherPrivilege"`
	UserInteractionRequired bool      `json:"userInteractionRequired"`
}

// MetricaCvssV3 representa la puntuación CVSS v3.0/v3.1.
type MetricaCvssV3 struct {
	Source              string    `json:"source"`
	Type                string    `json:"type"`
	CvssData            DatosCvss `json:"cvssData"`
	ExploitabilityScore float64   `json:"exploitabilityScore"`
	ImpactScore         float64   `json:"impactScore"`
}

// DatosCvss contiene el vector CVSS crudo y las puntuaciones.
type DatosCvss struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	BaseScore             float64 `json:"baseScore"`
	AccessVector          string  `json:"accessVector,omitempty"`     // V2
	AccessComplexity      string  `json:"accessComplexity,omitempty"` // V2
	Authentication        string  `json:"authentication,omitempty"`   // V2
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
	AttackVector          string  `json:"attackVector,omitempty"`       // V3
	AttackComplexity      string  `json:"attackComplexity,omitempty"`   // V3
	PrivilegesRequired    string  `json:"privilegesRequired,omitempty"` // V3
	UserInteraction       string  `json:"userInteraction,omitempty"`    // V3
	Scope                 string  `json:"scope,omitempty"`              // V3
	BaseSeverity          string  `json:"baseSeverity,omitempty"`       // V3
}

// Debilidad representa un CWE u otro tipo de debilidad.
type Debilidad struct {
	Source      string        `json:"source"`
	Type        string        `json:"type"`
	Description []Descripcion `json:"description"`
}

// Configuracion representa configuraciones de software afectadas (CPEs).
type Configuracion struct {
	Nodes []Nodo `json:"nodes"`
}

// Nodo representa una agrupación lógica de criterios de coincidencia.
type Nodo struct {
	Operator string            `json:"operator"`
	Negate   bool              `json:"negate"`
	CpeMatch []CoincidenciaCPE `json:"cpeMatch"`
}

// CoincidenciaCPE representa una cadena de coincidencia CPE específica.
type CoincidenciaCPE struct {
	Vulnerable      bool   `json:"vulnerable"`
	Criteria        string `json:"criteria"`
	MatchCriteriaID string `json:"matchCriteriaId"`
}

// Referencia representa un enlace a información externa.
type Referencia struct {
	URL    string `json:"url"`
	Source string `json:"source"`
}

// ObtenerDescripcion devuelve la descripción en el idioma solicitado o la primera disponible.
func (c *Vulnerabilidad) ObtenerDescripcion(idioma string) string {
	for _, d := range c.Descriptions {
		if d.Lang == idioma {
			return d.Value
		}
	}
	if len(c.Descriptions) > 0 {
		return c.Descriptions[0].Value
	}
	return ""
}
