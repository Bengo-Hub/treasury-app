# Contributing

Thank you for investing in the Treasury service. This guide helps you contribute safely and effectively.

## Prerequisites

- Go 1.22+
- golangci-lint, buf CLI
- Access to shared Vault secrets (for integration testing)
- Familiarity with double-entry accounting concepts

## Workflow

1. Branch from `main`: `git checkout -b feature/{ticket}`
2. Sync dependencies: `make tidy`
3. Implement changes + tests (`make test`)
4. Run linters (`make lint`) and ensure `go fmt ./...` passes
5. Update documentation (architecture, ledger principles, API contracts) when behaviour changes
6. Commit with Conventional Commits (e.g. `feat(ledger): add journal reversal support`)
7. Open a PR including:
   - Summary + business context
   - Testing evidence (unit/integration/perf)
   - Rollout/rollback strategy and feature flags
   - Migration notes (if schema/state changes)

## Standards

- Keep domain logic in `internal/modules` with interfaces for persistence/adapters
- Use context-aware functions and propagate request/tenant IDs
- Enforce idempotency for external side effects (payments, settlements)
- Structure logs with zap; include request and tenant metadata
- Add Prometheus metrics / OTEL spans for new long-running processes

## Testing Expectations

- Unit tests for new domain logic
- Integration tests when touching repositories, event pipelines, or external adapters (Testcontainers)
- Contract tests when modifying APIs/events consumed by other services
- Performance tests for high-volume flows (payments, settlements), triggered pre-release

## Documentation

- Update relevant docs under `docs/` and ensure [`docs/documentation-guide.md`](docs/documentation-guide.md) references new material
- Record release notes in [`CHANGELOG.md`](CHANGELOG.md)

By contributing you agree to the [Code of Conduct](CODE_OF_CONDUCT.md).
