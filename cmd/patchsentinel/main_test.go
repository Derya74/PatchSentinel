package main

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0o644)
}

const passDiff = `diff --git a/src/util.go b/src/util.go
--- a/src/util.go
+++ b/src/util.go
@@ -1,2 +1,4 @@
 package util
 
+func Ok() bool { return true }
`

func TestRunJSONPass(t *testing.T) {
	var out, errBuf bytes.Buffer
	code := run([]string{"-format", "json"}, strings.NewReader(passDiff), &out, &errBuf)
	if code != 0 {
		t.Fatalf("exit = %d, stderr=%s", code, errBuf.String())
	}
	var parsed map[string]any
	if err := json.Unmarshal(out.Bytes(), &parsed); err != nil {
		t.Fatalf("output not valid json: %v\n%s", err, out.String())
	}
	if parsed["verdict"] != "pass" {
		t.Errorf("verdict = %v, want pass", parsed["verdict"])
	}
}

func TestRunVersion(t *testing.T) {
	var out, errBuf bytes.Buffer
	code := run([]string{"-version"}, strings.NewReader(""), &out, &errBuf)
	if code != 0 {
		t.Fatalf("exit = %d", code)
	}
	if !strings.Contains(out.String(), "patchsentinel") {
		t.Errorf("version output = %q", out.String())
	}
}

func TestRunUnknownFormat(t *testing.T) {
	var out, errBuf bytes.Buffer
	code := run([]string{"-format", "xml"}, strings.NewReader(passDiff), &out, &errBuf)
	if code != 2 {
		t.Errorf("exit = %d, want 2 for unknown format", code)
	}
}

const failDiff = `diff --git a/c.go b/c.go
--- a/c.go
+++ b/c.go
