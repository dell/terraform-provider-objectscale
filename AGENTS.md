# AGENTS.md - Dell Terraform Provider for ObjectScale

## Project Overview

This is the Terraform provider for Dell ObjectScale S3-compatible object storage. It implements resources and data sources using HashiCorp's Terraform Plugin Framework, enabling infrastructure-as-code management of ObjectScale deployments.

- **Language:** Go 1.25
- **Module path:** `terraform-provider-objectscale`
- **Terraform Plugin Framework:** v1.15.1
- **SDK:** Internal client (no external SDK dependency)
- **Registry address:** `registry.terraform.io/dell/objectscale`
- **License:** Mozilla Public License 2.0

## Architecture

The provider follows the standard Terraform Plugin Framework architecture. It runs as a gRPC server that Terraform Core communicates with to manage ObjectScale resources.

### Provider Configuration

The provider authenticates to an ObjectScale management endpoint using endpoint, username, and password. Configuration can be supplied via HCL provider block or environment variables (`OBJECTSCALE_ENDPOINT`, `OBJECTSCALE_USERNAME`, `OBJECTSCALE_PASSWORD`, `OBJECTSCALE_INSECURE`, `OBJECTSCALE_TIMEOUT`).

### SDK Strategy

Uses an **internal client** — REST calls are implemented directly in provider code under `internal/client/` and `internal/clientgen/`. No external SDK dependency. An OpenAPI-generated client (`clientgen/`) provides API coverage. This gives full control but requires more maintenance.

### Resources and Data Sources

The provider exposes approximately 18 resources and 17 data sources covering ObjectScale entities such as accounts, buckets, users, IAM policies, object stores, and replication.

## Directory Structure

```
main.go                           Entry point (providerserver.Serve)
internal/
  provider/
    provider.go                   Provider configuration, resource/datasource registration
    *_resource.go                 Resource implementations
    *_resource_schema.go          Resource schema definitions
    *_datasource.go               Data source implementations
    *_test.go                     Unit and acceptance tests
  helper/                         Shared helper functions
  models/                         Terraform state model structs
  client/                         ObjectScale API client
  clientgen/                      OpenAPI-generated client code
clientgen_utils/                  OpenAPI spec and generation utilities
  openapi_specs/                  OpenAPI JSON specifications
examples/                         Example HCL configurations
docs/                             Generated documentation
about/                            Provider metadata
tools/                            Build and generation tools
```

## Build Commands

| Command | Description |
|---------|-------------|
| `make build` | Compile the provider binary |
| `make install` | Build and install to `~/.terraform.d/plugins/` |
| `make test` | Run unit tests |
| `make testacc` | Run acceptance tests (`TF_ACC=1`, requires live hardware) |
| `make check` | Run `gofmt`, `golangci-lint`, `go vet` |
| `make gosec` | Run security scan with `gosec` |
| `make cover` | Generate HTML coverage report |
| `make generate` | Run `go generate` (docs generation) |

## Testing

### Unit Tests (mockey)

- Test files follow `*_test.go` convention in `internal/provider/`.
- Frameworks: `github.com/stretchr/testify` (assertions), `github.com/bytedance/mockey` (function-level mocking).
- Run with `make test`.
- No hardware required.

### Acceptance Tests (terraform-plugin-testing)

- **Requires live ObjectScale hardware** with credentials set via environment variables.
- Creates real resources — clean up after failures.
- Run with `make testacc`.

### Running Tests

```bash
# Unit tests (no hardware)
make test

# Acceptance tests (requires live hardware)
export OBJECTSCALE_ENDPOINT="https://objectscale-mgmt-ip"
export OBJECTSCALE_USERNAME="admin"
export OBJECTSCALE_PASSWORD="secret"
export OBJECTSCALE_INSECURE="true"
make testacc
```

## Code Style and Conventions

### Code Organization Patterns

- **Resource pattern:** Each resource has up to three files: `<name>_resource.go`, `<name>_resource_schema.go`, plus helpers.
- **Models:** Terraform state structs in `internal/models/` using `tfsdk` struct tags.
- **Internal client:** `internal/client/` wraps REST API calls; `internal/clientgen/` provides OpenAPI-generated methods.
- All source lives under `internal/` to prevent external import.

### File Header

All source files must include the Dell copyright and MPL 2.0 license header.

## Common Development Tasks

### Adding a New Resource

1. Create resource, schema, and model files under `internal/` following existing patterns.
2. Add helper functions for API-to-Terraform mapping.
3. Register in `internal/provider/provider.go`.
4. Add unit and acceptance tests.
5. Create example HCL in `examples/resources/objectscale_<name>/`.
6. Run `make generate` to produce documentation.

### Regenerating the OpenAPI Client

Use the OpenAPI generator with specs in `clientgen_utils/openapi_specs/`.

## CI/CD

GitHub Actions workflows in `.github/workflows/`. GoReleaser configuration in `.goreleaser.yaml` builds cross-platform binaries.

## Code Ownership

All files are owned by the maintainers defined in `.github/CODEOWNERS`.
