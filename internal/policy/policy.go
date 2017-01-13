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
