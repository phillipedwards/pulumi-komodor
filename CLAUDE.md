# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Is

A Pulumi Terraform Bridge Provider that wraps the [Komodor Terraform provider](https://github.com/komodorio/terraform-provider-komodor), exposing its resources and data sources through Pulumi's multi-language SDK framework (Go, Python, Node.js, .NET).

The repo was bootstrapped from the `pulumi/pulumi-terraform-bridge` boilerplate and has been wired to the real Komodor Terraform provider (`github.com/komodorio/terraform-provider-komodor`).

## Commands

```bash
make provider            # Build just the provider binary
make build               # Build provider + all SDKs and install for testing
make generate            # Re-generate schema and SDK source from upstream TF provider
make lint_provider       # Run golangci-lint
make lint_provider.fix   # Auto-fix lint issues
make test_provider       # Run provider unit tests
make test                # Run integration/example tests (requires `make build` first)
```

Language-specific SDK targets: `generate_[dotnet|go|nodejs|python]`, `build_[dotnet|go|nodejs|python]`.

Tools are managed via `mise` — run `mise install` before anything if tools are missing.

## Architecture

### Build pipeline

1. `pulumi-tfgen-xyz` (code generator) — reads the upstream Terraform provider's schema and emits `schema.json` plus SDK source for each target language.
2. `pulumi-resource-xyz` (plugin binary) — the runtime provider; embeds `schema.json` and delegates all resource operations to the upstream TF provider via the bridge shim.
3. Multi-language SDKs in `sdk/` are generated output; edit the generator config in `provider/resources.go`, not the SDK files directly.

### Key file

`provider/resources.go` is the single source of truth for everything the Pulumi provider exposes. It returns a `tfbridge.ProviderInfo` that controls:
- Which upstream TF provider is wrapped (`P: shimv2.NewProvider(...)`)
- Resource/datasource token naming (`MustComputeTokens`)
- Per-language SDK options (import paths, package names, PyProject, etc.)
- Provider-level configuration schema and any extra types
- Auto-aliasing and autonaming settings

If you add, rename, or configure resources, data sources, or provider config, that's the only file to touch.

### Terraform framework shim

The upstream Komodor provider uses `terraform-plugin-sdk/v2`, so this bridge uses `shimv2`. If the upstream ever migrates to `terraform-plugin-framework`, swap `shimv2.NewProvider` for `pfbridge.ShimProvider` and update the `main.go` entry points accordingly (instructions are in comments in `resources.go`).

### Version injection

`provider/pkg/version/version.go` holds an empty `Version` string that is set at link time via `-ldflags`. Never hardcode a version there.

## Go module layout

| Directory | Module |
|-----------|--------|
| `provider/` | `github.com/pulumi/pulumi-xyz/provider` |
| `sdk/` | `github.com/phillipedwards/pulumi-komodor/sdk` |

The provider and SDK are separate Go modules. Run `go` commands from within the appropriate directory.
