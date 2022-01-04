package policy

import (
	"strings"
	"testing"
)

func TestDefaultIsNormalized(t *testing.T) {
	p := Default()
	if p.Version != 1 {
		t.Errorf("version = %d", p.Version)
	}
	if len(p.RiskPatterns) == 0 || len(p.SensitivePaths) == 0 {
		t.Fatalf("default policy should be populated")
	}
}

func TestLoadOverlaysDefaults(t *testing.T) {
	src := `{"thresholds":{"maxRiskScore":5}}`
	p, err := Load(strings.NewReader(src))
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if p.Thresholds.MaxRiskScore != 5 {
		t.Errorf("maxRiskScore = %d, want 5", p.Thresholds.MaxRiskScore)
	}
	// Defaults preserved for fields not present in the file.
	if len(p.RiskPatterns) == 0 {
		t.Errorf("risk patterns should be inherited from defaults")
