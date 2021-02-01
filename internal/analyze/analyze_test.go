package analyze

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Herricka/PatchSentinel/internal/diff"
	"github.com/Herricka/PatchSentinel/internal/policy"
)

func parse(t *testing.T, s string) diff.Patch {
	t.Helper()
	p, err := diff.Parse(strings.NewReader(s))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return p
}
