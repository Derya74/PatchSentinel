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

