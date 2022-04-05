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
  const s = report.summary;

  const lines: string[] = [];
  lines.push(`## PatchSentinel — ${v.emoji} ${v.label}`);
  lines.push("");
  lines.push("| Metric | Value |");
  lines.push("|---|---|");
  lines.push(`| Files changed | ${s.filesChanged} |`);
  lines.push(`| Lines added / removed | +${s.linesAdded} / -${s.linesRemoved} |`);
  lines.push(`| Test-to-code ratio | ${ratio(s.testToCodeRatio)} |`);
  lines.push(`| Risk score | ${s.riskScore} |`);
