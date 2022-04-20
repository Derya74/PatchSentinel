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

