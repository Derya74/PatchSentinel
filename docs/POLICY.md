# PatchSentinel Policy Reference

A **policy** is a JSON document describing the rules PatchSentinel evaluates a
patch against. If you pass `-policy path.json`, its fields are overlaid on top
of the built-in [default policy](#default-policy): any field you omit keeps its
default value. Unknown fields are rejected so typos fail loudly.

```bash
patchsentinel -policy examples/policy.strict.json examples/risky.diff
```

## Schema

```jsonc
{
  "version": 1,
  "sensitivePaths":   ["**/.github/workflows/**", "..."],
  "testPathPatterns": ["**/*_test.go", "..."],
  "riskPatterns": [
    { "id": "eval-usage", "regex": "\\beval\\s*\\(", "message": "…", "weight": 8 }
  ],
  "thresholds": {
    "maxRiskScore": 25,
