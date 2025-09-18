# Support

PatchSentinel is maintained in the open by one person with help from
occasional contributors. Here is where to take a question, and what to
expect.

## Where to ask

| Your question is about | Use |
|---|---|
| A bug, or a diff that gets the wrong verdict | [Bug report](https://github.com/Herricka/PatchSentinel/issues/new?template=bug_report.yml) |
| A signal, policy knob or output format you need | [Feature request](https://github.com/Herricka/PatchSentinel/issues/new?template=feature_request.yml) |
| "Is this behaviour intended?" | [Discussions](https://github.com/Herricka/PatchSentinel/discussions) |
| How to use or change PatchSentinel | [`README.md`](README.md) and [`docs/POLICY.md`](docs/POLICY.md) |
| A security problem | **Not** a public issue — see [`SECURITY.md`](SECURITY.md) |

Search the existing issues first; both a wrong-verdict report and a "how do I"
question usually already have an answer there.

## What to include in a bug report

A verdict is only as useful as the evidence you attach. Please include:

- the PatchSentinel version (`patchsentinel -version` or the release tag),
- the exact command you ran and the flags you passed,
- the diff, or the smallest excerpt of it that still reproduces the problem,
- the output you got and the output you expected.

A diff that trips a rule by design is not a bug. If a rule fired on
`eval(` inside a comment, the answer is in
[Heuristics, honestly](README.md#heuristics-honestly): the patterns are
regexes over added lines and they match shapes, not meaning.

## What to expect

- I read every issue. Comments usually arrive within a few days; a fix
  depends on whether it is a one-liner or a design question.
- Reproducible bugs get fixed first. Anything that changes a verdict ships
  with a test.
- Feature requests are discussed before they are built. Open an issue rather
  than a pull request if the shape of the change is still open.
- I do not accept feature-request pull requests that arrive before the
  discussion. Ask first; the answer is usually yes, and it is faster.

## Before you file

```bash
go build ./...
go test ./...
make ci          # what CI runs
```

If your checkout fails there, that is the bug - please report it.
