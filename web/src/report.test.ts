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
      testToCodeRatio: 0.2,
      riskScore: 30,
      repetitionRatio: 0.1,
      sensitiveFiles: 1,
      reviewMinutes: 8,
      contentHash: "deadbeef",
    },
    files: [
      { path: "a.go", added: 10, removed: 3, isTest: false, isSensitive: false, isBinary: false, status: "modified" },
      {
        path: ".github/workflows/ci.yml",
        added: 2,
        removed: 0,
        isTest: false,
        isSensitive: true,
        isBinary: false,
        status: "modified",
      },
    ],
    risks: [
      { patternId: "hardcoded-secret", message: "Possible secret | leak", weight: 10, path: "a.go", snippet: "key=...", count: 1 },
    ],
    findings: [
      { id: "risk-score", severity: "error", title: "Risk score exceeds maximum", detail: "score 30 >= 25" },
    ],
    verdict,
    exitCode: verdict === "fail" ? 2 : verdict === "warn" ? 1 : 0,
  };
}

test("isReport accepts valid and rejects invalid", () => {
  assert.equal(isReport(sampleReport()), true);
  assert.equal(isReport({}), false);
  assert.equal(isReport(null), false);
  assert.equal(isReport("nope"), false);
});

