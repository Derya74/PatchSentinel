// Package render formats an analyze.Report into human- and machine-friendly
// representations: a colourless plain-text summary and a GitHub-flavoured
// Markdown check summary. JSON rendering lives with the CLI (encoding/json).
package render

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Herricka/PatchSentinel/internal/analyze"
)

// Text renders a compact, deterministic plain-text report suitable for CI logs
// and terminals without ANSI support.
func Text(r analyze.Report) string {
	var b strings.Builder
	fmt.Fprintf(&b, "PatchSentinel report (schema v%d, policy v%d)\n", r.SchemaVersion, r.PolicyVersion)
	fmt.Fprintf(&b, "Verdict: %s\n", strings.ToUpper(string(r.Verdict)))
	b.WriteString(strings.Repeat("-", 60) + "\n")

	s := r.Summary
	fmt.Fprintf(&b, "Files changed      : %d\n", s.FilesChanged)
	fmt.Fprintf(&b, "Lines +/-          : +%d / -%d\n", s.LinesAdded, s.LinesRemoved)
	fmt.Fprintf(&b, "Production added   : %d\n", s.ProductionAdded)
	fmt.Fprintf(&b, "Test added         : %d\n", s.TestAdded)
	fmt.Fprintf(&b, "Test-to-code ratio : %.2f\n", s.TestToCodeRatio)
	fmt.Fprintf(&b, "Risk score         : %d\n", s.RiskScore)
