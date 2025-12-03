package models

import "testing"

func TestGetDescription(t *testing.T) {
	cve := Vulnerabilidad{
		Descriptions: []Descripcion{
			{Lang: "en", Value: "English Description"},
			{Lang: "es", Value: "Descripción en Español"},
		},
	}

	tests := []struct {
		name     string
		lang     string
		expected string
	}{
		{"English", "en", "English Description"},
		{"Spanish", "es", "Descripción en Español"},
		{"Unknown", "fr", "English Description"}, // Should return first available if not found
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cve.ObtenerDescripcion(tt.lang)
			if got != tt.expected {
				t.Errorf("ObtenerDescripcion(%q) = %q, want %q", tt.lang, got, tt.expected)
			}
		})
	}
}

func TestGetDescription_Empty(t *testing.T) {
	cve := Vulnerabilidad{}
	if got := cve.ObtenerDescripcion("en"); got != "" {
		t.Errorf("ObtenerDescripcion() = %q, want empty string", got)
	}
}
