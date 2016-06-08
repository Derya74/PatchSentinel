// Package diff parses unified diff files into a structured, deterministic model.
//
// The parser is tolerant of common variations produced by `git diff`,
// `diff -u`, and similar tools. It intentionally uses only the Go standard
// library so that PatchSentinel can be vendored and audited with zero
// third-party dependencies.
package diff

import (
	"bufio"
	"io"
	"strings"
)

// LineKind classifies a single line inside a hunk.
type LineKind int

const (
	// Context is an unchanged line present in both sides of the diff.
	Context LineKind = iota
	// Added is a line introduced by the patch (prefixed with '+').
	Added
	// Removed is a line deleted by the patch (prefixed with '-').
	Removed
)

