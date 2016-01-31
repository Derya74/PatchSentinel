<!-- PatchSentinel — the review desk between a diff and your main branch. -->

<p align="center">
  <img src="docs/assets/hero-gates.svg" alt="Diff lines riding a conveyor through the risk, tests, repetition and sensitive-path gates, exiting stamped pass, warn or fail." width="100%">
</p>

<h1 align="center">PatchSentinel</h1>

<p align="center">
  <em>An editorial review desk for diffs. Every change lines up, walks the gates, and leaves with a stamp — <strong>pass</strong>, <strong>warn</strong>, or <strong>fail</strong> — that CI can trust because it was earned the same way every time.</em>
</p>

<p align="center">
  <code>stdlib-only Go core</code> · <code>zero-dependency TypeScript renderer</code> · <code>no network</code> · <code>no LLM in the loop</code> · <code>byte-identical output</code>
</p>

---

Maintainers don't drown in bad code. They drown in **plausible** code — large, tidy-looking diffs produced in seconds that take a human an hour to actually vouch for. The secret in the base64 blob, the `InsecureSkipVerify: true` three files down, the forty near-identical getters, the 300 new lines with zero tests: they all *look* fine at a glance. PatchSentinel reads the whole diff, measures the boring-but-load-bearing stuff, and puts the risky material in front of your eyes before it reaches your users.

It does **not** claim to detect "AI-written code" — that claim isn't honest and this tool won't fake it. It measures transparent, explainable properties of a unified diff — repetition, risk-pattern hits, test-to-code balance, sensitive-path touches, review cost — and turns them into a reproducible verdict. See [Heuristics, honestly](#heuristics-honestly) for exactly where the lines are.

## Navigation

[60-second tour](#the-60-second-tour) · [Install](#install--build) · [Three diffs, three verdicts](#three-diffs-three-verdicts) · [Anatomy of a score](#anatomy-of-a-score) · [Policy knobs](#policy-knobs) · [Exit codes](#exit-codes) · [Report formats](#report-formats) · [GitHub Actions](#github-actions) · [Triage playbooks](#triage-playbooks) · [Heuristics, honestly](#heuristics-honestly) · [Edge cases](#edge-cases) · [Limitations](#limitations) · [Customization](#customization-recipes) · [Architecture](#architecture) · [Development](#development) · [Roadmap](#roadmap) · [License](#license)

## The 60-second tour

```bash
make build                                   # -> ./bin/patchsentinel
git diff origin/main... | ./bin/patchsentinel -format text
```

The core is a **Go CLI** built from the standard library alone: parse a unified diff (files or stdin), evaluate it against a policy, print evidence, exit with a code that mirrors the verdict. A separate dependency-free **TypeScript renderer** turns that JSON into a self-contained HTML page or GitHub check output. Two languages, one schema, zero runtime dependencies.

<p align="center">
  <img src="docs/assets/cockpit-report.svg" alt="A terminal window typing out the risky-example report: verdict FAIL, risk score 42, two sensitive files, exit 2, with live gauges." width="100%">
</p>

## Install & build

Requires **Go ≥ 1.24** and **Node ≥ 20** (Node only for HTML/GitHub rendering; the CLI stands alone).

```bash
make build      # compile the Go CLI into ./bin/patchsentinel
make test       # run Go + TypeScript test suites
make demo       # regenerate the example HTML + Markdown reports
make help       # list every target

# or raw:
go build -o bin/patchsentinel ./cmd/patchsentinel
cd web && npm install && npm run build      # dev-only deps; runtime stays zero-dep
```

