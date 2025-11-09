# Testing Strategy

## Pyramid

1. **Unit Tests (55%)** – Table-driven Go tests for domain services (ledger postings, reconciliation rules, payment workflows)
2. **Integration Tests (25%)** – Testcontainers for Postgres/Redis/NATS verifying repository + event pipeline
3. **Contract Tests (10%)** – Pact/Connect conformance with consuming services (food delivery, notifications)
4. **Performance Tests (5%)** – k6 and vegeta for payment/settlement latency targets
5. **Chaos/Failover (5%)** – Simulate provider outages, transaction rollback behaviours

## Tooling

- `go test ./...` with `-race` for concurrency issues
- `github.com/stretchr/testify` for assertions, `testify/suite` for complex contexts
- `github.com/testcontainers/testcontainers-go` for dependency orchestration (planned)
- `buf` for API contract conformance (once ConnectRPC integrated)
- `gosec` for security static analysis, `golangci-lint` for lint aggregator

## Coverage Targets

- Statement coverage ≥ 85%
- Ledger module ≥ 90%
- Payments/settlements invariants verified by property-based tests (rapidcheck/quicktest)

## CI Gates

- Lint + test (with coverage) required before merge
- Automatic migration check to ensure schema diff applied
- Smoke tests executed post-deploy verifying ledger posting -> reconciliation pipeline

## Future Enhancements

- Golden files for ledger journal templates
- Deterministic simulation harness for multi-currency settlements
- Synthetic monitoring for dunning workflows via notifications integration
