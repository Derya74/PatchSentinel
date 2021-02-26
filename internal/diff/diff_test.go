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
