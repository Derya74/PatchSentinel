# Changelog

All notable changes to PatchSentinel are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-08-31

### Added

- **Go CLI (`patchsentinel`)** — deterministic AI contribution quality gate,
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
- **TypeScript renderer (`web/`)** — zero runtime dependencies:
  - Self-contained HTML report (inline CSS, no network assets).
  - GitHub-style Markdown check summary.
  - GitHub Actions annotation formatter (`::error` / `::warning`).
  - Node CLI reading the Go JSON from a file or stdin.
- Example diffs (`clean`, `risky`, `generated`) and a sample strict policy.
- Documentation: rich `README.md` with Mermaid diagrams, `docs/POLICY.md`.
- Tooling: `Makefile`, GitHub Actions CI, `.gitignore`, MIT `LICENSE`.
- Focused test suites for Go (parser, glob, analyzer, policy, render, CLI) and
  TypeScript (renderer, formatter, arg parsing, injection safety).

[1.0.0]: https://example.com/patchsentinel/releases/tag/v1.0.0

> fix: clamp internal/glob to the supported range

> test: cover benches edge cases

> test: regression for the internal/glob overflow

> chore: tidy internal/analyze

> feat: expose internal/render as a library API

> perf: reduce allocations in internal/glob

> test: regression for the internal/diff overflow

> refactor: split cmd/patchsentinel out of the hot path

> docs: document internal/policy

> fix: handle empty internal/render

> docs: document internal/diff

> chore: bump deps for internal/glob

> feat: add scripts handling
