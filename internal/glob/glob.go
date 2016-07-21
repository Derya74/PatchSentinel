// Package glob implements a small, dependency-free path matcher supporting the
// "**" (match across separators) wildcard in addition to the "*" and "?"
// wildcards understood by path.Match.
//
// Matching is done by translating the glob into a regular expression once and
// caching nothing (patterns are few and compilation is cheap), keeping the
// implementation simple and deterministic. Paths are compared using forward
// slashes; callers should normalize separators beforehand.
package glob

import (
	"regexp"
	"strings"
	"sync"
)

var (
	mu    sync.Mutex
	cache = map[string]*regexp.Regexp{}
)

// Match reports whether name matches pattern. It always uses '/' as the path
// separator. An invalid pattern never matches (returns false) rather than
