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
