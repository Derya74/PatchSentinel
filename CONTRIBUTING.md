# Contributing to PatchSentinel

Thanks for helping. This project is small on purpose: a standard-library-only
Go core and a zero-dependency TypeScript renderer, both of which have to keep
producing byte-identical output for the same input.

## Ground rules

- **Discuss before you build.** If the shape of the change is still open,
  open an issue first. A pull request that solves a problem nobody agreed on
  usually gets closed.
- **One concern per pull request.** Refactors that arrive bundled with a
  behaviour change are hard to review, and this project reviews its own work
  the same way it asks you to review a diff.
- **Behaviour changes need tests.** A new signal, a changed default, or a
  new output format ships with a test that fails without it.
- **Determinism is the promise.** Same diff plus same policy must give the
  same bytes, on every platform, in every run. No clocks, no randomness, no
  map iteration order leaking into output, no network, no shelling out.

## Getting set up

```bash
go build ./...          # Go core
go test ./...           # Go tests
make test-ts            # type-check, build and test the renderer
make ci                 # vet + both test suites, exactly what CI runs
```

Node 20+ is needed only for the `web/` half. The Go CLI stands alone.

## Working on the code

```
cmd/patchsentinel/      CLI entrypoint: flags, IO, exit codes
internal/diff/          unified diff parser
internal/policy/        default policy + JSON overlay loader
internal/glob/          the ** matcher used by sensitive paths
internal/analyze/       signals and the verdict
internal/render/        json, text, markdown, annotations
web/src/                dependency-free HTML / GitHub renderer
examples/               the three demo diffs + a strict policy
docs/POLICY.md          the full policy and report schema
```

Conventions that are not negotiable here:

- stdlib only in `internal/` - a new dependency has to earn its keep against
  a zero-dependency promise that users rely on;
- sort anything you iterate before it reaches output;
- escape per output format - HTML, Markdown and workflow annotations each
  need their own escaping, and `web/src/report.test.ts` is where that is
  pinned down;
- a false positive gets a policy knob, not a special case in the analyzer.

## Pull requests

1. Branch from `main`.
2. Make `make ci` pass.
3. Add a `CHANGELOG.md` entry under `[Unreleased]`.
4. Fill in the pull request template, including the before/after verdict for
   whatever example you changed behaviour on.

Small pull requests with a clear description get reviewed quickly. Large ones
get split.

## Reporting bugs and security issues

Bugs go to the [bug template](.github/ISSUE_TEMPLATE/bug_report.yml) with a
reproducible diff. Security issues do **not** go to a public issue - see
[`SECURITY.md`](SECURITY.md).

## Code of Conduct

Participation is governed by [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).
