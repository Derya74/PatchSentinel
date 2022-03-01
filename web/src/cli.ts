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

import { readFileSync, writeFileSync } from "node:fs";
import { pathToFileURL } from "node:url";
import { isReport, Report } from "./types.js";
import { renderHtml } from "./htmlReport.js";
import { renderCheckSummary, renderAnnotations } from "./githubCheck.js";

interface Options {
  format: "html" | "github" | "annotations";
  input?: string;
  output?: string;
}

