# Treasury Service

Go-powered treasury and payments orchestration service supporting collections, disbursements, settlements, reconciliation, invoicing, and compliance for all BengoBox products.

## Highlights

- **Clean architecture** with explicit domain modules (`ledger`, `payments`, `settlements`)
- **API contracts** delivered via REST (chi) with future ConnectRPC/gRPC gateway
- **Data layer**: PostgreSQL (double-entry ledger), Redis (idempotency + caching), NATS JetStream event bus, S3-compatible object storage
- **Observability**: zap structured logging, Prometheus metrics, OTEL-ready instrumentation
- **Security**: Tenant-aware RBAC, secrets provider abstraction, Vault-ready integration, request ID + tenant middleware

## Getting Started

```bash
cp config/app.env.example .env
make tidy
make run
```

Port mapping:

- Local development defaults to **http://localhost:4001**.
- In Kubernetes environments the chart sets `TREASURY_HTTP_PORT=4000`, aligning with the shared ingress port convention used by other backend services.

### Environment Variables

All configuration keys are prefixed with `TREASURY_`. See [`config/app.env.example`](config/app.env.example) for defaults covering HTTP/gRPC ports, Postgres, Redis, NATS, storage, and telemetry endpoints.

### Make Targets

| Command | Description |
| ------- | ----------- |
| `make run` | Run the HTTP API (`cmd/api`) |
| `make worker` | Run worker binary (JetStream consumers, scheduled jobs) |
| `make build` | Build API and worker binaries into `bin/` |
| `make test` | Execute Go test suites |
| `make lint` | Run golangci-lint (install separately) |
| `make proto` | Generate protobuf/Connect stubs via Buf (config forthcoming) |

## Project Layout

```
cmd/
  api/         # HTTP/Connect service entry point
  worker/      # Background job processor (JetStream, schedulers)
internal/
  app/         # Application bootstrap and lifecycle management
  config/      # Environment-backed configuration loader
  http/        # Routes, handlers, DTOs
  modules/     # Ledger, payments, settlements domain modules
  platform/    # Database, cache, events, storage, secrets adapters
  shared/      # Logger and HTTP middleware utilities
```

## Documentation

Documentation lives under `docs/` and is indexed in [`docs/documentation-guide.md`](docs/documentation-guide.md):

- [`docs/architecture.md`](docs/architecture.md) – clean architecture layout, module boundaries, integration points
- [`docs/development-workflow.md`](docs/development-workflow.md) – local setup, CI/CD, branching strategy
- [`docs/testing-strategy.md`](docs/testing-strategy.md) – testing pyramid, tooling, coverage goals
- [`docs/ledger-principles.md`](docs/ledger-principles.md) – double-entry ledger rules and invariants
- [`docs/api-contracts.md`](docs/api-contracts.md) – REST/Webhook/Events contract guidelines

## Community Files

- [`CONTRIBUTING.md`](CONTRIBUTING.md)
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md)
- [`SECURITY.md`](SECURITY.md)
- [`SUPPORT.md`](SUPPORT.md)
- [`CHANGELOG.md`](CHANGELOG.md)

## Roadmap Next Steps

- Implement domain services for ledger postings, reconciliation jobs, and payments orchestration
- Add ConnectRPC/gRPC gateway (`cmd/api/grpc`) and Buf configuration
- Wire Vault secrets provider and MinIO/S3 artifact storage
- Integrate treasury events with notifications app (dunning, invoices) and food-delivery backend
- Expand testing harness with Testcontainers and k6 performance scenarios
