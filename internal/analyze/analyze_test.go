package analyze

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Herricka/PatchSentinel/internal/diff"
	"github.com/Herricka/PatchSentinel/internal/policy"
)

func parse(t *testing.T, s string) diff.Patch {
	t.Helper()
	p, err := diff.Parse(strings.NewReader(s))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return p
}

const cleanDiff = `diff --git a/src/util.go b/src/util.go
--- a/src/util.go
+++ b/src/util.go
@@ -1,3 +1,6 @@
 package util
 
+func Add(a, b int) int {
+	return a + b
+}
diff --git a/src/util_test.go b/src/util_test.go
new file mode 100644
--- /dev/null
+++ b/src/util_test.go
@@ -0,0 +1,5 @@
+package util
+import "testing"
+func TestAdd(t *testing.T) {
+	if Add(1,2) != 3 { t.Fail() }
+}
`

func TestAnalyzeCleanPasses(t *testing.T) {
	rep := Analyze(parse(t, cleanDiff), policy.Default())
	if rep.Verdict != Pass {
		t.Errorf("verdict = %s, want pass", rep.Verdict)
	}
	if rep.ExitCode != 0 {
		t.Errorf("exit = %d, want 0", rep.ExitCode)
	}
	if rep.Summary.TestAdded == 0 {
		t.Errorf("expected test lines counted")
	}
	if rep.Summary.ProductionAdded == 0 {
		t.Errorf("expected production lines counted")
	}
}

const secretDiff = `diff --git a/src/config.go b/src/config.go
--- a/src/config.go
+++ b/src/config.go
@@ -1,2 +1,5 @@
 package config
 
+var apiKey = "sk_live_abcdef1234567890"
+var password = "hunter2hunter2"
+func run() { eval("danger") }
`

func TestAnalyzeSecretFails(t *testing.T) {
	rep := Analyze(parse(t, secretDiff), policy.Default())
	if rep.Verdict != Fail {
		t.Errorf("verdict = %s, want fail (risk=%d)", rep.Verdict, rep.Summary.RiskScore)
	}
	if rep.Summary.RiskScore < policy.Default().Thresholds.MaxRiskScore {
		t.Errorf("risk score %d should exceed max", rep.Summary.RiskScore)
	}
	// Ensure the hardcoded-secret pattern was detected.
	found := false
	for _, r := range rep.Risks {
		if r.PatternID == "hardcoded-secret" {
			found = true
		}
	}
	if !found {
		t.Errorf("hardcoded-secret not detected; risks=%+v", rep.Risks)
	}
}

func sensitiveDiff() string {
	return `diff --git a/.github/workflows/deploy.yml b/.github/workflows/deploy.yml
--- a/.github/workflows/deploy.yml
+++ b/.github/workflows/deploy.yml
@@ -1,2 +1,3 @@
 name: deploy
 on: push
+run: echo hi
`
}

func TestAnalyzeSensitivePath(t *testing.T) {
