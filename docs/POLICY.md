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
    "warnRiskScore": 10,
    "minTestRatio": 0.25,
    "minProductionLinesForTestCheck": 40,
    "maxRepetitionRatio": 0.60,
    "maxReviewMinutes": 90,
    "sensitivePathIsWarning": false
  },
  "reviewCost": {
    "minutesPerAddedLine": 0.30,
    "minutesPerRemovedLine": 0.10,
    "minutesPerFile": 2.0,
    "sensitiveMultiplier": 1.5
  }
}
```

## Glob syntax

Path patterns (`sensitivePaths`, `testPathPatterns`) support:

| Token | Meaning                                            |
|-------|----------------------------------------------------|
| `*`   | Any run of characters **within** a path segment    |
