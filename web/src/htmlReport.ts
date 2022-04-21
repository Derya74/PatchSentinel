/**
 * renderHtml turns a PatchSentinel Report into a single self-contained HTML
 * document. No external assets, scripts, or network calls are used: all CSS is
 * inlined so the report can be attached to a CI artifact and opened offline.
 */

import { Report, FileStat, RiskHit, Finding } from "./types.js";
import {
  escapeHtml,
  ratio,
  minutes,
  verdictMeta,
  severityEmoji,
} from "./util.js";

export function renderHtml(report: Report): string {
  const v = verdictMeta(report.verdict);
  const s = report.summary;

  const metrics: Array<[string, string]> = [
    ["Files changed", String(s.filesChanged)],
    ["Lines added", `+${s.linesAdded}`],
    ["Lines removed", `-${s.linesRemoved}`],
    ["Production added", String(s.productionAdded)],
    ["Test added", String(s.testAdded)],
    ["Test-to-code ratio", ratio(s.testToCodeRatio)],
    ["Risk score", String(s.riskScore)],
    ["Repetition ratio", ratio(s.repetitionRatio)],
    ["Sensitive files", String(s.sensitiveFiles)],
    ["Est. review time", minutes(s.reviewMinutes)],
  ];

  return `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>PatchSentinel Report — ${v.label}</title>
<style>
${styles()}
</style>
</head>
<body>
<main class="wrap">
