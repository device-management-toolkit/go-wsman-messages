# Contributing to go-wsman-messages

Thank you for your interest in contributing. This document covers the local workflow, the quality gates a PR must clear, and the commit / PR conventions that drive the automated release pipeline. The architectural rules for adding new WSMAN messages are documented in [CLAUDE.md](./CLAUDE.md).

## Local development

### Prerequisites

- Go **1.25** (see `go.mod`). The CI matrix exercises Linux and Windows; write platform-agnostic code, particularly in `pkg/wsman/local/`.
- `gofumpt` — `go install mvdan.cc/gofumpt@latest`. **Must be run before `golangci-lint --fix`** (see workflow below); otherwise `golangci-lint --fix` and `gofumpt` produce conflicting edits that corrupt formatting.
- `golangci-lint` v2 — either install locally per the [official instructions](https://golangci-lint.run/welcome/install/) or invoke via Docker (`docker run --rm -v "$(pwd):/app" -w /app golangci/golangci-lint:latest golangci-lint run -v`). The repo's `.golangci.yml` registers `gofmt`, `gofumpt`, `goimports`, and `gci` under `formatters`, so a verification-only `golangci-lint run` enforces formatting **and** linting in one shot.

### Common commands

```sh
go test ./...                                            # full suite
go test ./pkg/wsman/amt/alarmclock -run TestPositive     # one test
go test -covermode=atomic -coverprofile=coverage.out -race -v ./...

# Formatting + lint — in this order
gofumpt -l -w -extra ./                                  # 1. format first
golangci-lint run --fix --config=./.golangci.yml ./...   # 2. then lint --fix
golangci-lint run --config=./.golangci.yml ./...         # 3. verify clean

go vet ./...
```

**The order matters.** Running `golangci-lint --fix` before `gofumpt` will rewrite imports and line breaks in a way that `gofumpt` then has to undo on the next pass, producing conflicting edits that corrupt formatting. Always: `gofumpt -l -w -extra ./` first, then `golangci-lint --fix`, then a final verification run. `golangci-lint` registers `gofmt`/`gofumpt`/`goimports`/`gci` under `formatters` in `.golangci.yml`, so the final verification run gates both formatting and lint rules.

### Test coverage requirements

Coverage is the primary quality gate for this library: every exported message is a wire-format contract with AMT firmware and a regression can disrupt fleets running `rpc-go` or `console` in production.

- **Minimum: 85% per-package line coverage.** PRs that drop a package below this threshold will be rejected. Add tests in the **same PR** as the code change, not as a follow-up.
- **Goal: 100%.** Treat uncovered branches as defects to be fixed. The narrow exceptions worth annotating and leaving uncovered are: (1) platform-specific syscall failures that can't run on the CI runners, and (2) defensive panics guarding invariants the Go type system already enforces (prefer removing them).
- Inspect gaps with `go tool cover -func=coverage.out | grep -v '100.0%'`. Codecov reports per-package coverage on every PR.
- Every new service needs a `service_test.go` exercising Get / Enumerate / Pull and each custom method, asserting byte-exact XML output. Add the corresponding response fixture under `pkg/wsman/wsmantesting/responses/<namespace>/<lowercase-method>.xml`.

## <a name="commit"></a> Commit Message Guidelines

We have precise rules over how our git commit messages should be formatted. This leads to more readable messages that are easy to follow when looking through the project history and is required for the automated semantic-release pipeline that publishes new versions of the module.

### Commit Message Format

Each commit message consists of a **header**, a **body** and a **footer**. The header has a special format that includes a **type**, a **scope** and a **subject**:

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

The **header** with **type** is mandatory. The **scope** of the header is optional as far as the automated PR checks are concerned, but reviewers **may request** you provide an applicable scope.

Any line of the commit message should be at most ~72 characters for readability on GitHub and in terminal git tools. The hard CI cap (set in `.github/commitlint.config.cjs`) is **200** characters per body line.

The footer should contain a reference to a GitHub issue using the [GitHub closing-keyword syntax](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue) (e.g. `Closes: #1234`, `Fixes: #1234`, or `Resolves: #1234`) so the issue auto-closes when the PR merges.

Example:

```
feat(amt): add CIM_AssociatedPowerManagementService wrapper

Wires a new service on amt.Messages that mirrors the AMT SDK's
AssociatedPowerManagementService class. The wrapper reuses the
base.WSManService Get/Enumerate/Pull primitives and adds typed
ReturnValue lookups for the firmware-side enum.

Closes: #678
```

### Revert

If the commit reverts a previous commit, it should begin with `revert: `, followed by the header of the reverted commit. In the body it should say: `This reverts commit <hash>.`, where the hash is the SHA of the commit being reverted.

### Type

Must be one of the following:

- **feat**: A new feature (triggers a **minor** release).
- **fix**: A bug fix (triggers a **patch** release).
- **perf**: A code change that improves performance (triggers a **patch** release).
- **chore**: Maintenance work not covered by the other types (triggers a **patch** release per `.releaserc.json`).
- **docs**: Documentation only changes (no release).
- **style**: Changes that do not affect the meaning of the code — white-space, formatting, missing semicolons, etc. (no release).
- **refactor**: A code change that neither fixes a bug nor adds a feature (no release).
- **test**: Adding missing tests or correcting existing tests (no release).
- **build**: Changes that affect the build system or external dependencies (no release; dependabot uses `build(deps)` / `build(deps-dev)`).
- **ci**: Changes to the CI configuration files and scripts (no release).
- **revert**: Reverts a previous commit (release impact depends on the reverted commit).

`BREAKING CHANGE:` in the footer triggers a **major** release regardless of type. Use it for any change to the emitted XML envelopes, exported Go types, or package layout — see "Pull Requests practices" below.

### Scope

The conventional scopes used in this repo are listed below. `.github/commitlint.config.cjs` does **not** currently enforce a `scope-enum` allowlist (it sets `body-max-line-length` and `subject-case` only), so a commit with an off-list scope will pass CI — but reviewers may request you align with the list during review for consistency with the historical `git log` and the auto-generated release notes.

- **amt**: A change or addition to an AMT class under `pkg/wsman/amt/`.
- **cim**: A change or addition to a CIM class under `pkg/wsman/cim/`.
- **ips**: A change or addition to an IPS class under `pkg/wsman/ips/`.
- **apf**: A change or addition to the Asynchronous Protocol Framework (`pkg/apf/`).
- **client**: A change to the transport layer (`pkg/wsman/client/`).
- **local**: A change to the LMS/LME local-host transport (`pkg/wsman/local/`).
- **deps**: A change or addition to dependencies (primarily used by dependabot).
- **deps-dev**: A change or addition to developer dependencies (primarily used by dependabot).
- **gh-actions**: A change or addition to GitHub Actions workflows.
- _no scope_: If no scope applies, omit it.

Sub-scopes are permitted when a class warrants more specificity (e.g. `feat(amt/boot): …`), but the outer scope must be one of the above.

### Body

Just as in the **subject**, use the imperative, present tense: "change" not "changed" nor "changes". Detailed guideline ([reference](https://chris.beams.io/posts/git-commit/)):

```
More detailed explanatory text, if necessary. Wrap it to about 72
characters or so. The blank line separating the summary from the
body is critical (unless you omit the body entirely); various tools
like `log`, `shortlog` and `rebase` can get confused if you run the
two together.

Explain the problem that this commit is solving. Focus on why you
are making this change as opposed to how (the code explains that).
Are there side effects or other unintuitive consequences of this
change? Here's the place to explain them.

 - Bullet points are okay, too

 - Typically a hyphen or asterisk is used for the bullet, preceded
   by a single space, with blank lines in between, but conventions
   vary here
```

### Footer

The footer should contain a reference to the GitHub issue this commit **Closes**, **Fixes**, or **Resolves** (e.g. `Closes: #1234`). See [GitHub's closing-keyword syntax](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue) for the keywords that trigger auto-close on merge.

The footer should also contain any information about **Breaking Changes**.

**Breaking Changes** should start with the words `BREAKING CHANGE:` followed by a space or two newlines. Because this module is depended on by `rpc-go` (running on managed devices) and `console`, use a `BREAKING CHANGE:` footer for any change to:

- The bytes of an emitted WSMAN envelope (whitespace, attribute order, namespace prefix all count — they're observable to AMT firmware).
- An exported Go type, method, or field on `wsman.Messages` / any `pkg/wsman/<ns>/<class>` package.
- Package paths or directory layout — consumers import individual sub-packages.

semantic-release reads the footer and cuts a major version automatically.

## Pull Requests practices

- **One concern per PR.** Don't bundle a new IPS message with a `client/` refactor and a `gofumpt` sweep — they belong in separate PRs so the diff that changes the wire format can be reviewed in isolation.
- **Order PRs around the release impact.** Land internal helpers and test scaffolding under `refactor:` / `test:` first; ship the user-visible new message as the `feat:` that triggers the release.
- **Coverage gates must pass.** Run `go test -covermode=atomic -coverprofile=coverage.out ./...` and confirm the affected package(s) remain at ≥85%, with 100% as the goal.
- **Before requesting review:** run `gofumpt -l -w -extra ./` → `golangci-lint run --fix --config=./.golangci.yml ./...` → `golangci-lint run --config=./.golangci.yml ./...` (must be clean) → `go test ./...` (must pass). The gofumpt-before-lint-fix ordering is required; flipping it corrupts formatting.
- The PR title should follow the same format as the [commit header](#commit-message-format) — semantic-release will use it when the PR is squashed.
- The PR author is responsible for merging their own PR after review approval and a green CI run.
- When merging, preserve git linear history. The PR author selects `Rebase and merge` or `Squash and merge`, whichever fits the history best.

## Security

See [SECURITY.md](./SECURITY.md). Do not file public issues for vulnerabilities — report them through the Intel vulnerability handling process linked there.
