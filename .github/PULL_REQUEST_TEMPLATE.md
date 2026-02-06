<!-- Thanks for helping. One concern per pull request, discussed before it is built. -->

## What this changes

<!-- One or two sentences. What does the reader see, gate on, or maintain differently? -->

## Why

<!-- Link the issue or the discussion this responds to. -->

## Verification

<!-- Exactly what you ran, and the output you got. -->

```text
$ make ci
```

## Checklist

- [ ] `go build ./...` and `go test ./...` pass.
- [ ] `make test-ts` passes, if `web/` is touched.
- [ ] `make ci` passes - this is exactly what CI runs.
- [ ] A behaviour change ships with a test that fails without it.
- [ ] Anything iterated before it reaches output is sorted, so the same input still produces the same bytes.
- [ ] Values landing in HTML, Markdown or a workflow annotation are escaped for that format.
- [ ] `CHANGELOG.md` has an entry under `[Unreleased]`.
- [ ] No new runtime dependency in `internal/`.

## Before / after

<!-- For a verdict or output change, paste both. Leave as-is otherwise. -->

| Diff | Before | After |
|---|---|---|
| `examples/risky.diff` | | |

## Notes for the reviewer

<!-- Anything non-obvious: a trade-off you made, a case you deliberately left out. -->

<!-- draft note 1229 -->
