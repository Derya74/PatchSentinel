/**
 * GitHub-style formatters. `renderCheckSummary` builds a Markdown body suitable
 * for the GitHub Checks API `output.summary` field or a PR comment.
 * `renderAnnotations` emits GitHub Actions workflow commands (::error /
 * ::warning) that surface findings inline in CI logs.
 *
 * Zero runtime dependencies.
 */

import { Report, Finding, RiskHit } from "./types.js";
import {
  escapePipe,
  ratio,
  minutes,
  verdictMeta,
  severityEmoji,
} from "./util.js";

export function renderCheckSummary(report: Report): string {
  const v = verdictMeta(report.verdict);
