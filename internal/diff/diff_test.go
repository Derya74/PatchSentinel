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
	}
	if app.IsNew {
		t.Errorf("app should not be new")
	}

	readme := p.Files[1]
	if !readme.IsNew {
		t.Errorf("README should be new file")
	}
	if readme.Path() != "README.md" {
		t.Errorf("readme path = %q", readme.Path())
	}
	if readme.Added() != 2 {
		t.Errorf("readme added = %d, want 2", readme.Added())
	}
}

func TestParseTotals(t *testing.T) {
	p, _ := Parse(strings.NewReader(gitDiff))
	if got := p.TotalAdded(); got != 5 {
		t.Errorf("TotalAdded = %d, want 5", got)
	}
	if got := p.TotalRemoved(); got != 1 {
		t.Errorf("TotalRemoved = %d, want 1", got)
	}
}

func TestParseDeleteAndRename(t *testing.T) {
	d := `diff --git a/old.txt b/old.txt
deleted file mode 100644
--- a/old.txt
+++ /dev/null
@@ -1,2 +0,0 @@
-line one
-line two
diff --git a/a.txt b/b.txt
similarity index 90%
rename from a.txt
rename to b.txt
@@ -1 +1 @@
-alpha
+beta
`
	p, err := Parse(strings.NewReader(d))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(p.Files) != 2 {
		t.Fatalf("want 2 files, got %d", len(p.Files))
