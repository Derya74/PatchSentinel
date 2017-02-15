// Package analyze turns a parsed patch and a policy into a deterministic
// Report: a structured set of evidence (risk signals, repetition, test ratio,
// sensitive paths, review cost) plus a policy verdict.
//
// Determinism is a first-class requirement: given the same patch and policy,
// the JSON output is byte-for-byte identical across runs and platforms. All
// iteration happens over sorted collections and floating-point results are
// rounded to a fixed precision.
package analyze

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/Herricka/PatchSentinel/internal/diff"
	"github.com/Herricka/PatchSentinel/internal/glob"
	"github.com/Herricka/PatchSentinel/internal/policy"
)

// Verdict is the top-level gate decision.
type Verdict string

const (
	// Pass means the change satisfies every gate.
	Pass Verdict = "pass"
	// Warn means the change is acceptable but has notable signals.
	Warn Verdict = "warn"
	// Fail means at least one hard gate was violated.
	Fail Verdict = "fail"
)

// Report is the complete, serializable analysis result.
type Report struct {
	// SchemaVersion identifies the output format.
	SchemaVersion int `json:"schemaVersion"`
	// Tool is the producing tool name.
	Tool string `json:"tool"`
	// PolicyVersion echoes the evaluated policy version.
	PolicyVersion int `json:"policyVersion"`

	Summary  Summary    `json:"summary"`
	Files    []FileStat `json:"files"`
	Risks    []RiskHit  `json:"risks"`
	Findings []Finding  `json:"findings"`
	Verdict  Verdict    `json:"verdict"`
	ExitCode int        `json:"exitCode"`
}

// Summary aggregates headline numbers.
type Summary struct {
	FilesChanged    int     `json:"filesChanged"`
	LinesAdded      int     `json:"linesAdded"`
	LinesRemoved    int     `json:"linesRemoved"`
	ProductionAdded int     `json:"productionAdded"`
	TestAdded       int     `json:"testAdded"`
	TestToCodeRatio float64 `json:"testToCodeRatio"`
	RiskScore       int     `json:"riskScore"`
	RepetitionRatio float64 `json:"repetitionRatio"`
	SensitiveFiles  int     `json:"sensitiveFiles"`
	ReviewMinutes   float64 `json:"reviewMinutes"`
	ContentHash     string  `json:"contentHash"`
}

// FileStat is the per-file breakdown.
type FileStat struct {
	Path        string `json:"path"`
	OldPath     string `json:"oldPath,omitempty"`
	Added       int    `json:"added"`
	Removed     int    `json:"removed"`
	IsTest      bool   `json:"isTest"`
	IsSensitive bool   `json:"isSensitive"`
	IsBinary    bool   `json:"isBinary"`
	Status      string `json:"status"` // added | deleted | renamed | modified
}

// RiskHit is a single risk-pattern match with location context.
type RiskHit struct {
	PatternID string `json:"patternId"`
	Message   string `json:"message"`
	Weight    int    `json:"weight"`
	Path      string `json:"path"`
	Snippet   string `json:"snippet"`
	Count     int    `json:"count"`
}

// Severity classifies a finding.
type Severity string

const (
	SeverityInfo  Severity = "info"
	SeverityWarn  Severity = "warn"
	SeverityError Severity = "error"
)

// Finding is a human-facing gate result explaining why the verdict is what it
// is. Findings are the primary evidence surfaced to reviewers.
type Finding struct {
	ID       string   `json:"id"`
	Severity Severity `json:"severity"`
	Title    string   `json:"title"`
	Detail   string   `json:"detail"`
}

// Analyze evaluates a patch against a policy and returns a deterministic
// Report.
func Analyze(p diff.Patch, pol policy.Policy) Report {
	rep := Report{
		SchemaVersion: 1,
		Tool:          "patchsentinel",
		PolicyVersion: pol.Version,
	}

	// Compile risk patterns once. Invalid regexes are skipped defensively.
	type compiledPattern struct {
		def policy.RiskPattern
		re  *regexp.Regexp
	}
	compiled := make([]compiledPattern, 0, len(pol.RiskPatterns))
	for _, rp := range pol.RiskPatterns {
		re, err := regexp.Compile(rp.Regex)
