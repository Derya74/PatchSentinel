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
