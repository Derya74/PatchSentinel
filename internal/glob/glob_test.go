package glob

import "testing"

func TestMatch(t *testing.T) {
	cases := []struct {
		pattern string
		name    string
		want    bool
	}{
		{"**/*_test.go", "internal/diff/diff_test.go", true},
		{"**/*_test.go", "main.go", false},
		{"**/.github/workflows/**", ".github/workflows/ci.yml", true},
		{"**/.github/workflows/**", "src/.github/workflows/deploy.yaml", true},
		{"**/auth/**", "pkg/auth/login.go", true},
		{"**/auth/**", "pkg/authz/rbac.go", false},
		{"*.tf", "main.tf", true},
		{"**/*.tf", "infra/prod/main.tf", true},
		{"**/Dockerfile", "build/Dockerfile", true},
