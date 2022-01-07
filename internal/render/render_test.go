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
