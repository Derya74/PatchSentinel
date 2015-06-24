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
