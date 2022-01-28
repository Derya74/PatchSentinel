package render

import (
	"strings"
	"testing"

	"github.com/Herricka/PatchSentinel/internal/analyze"
)

func sample() analyze.Report {
	return analyze.Report{
		SchemaVersion: 1,
		Tool:          "patchsentinel",
		PolicyVersion: 1,
		Summary: analyze.Summary{
			FilesChanged:    2,
			LinesAdded:      12,
			LinesRemoved:    3,
			ProductionAdded: 10,
			TestAdded:       2,
			TestToCodeRatio: 0.2,
			RiskScore:       11,
			RepetitionRatio: 0.1,
			SensitiveFiles:  1,
			ReviewMinutes:   8,
			ContentHash:     "deadbeef",
		},
		Files: []analyze.FileStat{
			{Path: "a.go", Added: 10, Removed: 3, Status: "modified"},
			{Path: ".github/workflows/ci.yml", Added: 2, IsSensitive: true, Status: "modified"},
		},
		Risks: []analyze.RiskHit{
			{PatternID: "debug-print", Message: "Debug print | left", Weight: 1, Count: 1, Path: "a.go", Snippet: "fmt.Println(x)"},
		},
		Findings: []analyze.Finding{
			{ID: "sensitive-path", Severity: analyze.SeverityError, Title: "Sensitive paths modified", Detail: "1 file"},
		},
		Verdict:  analyze.Fail,
		ExitCode: 2,
	}
}

func TestText(t *testing.T) {
	out := Text(sample())
	for _, want := range []string{"Verdict: FAIL", "Risk score         : 11", "Findings:", "Sensitive paths modified", ".github/workflows/ci.yml"} {
		if !strings.Contains(out, want) {
			t.Errorf("text output missing %q\n---\n%s", want, out)
		}
