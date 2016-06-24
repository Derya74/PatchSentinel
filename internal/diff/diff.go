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

// Line is a single physical line inside a hunk.
type Line struct {
	Kind LineKind
	Text string
}

// Hunk is a contiguous change region within a file, delimited by an
// "@@ -a,b +c,d @@" header in the unified format.
type Hunk struct {
	Header string
	Lines  []Line
}

// FileChange represents all edits applied to a single file path.
type FileChange struct {
	// OldPath and NewPath are the pre- and post-image paths. For renames
	// they differ; for creates OldPath is "/dev/null"; for deletes NewPath
	// is "/dev/null".
	OldPath string
	NewPath string
	// IsNew, IsDelete and IsRename are derived from the file header lines.
	IsNew    bool
	IsDelete bool
	IsRename bool
