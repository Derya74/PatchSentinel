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
		{"**/*Dockerfile*", "build/Dockerfile.prod", true},
		{"src/*.go", "src/app.go", true},
		{"src/*.go", "src/sub/app.go", false},
		{"**/*.env*", "config/.env.local", true},
		{"a/b?c", "a/bxc", true},
		{"a/b?c", "a/bxxc", false},
		{"**/migrations/**", "db/migrations/001_init.sql", true},
	}
	for _, c := range cases {
		if got := Match(c.pattern, c.name); got != c.want {
			t.Errorf("Match(%q, %q) = %v, want %v", c.pattern, c.name, got, c.want)
		}
	}
}

func TestMatchWindowsSeparators(t *testing.T) {
	if !Match("**/auth/**", "pkg\\auth\\login.go") {
		t.Errorf("windows separators should be normalized")
	}
}
