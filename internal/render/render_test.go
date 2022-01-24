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
