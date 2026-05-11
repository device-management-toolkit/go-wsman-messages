# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

The filename is historical; the content is tool-neutral and applies to any AI coding assistant (Claude Code, Codex, Cursor, Aider, Continue, Gemini CLI, GitHub Copilot, etc.). `AGENTS.md` and `.github/copilot-instructions.md` are pointers to this file — keep edits here.

## Overview

`github.com/device-management-toolkit/go-wsman-messages/v2` is a Go library that constructs WS-Management (WSMAN) SOAP envelopes addressed to Intel® Active Management Technology (AMT) firmware, sends them over the appropriate transport (HTTP Digest, TLS, or raw TCP for redirection / CIRA tunnels), and unmarshals the AMT response into typed Go structs. It is the Go-side analogue of the TypeScript [`@device-management-toolkit/wsman-messages`](https://github.com/device-management-toolkit/wsman-messages) library, but unlike that one it owns the **transport** as well as the **envelope** — callers hand it a `client.Parameters` and get back parsed data, not strings.

Public surface is the `wsman.Messages` struct (`pkg/wsman/messages.go`), which embeds three namespace sub-Messages: `AMT`, `CIM`, `IPS`. Consumers — [`rpc-go`](https://github.com/device-management-toolkit/rpc-go) (the in-host RPC client) and [`console`](https://github.com/device-management-toolkit/console) — construct one of these and call e.g. `m.AMT.GeneralSettings.Get()`. Full AMT SDK reference: <https://www.intel.com/content/www/us/en/developer/tools/active-management-technology-sdk/overview.html>.

## Commands

- `go test ./...` — runs the full suite. The fast path during development. CI uses `go test -covermode=atomic -coverprofile=coverage.out -race -v ./...`; run with `-race` locally before opening a PR if you touched `pkg/wsman/client/` or `pkg/apf/` (anywhere goroutines or shared connection state are in play).
- Single package: `go test ./pkg/wsman/amt/alarmclock/...`. Single test: `go test ./pkg/wsman/amt/alarmclock -run TestPositiveAMT_AlarmClockService`. Add `-v` to see sub-test names.
- **Formatting + lint workflow (order matters):**
  1. `gofumpt -l -w -extra ./` — **run this first.** Install with `go install mvdan.cc/gofumpt@latest`.
  2. `golangci-lint run --fix --config=./.golangci.yml ./...` — apply lint auto-fixes.
  3. `golangci-lint run --config=./.golangci.yml ./...` — verify clean.

  Running `golangci-lint --fix` **before** `gofumpt` will rewrite imports / lines in a way that gofumpt then has to undo, leaving conflicting edits and broken formatting. Always gofumpt first, then lint-fix.

  The v2 config registers `gofmt`, `gofumpt`, `goimports`, and `gci` under `formatters` (gci section order: `standard → default → prefix(github.com/device-management-toolkit/go-wsman-messages/v2) → blank → dot → alias → localmodule`), so a verification-only run (`golangci-lint run …` without `--fix`) is the single command that gates both formatting and the lint rules. CI runs `reviewdog/action-golangci-lint`; for an exact reproduction locally use `docker run --rm -v "$(pwd):/app" -w /app golangci/golangci-lint:latest golangci-lint run -v`. Notable enabled checkers: `errorlint`, `exhaustive`, `forbidigo`, `gochecknoinits`, `nlreturn`, `wsl`, `staticcheck`, `unused`, `nolintlint`, `predeclared`, `thelper`, `tparallel`. `goconst` triggers on 2+ occurrences of a 2+ character string — extract constants liberally. `linters.default: none` means only the linters explicitly listed in `linters.enable` run; `errcheck` is intentionally not enabled despite the settings block in `.golangci.yml`.
- `go vet ./...` — runs in CI under the `formatting` job alongside the gofmt check.
- `go install github.com/jstemmer/go-junit-report/v2@latest` — required only if you want to reproduce the CI's JUnit-formatted test output locally; not needed for normal development.

Toolchain: Go **1.25** (`go.mod`). CI builds on `windows-2019`, `windows-2022`, `ubuntu-22.04`, `ubuntu-20.04` — write platform-agnostic code.

### Test coverage policy

Coverage is the primary quality gate for this library because every exported message is a wire-format contract with AMT firmware and a regression can brick a fleet. Targets:

- **Minimum 85%** package-level coverage. PRs that drop a package below this threshold will be rejected — add tests in the same PR as the code change, not as a follow-up.
- **Goal 100%.** Treat any uncovered branch as a defect to be fixed, not an accepted gap. The exceptions worth annotating with a comment and leaving uncovered are:
  - Platform-specific syscall failures that can't be exercised on the CI runners.
  - `panic` arms guarding invariants the type system already prevents (prefer removing them — see "Implementation guidelines" below).
- Run `go test -covermode=atomic -coverprofile=coverage.out ./...` then `go tool cover -func=coverage.out | grep -v '100.0%'` to find the gaps. Codecov runs on every push to `main` and every PR.
- Every new service must ship a `service_test.go` (or equivalent) covering Get/Enumerate/Pull and every custom method, asserting against the **byte-exact** XML input emitted by the service. The fixture for the corresponding response lives at `pkg/wsman/wsmantesting/responses/<namespace>/<lowercased-method>.xml` — the `wsmantesting.MockClient` reads it from disk based on `client.CurrentMessage`, so the fixture file must exist or the test will fail with an open-file error rather than an assertion mismatch.

## Architecture

### Top-level structure

```
pkg/
  amterror/       — AMT-specific HTTP / SOAP fault decoding shared by client and services
  apf/            — Asynchronous Protocol Framework: bytestream framing used by CIRA tunnels
  common/         — Cross-namespace XML response shapes (EnumerateResponse, PullResponse envelope wrappers)
  config/         — YAML-tagged remote-management Configuration / RemoteManagement structs (provisioning input schema)
  security/       — Encrypt/decrypt helpers and OS-keyring-backed (`go-keyring`) Storage for secret material
  wsman/
    messages.go   — wsman.Messages: the top-level entry point exposed to consumers
    types.go      — top-level type aliases / interface shapes
    amt/ cim/ ips/  — one directory per WSMAN namespace (see below)
    base/         — generic WSManService[T] used by every service; the Go equivalent of the JS Base class
    client/       — HTTP Digest + TLS + raw-TCP transports, CIRA channel wiring
    common/       — cross-namespace response models (EnumerateResponse, PullResponse)
    wsmantesting/ — MockClient + ExpectedResponse helper + recorded XML response fixtures
internal/
  message/        — WSManMessageCreator (XML envelope assembly) and message.Base (low-level Get/Put/Enum/Pull primitives); not part of the public API
```

### The namespace package shape

Each AMT/CIM/IPS namespace is built from one directory per WSMAN class. Every class directory follows the **same file shape** — keep them parallel and reach for the same names:

- `decoder.go` — string constants for resource URIs (`AMTAlarmClockService = "AMT_AlarmClockService"`), method names (`AddAlarm = "AddAlarm"`), and `ReturnValue` string maps. The decoder file is the canonical place for "what does integer N from the firmware mean" lookups.
- `decoder_test.go` — tests for the `ReturnValue.String()` lookups and any other helpers in `decoder.go`.
- `types.go` — Go structs for `Response`, `Body`, the entity itself (e.g. `AlarmClockService`), `PullResponse`, custom-method input/output types, plus enum-style typed integers (e.g. `type ReturnValue int`). Preserve doc comments that map integer values to firmware-side semantics — they are the only documentation downstream consumers have.
- `marshal.go` — JSON / YAML marshallers on `Response` (`.JSON()`, `.YAML()`). Almost always boilerplate; copy from a neighbouring package.
- `service.go` — the `Service` (or `Settings`, `Table`, etc.) struct, which embeds `base.WSManService[Response]` and adds any custom methods on top. The constructor name is `New<ClassName>WithClient(creator, client)` and **must** be wired into the namespace's `messages.go`.
- `service_test.go` (or `<file>_test.go`) — table-driven `TestPositive<Class>` tests that call each method, capture `response.XMLInput`, and `assert.Equal` it against `wsmantesting.ExpectedResponse(messageID, ...)`. See "Authoring a new service" below.

The namespace's `messages.go` declares a `Messages` struct with one field per service and a `NewMessages(client)` constructor that instantiates every service against a shared `*message.WSManMessageCreator`. **Every new service must be added to that `Messages` struct and constructor** — that's the entry point consumers reach for.

### Generic service base

`pkg/wsman/base/wsmanservice.go` defines `WSManService[T any]`:

```go
type WSManService[T any] struct {
    Base message.Base
}

func NewService[T any](creator *message.WSManMessageCreator, resourceURI string, client client.WSMan) WSManService[T]

func (s WSManService[T]) Get() (T, error)               // GET with no selector
func (s WSManService[T]) GetByName(name string) (T, error)
func (s WSManService[T]) GetByInstanceID(id string) (T, error)
func (s WSManService[T]) Enumerate() (T, error)         // returns EnumerationContext inside T
func (s WSManService[T]) Pull(ctx string) (T, error)    // follow-up pull using the context
func (s WSManService[T]) Put(request any) (T, error)    // PUT with namespace-injection on the request
```

`Get`/`Enumerate`/`Pull`/`Put` build the XML via `message.Base`, execute it through the supplied `client.WSMan`, then `xml.Unmarshal` the response into `T`. They also use reflection to inject `*client.Message` into a `Message` field on `T` (so callers can inspect the raw XML in/out) and, for `Put`, to set an `H` field on the request struct to the right schema prefix (`AMTSchema + AMT_GeneralSettings`, etc.). This is the Go analogue of the TypeScript library's `Base` class — services should reuse these whenever the operation shape matches.

For methods AMT models with `*_INPUT` shapes that `base.WSManService[T]` doesn't expose directly, prefer a helper inside `internal/message/` (next to `CreateBody` / `createCommonBodyPull` / `createCommonBodyRequestStateChange`) over open-coding `<Body>` in the service. The existing custom-method services (`alarmclock.AddAlarm`, `cim/boot.SetBootConfigRole`, `remoteaccess.AddRemoteAccessPolicyRule`, etc.) inline `<Body>` with a `strings.Builder` — that's the legacy pattern; new code shouldn't copy it (see Implementation guidelines below).

### XML assembly (`internal/message/`)

The envelope shape lives in `internal/message/wsman.go`:

- `WSManMessageCreator` holds `MessageID` (auto-incremented per `CreateHeader` call), the `<?xml…?><Envelope …>` prefix, the closing `</Envelope>`, the anonymous reply-to address, the `PT60S` default operation timeout, and the resource URI base for the namespace (one of `AMTSchema`/`CIMSchema`/`IPSSchema`). `CreateHeader` emits `<a:Action>`, `<a:To>`, `<w:ResourceURI>`, `<a:MessageID>`, `<a:ReplyTo>`, `<w:OperationTimeout>`, and an optional `SelectorSet`. `CreateXML(header, body)` concatenates prefix + header + body + suffix.
- `message.Base` (`internal/message/base.go`) wraps `WSManMessageCreator` with `ClassName` and a `client.WSMan` and exposes `Get(selector)`, `Enumerate()`, `Pull(ctx)`, `Put(request, hasSelector, selectorSet)`, `Delete(selectorSet)`, `RequestStateChange(...)`. The `base.WSManService[T]` wrapper above defers to these.
- `OBJtoXML` + reflection-based namespacing handles model→XML conversion for `Put`. The package is `internal/` for a reason: the API is messy (reflection, ad-hoc enums for namespace prefixes) and we explicitly do not want consumers depending on it. If you need to expose helpers, expose them through `pkg/wsman/base` instead.

### Transports (`pkg/wsman/client/`)

`client.WSMan` is the interface every service talks to:

```go
type WSMan interface {
    Post(msg string) ([]byte, error)       // HTTP/HTTPS Digest path
    Connect() error                         // TCP / redirection / CIRA path
    Send([]byte) error
    Receive() ([]byte, error)
    CloseConnection() error
    IsAuthenticated() bool
    GetServerCertificate() (*tls.Certificate, error)
}
```

Three concrete implementations:

- `Target` (`wsman.go`) — HTTP(S) Digest client, used for normal WSMAN-over-HTTP traffic. Constructed via `NewWsman(Parameters)`. Honours `UseTLS`, `SelfSignedAllowed`, `PinnedCert`, `UseDigest`. Logs raw XML on stdout when `LogAMTMessages: true`.
- `WsmanTCP` (`wsman_tcp.go`) — TCP transport used for KVM / serial-over-LAN / storage-redirection. Constructed via `NewWsmanTCP(Parameters)` when `Parameters.IsRedirection` is set. `NewMessages` in `pkg/wsman/messages.go` picks between the two automatically.
- `CIRARedirectionTarget` (`cira_redirection.go`) — wraps a `CIRAChannelManager` to ship raw redirection bytes through an existing CIRA APF tunnel. Constructed via the top-level `wsman.NewCIRARedirectionMessages(manager)`; this path **does not** instantiate the AMT/CIM/IPS message builders, because redirection is a binary stream, not a WSMAN message bus.

### Test harness (`pkg/wsman/wsmantesting/`)

- `MockClient` implements `client.WSMan` by reading the response XML from disk: `../../wsmantesting/responses/<PackageUnderTest>/<lowercase-CurrentMessage>.xml`. The relative path means tests must run from inside the service's own directory (the standard `go test ./...` flow does this).
- `ExpectedResponse(messageID, resourceURIBase, method, action, address, body)` builds the **expected** XML envelope using the same template the production `WSManMessageCreator` produces. Tests `assert.Equal(t, expectedXMLInput, response.XMLInput)` — i.e. they verify byte-for-byte that the production code emits the envelope the test believes it should. This is the wire-format contract gate.
- Recorded response fixtures live under `pkg/wsman/wsmantesting/responses/{amt,cim,ips}/`. Add a fixture whenever you add a new method — the file name is `<lowercase-method>.xml` (e.g. `get.xml`, `pull.xml`, `addalarm.xml`).

### Why byte-exact XML assertions are deliberate

AMT firmware parses envelopes with a fixed schema; whitespace, namespace prefix, and attribute order all matter, and a "harmless" tidy-up of the envelope can break every downstream device call. The test suite asserts on the **exact** XML byte string emitted by the service — this is intentional, mirrors what the TypeScript `wsman-messages` library does on the JS side, and is what keeps the two libraries' wire output compatible. Do not refactor tests into "build the expected XML using the same helper the production code uses" — that would let a bug in `WSManMessageCreator` go undetected because both sides would change together.

### Authoring a new service / message

1. Create `pkg/wsman/<ns>/<class>/` with `decoder.go` (constants and return-value maps), `decoder_test.go` (lookup tests), `types.go` (response/body shapes), `marshal.go` (`.JSON()`/`.YAML()` boilerplate copied from a neighbour), and `service.go`.
2. In `service.go`, declare `type Service struct { base.WSManService[Response] }` and `NewServiceWithClient(creator, client)` that calls `base.NewService[Response](creator, "<NamespacePrefix>_<ClassName>", client)`. For Get / Enumerate / Pull / Put you're done — `WSManService` provides them.
3. For Get / Enumerate / Pull / Put / Delete / RequestStateChange, use the helpers on `base.WSManService[T]` (or, for shapes that base doesn't expose, `s.Base` from `internal/message`'s `Base`). **For custom `*_INPUT` methods, add a new helper to `internal/message/` (or `pkg/wsman/base/`) and call it from the service** — do not open-code `<Body>...</Body>` in `service.go`. The handful of services that currently do (`alarmclock.AddAlarm`, `cim/boot.SetBootConfigRole`, `remoteaccess.AddRemoteAccessPolicyRule`, etc.) are legacy from before the generic helpers existed; they're tech debt, not the pattern for new code. Never inline `<Envelope>` or `<Header>` literals either.
4. Add the service to the namespace `messages.go` `Messages` struct **and** its `NewMessages` constructor. Both must list the service or consumers can't reach it.
5. Add fixtures: `pkg/wsman/wsmantesting/responses/<ns>/<lowercase-method>.xml` for every method the test exercises (Get, Enumerate, Pull, and each custom method).
6. Add `service_test.go` with a `TestPositive<Class>` table-driven test. The standard pattern is at the top of `pkg/wsman/amt/alarmclock/service_test.go` — copy it. Each case captures `response.XMLInput` and asserts `expectedXMLInput == response.XMLInput` and `expectedResponse == response.Body`. Hit ≥85% line coverage (target 100%), including the error/negative paths — add a corresponding `TestNegative<Class>` if the service has branches the positive table can't reach.

## Consumer contract (CRITICAL)

This library is published as a Go module and consumed by first-party Go projects in the device-management-toolkit — [`rpc-go`](https://github.com/device-management-toolkit/rpc-go) (the in-host RPC client running on managed devices) and [`console`](https://github.com/device-management-toolkit/console) (the Go/Wails desktop console). Implications:

- **Every emitted XML envelope is a public wire-format contract.** Whitespace, attribute order, or namespace-prefix changes are behaviour changes for AMT firmware on the wire. If a `service_test.go` expectation needs updating because the envelope changed, treat the change as breaking unless you can show the new envelope is byte-equivalent for the firmware (it usually isn't).
- **Exported Go shapes are also public API.** Renaming a method on `wsman.Messages.*`, changing the order/optionality of a struct field, restructuring `Response`/`Body`, or moving a package (consumers import e.g. `pkg/wsman/amt/general`, `pkg/wsman/client`) is a breaking change.
- **Prefer additive evolution.** New methods on a service, new fields appended to a response struct, new packages — fine. Renaming, removing, or tightening existing shapes — breaking, and must be flagged with a `BREAKING CHANGE:` footer so semantic-release cuts a major.
- **The module is `…/v2`.** A major bump requires moving to `/v3` with the corresponding directory rename. Don't take this on without explicit user direction.

## Implementation guidelines (non-negotiable)

- **Never inline `<Envelope>`, `<Header>`, or `<Body>` literals in new service code.** `<Envelope>` lives in `WSManMessageCreator.XMLCommonPrefix/End`, `<Header>` is built by `WSManMessageCreator.CreateHeader`, and `<Body>` is owned by `internal/message/` — `DeleteBody` / `EnumerateBody` / `GetBody` constants plus the `CreateBody` / `createCommonBodyPull` / `createCommonBodyRequestStateChange` helpers, surfaced through `base.WSManService[T]`. Standard operations (Get / Enumerate / Pull / Put / Delete / RequestStateChange) are already covered. For a **custom** `*_INPUT` method where no helper exists, add a new helper to `internal/message/` (or `pkg/wsman/base/`) and call it from the service; do not open-code `<Body>...</Body>` in the service file. The eight files that currently inline `<Body>` (`pkg/wsman/amt/{alarmclock,boot/settingdata,remoteaccess,tls/credentialcontext}/…`, `pkg/wsman/cim/{boot/service,boot/configsetting,power/managementservice}/…`, `pkg/wsman/ips/power/managementservice`) are legacy that predates the generic helpers — treat them as tech debt to migrate, not patterns to copy.
- **Return `error`, never `panic`, on AMT or transport failures.** AMT faults come back as `amterror.AMTError`-typed values; wrap with `fmt.Errorf("...: %w", err)` so `errors.Is` / `errors.As` work end-to-end. The `errorlint` linter enforces this.
- **Don't add `init()` functions or package-level mutable state.** `gochecknoinits` and `goconst` are enabled. If you need configurable behaviour, expose it through a struct field on the service or a `client.Parameters` field, not a global.
- **Keep `pkg/wsman/amt`, `pkg/wsman/cim`, `pkg/wsman/ips` parallel.** Same per-class file layout (`decoder.go`/`decoder_test.go`/`types.go`/`marshal.go`/`service.go`/`service_test.go`), same constructor naming (`New<Class>WithClient`), same wiring pattern in `messages.go`. Consumers and test fixtures both depend on this regularity — diverging in one namespace breaks the muscle memory and forces them to learn three APIs.
- **Preserve doc comments on enum-style types.** `type ReturnValue int` and the string-map lookups in each `decoder.go` are the only place firmware-side integer semantics are documented. Don't drop or paraphrase them.
- **Honor `internal/message/` boundaries.** That package is internal on purpose; if a service needs a helper, add it to `pkg/wsman/base` rather than re-exporting from `internal/`.
- **Keep PRs small and scoped to one concern.** Don't bundle a new IPS message with a refactor to `client/` and a `gofumpt` sweep — they belong in separate PRs so the diff that changes wire format can be reviewed in isolation. This library ships to fleets of in-field devices via `rpc-go`; reviewers need to see envelope-affecting changes clearly.
- **Order PRs around the semver release impact.** Releases are automated from conventional commits by semantic-release. `feat:` cuts a **minor**, `fix:` / `perf:` cuts a **patch**, `BREAKING CHANGE:` cuts a **major**, `chore:` cuts a **patch** (per `.releaserc.json`'s extra rule), `refactor:` / `docs:` / `test:` / `style:` / `build:` / `ci:` do **not** cut a release. Land internal helpers and test scaffolding under `refactor:` / `test:` first; ship the user-visible new message as the `feat:` that triggers the release.
- **Before declaring work done, both must be green:** `go test ./...` and `golangci-lint run --config=./.golangci.yml ./...` (which also enforces `gofmt`/`gofumpt`/`goimports`/`gci` via the `formatters` section). CI runs the same set; failures there will block the PR.

## Commit conventions (see CONTRIBUTING.md)

Format: `<type>(<scope>): <subject>` with body and optional footer. Types: `feat | fix | docs | style | refactor | perf | test | build | ci | chore | revert`. The conventional scopes used in this repo (matching the historical `git log` — not currently enforced by `.github/commitlint.config.cjs`, which only sets `body-max-line-length` and `subject-case`): `amt`, `cim`, `ips`, `apf`, `client`, `deps`, `deps-dev`, `gh-actions`. Sub-scopes are permitted (e.g. `feat(amt/boot): …`). Reviewers may request you align with this list during review. The lint config caps `body-max-line-length` at 200; `CONTRIBUTING.md` asks contributors to keep all lines around 72 chars for readability in git tooling. The footer must reference a GitHub issue with a closing keyword (`Closes: #1234` / `Fixes: #1234` / `Resolves: #1234`).
