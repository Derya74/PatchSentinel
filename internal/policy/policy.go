// Package policy defines the configurable rules PatchSentinel evaluates a
// patch against, plus a deterministic loader with sensible built-in defaults.
package policy

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
)

// Policy is the full rule set. All thresholds are inclusive unless noted.
type Policy struct {
	// Version identifies the policy schema for forward compatibility.
	Version int `json:"version"`

	// SensitivePaths are glob patterns (see internal/glob) that mark files
	// requiring extra scrutiny (auth, infra, CI, secrets, etc.).
	SensitivePaths []string `json:"sensitivePaths"`

	// RiskPatterns are regular expressions applied to added lines. Each match
	// contributes its Weight to the risk score.
	RiskPatterns []RiskPattern `json:"riskPatterns"`

	// TestPathPatterns identify files that count as tests for the
	// test-to-code ratio.
	TestPathPatterns []string `json:"testPathPatterns"`

	// Thresholds control verdict computation.
	Thresholds Thresholds `json:"thresholds"`

	// ReviewCost tunes the review-effort model.
	ReviewCost ReviewCost `json:"reviewCost"`
}

// RiskPattern is a named regular expression with a severity weight.
type RiskPattern struct {
	ID      string `json:"id"`
	Regex   string `json:"regex"`
	Message string `json:"message"`
	Weight  int    `json:"weight"`
}

// Thresholds gate the overall verdict.
type Thresholds struct {
	// MaxRiskScore is the highest tolerated cumulative risk score before the
	// verdict fails. Exceeding it is a hard fail.
	MaxRiskScore int `json:"maxRiskScore"`
	// WarnRiskScore raises a warning (not a failure) once reached.
	WarnRiskScore int `json:"warnRiskScore"`
	// MinTestRatio is the minimum acceptable test-to-code line ratio when the
	// change touches non-trivial production code.
	MinTestRatio float64 `json:"minTestRatio"`
	// MinProductionLinesForTestCheck disables the test-ratio gate for tiny
	// changes below this many added production lines.
	MinProductionLinesForTestCheck int `json:"minProductionLinesForTestCheck"`
	// MaxRepetitionRatio flags likely generated/boilerplate code when the
	// fraction of duplicated added lines meets or exceeds this value.
	MaxRepetitionRatio float64 `json:"maxRepetitionRatio"`
	// MaxReviewMinutes fails changes whose estimated review time is too high.
	MaxReviewMinutes int `json:"maxReviewMinutes"`
	// SensitivePathIsWarning downgrades sensitive-path hits from fail to warn.
	SensitivePathIsWarning bool `json:"sensitivePathIsWarning"`
}

// ReviewCost parameters feed a simple, explainable effort model.
type ReviewCost struct {
	// MinutesPerAddedLine is the base cost of reading one new line.
	MinutesPerAddedLine float64 `json:"minutesPerAddedLine"`
	// MinutesPerRemovedLine is the (lower) cost of validating a deletion.
	MinutesPerRemovedLine float64 `json:"minutesPerRemovedLine"`
	// MinutesPerFile accounts for context-switching between files.
	MinutesPerFile float64 `json:"minutesPerFile"`
	// SensitiveMultiplier scales the cost of touching sensitive files.
	SensitiveMultiplier float64 `json:"sensitiveMultiplier"`
