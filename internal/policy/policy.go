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

