/**
 * Tests for the TypeScript renderer/formatter using Node's built-in test
 * runner (node:test) and assertions (node:assert). No third-party test deps.
 */

import { test } from "node:test";
import assert from "node:assert/strict";

import { Report, isReport } from "./types.js";
import { renderHtml } from "./htmlReport.js";
import { renderCheckSummary, renderAnnotations } from "./githubCheck.js";
import { escapeHtml, escapePipe } from "./util.js";
import { parseArgs, render, HelpRequested } from "./cli.js";

function sampleReport(verdict: Report["verdict"] = "fail"): Report {
  return {
    schemaVersion: 1,
    tool: "patchsentinel",
    policyVersion: 1,
    summary: {
      filesChanged: 2,
      linesAdded: 12,
      linesRemoved: 3,
      productionAdded: 10,
      testAdded: 2,
