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
