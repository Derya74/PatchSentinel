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
	fmt.Fprintf(&b, "Repetition ratio   : %.2f\n", s.RepetitionRatio)
	fmt.Fprintf(&b, "Sensitive files    : %d\n", s.SensitiveFiles)
	fmt.Fprintf(&b, "Est. review time   : %.0f min\n", s.ReviewMinutes)
	fmt.Fprintf(&b, "Content hash       : %s\n", s.ContentHash)

	if len(r.Findings) > 0 {
		b.WriteString("\nFindings:\n")
		for _, f := range r.Findings {
			fmt.Fprintf(&b, "  [%s] %s\n    %s\n", strings.ToUpper(string(f.Severity)), f.Title, f.Detail)
		}
	}

	if len(r.Risks) > 0 {
		b.WriteString("\nRisk signals:\n")
		for _, h := range r.Risks {
			fmt.Fprintf(&b, "  (w%d x%d) %s [%s] %s\n    %s\n",
				h.Weight, h.Count, h.PatternID, h.Path, h.Message, h.Snippet)
		}
	}

	if len(r.Files) > 0 {
		b.WriteString("\nFiles:\n")
		for _, f := range r.Files {
			tags := fileTags(f)
			fmt.Fprintf(&b, "  %-9s +%-4d -%-4d %s%s\n", f.Status, f.Added, f.Removed, f.Path, tags)
		}
	}
	return b.String()
}

func fileTags(f analyze.FileStat) string {
	var tags []string
	if f.IsTest {
		tags = append(tags, "test")
	}
	if f.IsSensitive {
		tags = append(tags, "sensitive")
	}
	if f.IsBinary {
		tags = append(tags, "binary")
