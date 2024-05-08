// Command patchsentinel is a deterministic AI contribution quality gate.
//
// It reads one or more unified diff files (or stdin), analyzes them against a
// policy, and emits structured evidence plus a pass/warn/fail verdict. The
// process exit code mirrors the verdict (0 pass, 1 warn, 2 fail) so it can be
// wired directly into CI pipelines.
//
// Usage:
//
//	patchsentinel [flags] [diff-file ...]
//
// Flags:
//
//	-format   output format: json | text | markdown | annotations (default json)
//	-policy   path to a JSON policy file (defaults to built-in policy)
//	-out      write output to a file instead of stdout
//	-strict   treat warnings as failures (exit 2 on warn)
//	-version  print version and exit
//
// With no diff-file arguments, the diff is read from stdin.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Herricka/PatchSentinel/internal/analyze"
	"github.com/Herricka/PatchSentinel/internal/diff"
	"github.com/Herricka/PatchSentinel/internal/policy"
	"github.com/Herricka/PatchSentinel/internal/render"
)

// version is overridable at build time via -ldflags "-X main.version=...".
var version = "1.0.0"

func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

// run is the testable core of the CLI. It returns the process exit code.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("patchsentinel", flag.ContinueOnError)
