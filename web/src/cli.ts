/**
 * PatchSentinel report CLI (TypeScript / Node, zero dependencies).
 *
 * Reads a PatchSentinel JSON report (produced by the Go CLI) from a file or
 * stdin and renders it as either a self-contained HTML page, a GitHub check
 * summary (Markdown), or GitHub Actions annotations.
 *
 * Usage:
 *   node dist/cli.js --format html   [--in report.json] [--out report.html]
 *   node dist/cli.js --format github [--in report.json]
 *   node dist/cli.js --format annotations
 *   patchsentinel -format json <diff> | node dist/cli.js --format html
 */

