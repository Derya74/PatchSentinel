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
