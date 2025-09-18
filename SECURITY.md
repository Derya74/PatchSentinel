# Security Policy

PatchSentinel reads untrusted diffs and prints them back to you. That makes
the parser a security-relevant surface, and reports about it are taken
seriously.

## Supported versions

| Version | Supported |
|---|---|
| `main` | yes |
| `1.0.x` | yes |
| `< 1.0` | no - upgrade |

Fixes land on `main` first and are released from there.

## Reporting a vulnerability

**Do not open a public issue.** Send a private report through GitHub's
[private vulnerability reporting](https://github.com/Herricka/PatchSentinel/security/advisories/new)
page (Security -> Report a vulnerability).

Please include:

- what you ran and what happened,
- the input diff, or a minimal reproduction,
- the impact you believe it has.

You can expect an acknowledgement within a week. Once a fix is ready you get
a credit in the advisory unless you ask otherwise, and a CVE is requested
only where the impact warrants one.

## What counts as a vulnerability here

In scope:

- a panic, hang, or unbounded memory or CPU consumption when parsing a
  crafted or merely large diff;
- a path traversal or file read outside the paths you asked for;
- injection through a rendered report - HTML or Markdown that a workflow
  would interpret, or an annotation that escapes into a workflow command;
- a policy file that is loaded from somewhere other than where you pointed,
  or that is evaluated with more permissions than intended.

Out of scope:

- a rule that fires on code you consider safe (that is a false positive -
  open a normal issue, the tuning answer is a policy overlay);
- a verdict you disagree with;
- anything requiring you to write a deliberately malicious policy file
  yourself.

## Hardening notes

PatchSentinel reads from a file or standard input and never opens a network
connection, shells out, or evaluates the diff it is given. Diff content is
carried as data, and every value interpolated into an HTML report, a Markdown
report, or a workflow annotation is escaped for that format.
