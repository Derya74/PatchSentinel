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
