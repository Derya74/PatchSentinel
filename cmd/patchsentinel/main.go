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
