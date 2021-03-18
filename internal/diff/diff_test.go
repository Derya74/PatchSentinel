package diff

import (
	"strings"
	"testing"
)

const gitDiff = `diff --git a/src/app.go b/src/app.go
index 1234567..89abcde 100644
--- a/src/app.go
+++ b/src/app.go
@@ -1,4 +1,6 @@
 package main
 
+import "fmt"
+
 func main() {
-	println("hi")
+	fmt.Println("hello")
 }
diff --git a/README.md b/README.md
new file mode 100644
index 0000000..1111111
--- /dev/null
+++ b/README.md
@@ -0,0 +1,2 @@
+# Title
+body
`

func TestParseGitDiff(t *testing.T) {
	p, err := Parse(strings.NewReader(gitDiff))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(p.Files) != 2 {
		t.Fatalf("want 2 files, got %d", len(p.Files))
	}

	app := p.Files[0]
	if app.Path() != "src/app.go" {
		t.Errorf("path = %q", app.Path())
	}
	if app.Added() != 3 {
		t.Errorf("added = %d, want 3", app.Added())
	}
	if app.Removed() != 1 {
		t.Errorf("removed = %d, want 1", app.Removed())
