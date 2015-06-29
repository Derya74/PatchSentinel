# Changelog

All notable changes to PatchSentinel are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Nothing yet — open an issue describing the signal or policy knob you need
  before it lands here.

## [1.0.0] - 2026-08-31

### Added

- **Go CLI (`patchsentinel`)** - deterministic AI contribution quality gate,
  standard library only:
  - Unified diff parser supporting `git diff`, plain `diff -u`, renames,
    deletes, new files and binary files.
  - Signal engine computing:
    - risk score from configurable weighted regex patterns,
    - generated-code repetition ratio,
    - test-to-code ratio (size-gated),
    - sensitive-path detection via a `**` glob matcher,
    - explainable large-change review-cost estimate,
    - a stable SHA-256 content hash for reproducibility.
  - Policy engine with a built-in default policy and a JSON loader that
    overlays partial policies and rejects unknown fields.
  - Verdict model (`pass` / `warn` / `fail`) mapped to exit codes `0` / `1` / `2`,
    with a `-strict` mode that escalates warnings to failures.
  - Output formats: `json`, `text`, `markdown`, `annotations`.
- **TypeScript renderer (`web/`)** - zero runtime dependencies:
  - Self-contained HTML report (inline CSS, no network assets).
  - GitHub-style Markdown check summary.
  - GitHub Actions annotation formatter (`::error` / `::warning`).
  - Node CLI reading the Go JSON from a file or stdin.
- Example diffs (`clean`, `risky`, `generated`) and a sample strict policy.
- Documentation: rich `README.md` with Mermaid diagrams, `docs/POLICY.md`.
- Tooling: `Makefile`, GitHub Actions CI, `.gitignore`, MIT `LICENSE`.
- Focused test suites for Go (parser, glob, analyzer, policy, render, CLI) and
  TypeScript (renderer, formatter, arg parsing, injection safety).

## [0.9.0] - 2026-05-12

### Added

- `-strict` flag: promote a `warn` verdict to exit code `2` without editing
  the policy, so warnings can block a pipeline outright.
- `annotations` output format emitting `::error` / `::warning` workflow
  commands, with `%` and newline escaping.

### Changed

- The review-cost estimate is now reported per file in the JSON payload, not
  only as a total, so a reviewer can see which file carries the cost.

## [0.8.0] - 2026-01-20

### Added

- Repetition signal: fraction of added meaningful production lines that are
  duplicates of an earlier line in the same change, reported as
  `repetitionRatio` and gated at `maxRepetitionRatio`.
- `examples/generated.diff` - a uniform file that trips the repetition gate.

### Fixed

- Trivial lines (braces, blanks, sub-4-character lines) no longer count
  towards the repetition ratio, which was inflated by boilerplate braces.

## [0.7.0] - 2025-11-04

### Added

- Sensitive-path detection with a dedicated `**` glob matcher
  (`internal/glob`), matching paths segment-wise with `/` separators and
  normalizing Windows `\` to `/`.
- `sensitivePathIsWarning` policy switch so a repo can route sensitive files
  to a domain owner instead of blocking the merge.
