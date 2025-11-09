# Development Workflow

## Prerequisites

- Go 1.22+
- PostgreSQL 14+, Redis 7+, NATS JetStream
- golangci-lint, buf CLI (for ConnectRPC), Docker (optional for dependencies)

## Setup

```bash
cp config/app.env.example .env
make tidy
make run
```

Use Docker Compose (to be added) for local Postgres/Redis/NATS if needed. Seed scripts and migrations will live under `db/migrations`.

## Branching & Commits

- Trunk-based development with short-lived branches (`feature/`, `fix/`, `chore/`)
- Conventional Commits powering semantic releases (e.g. `feat(ledger): support multi-currency journals`)
- Pull requests must:
  - Pass `make lint` and `make test`
  - Include migration scripts (if schema changes)
  - Update documentation (architecture, ledger principles) where relevant
  - Outline rollout/rollback, feature flags, data migrations

## CI/CD Pipeline

1. **Lint/Test** – golangci-lint, go test with race + coverage
2. **Build** – multi-stage Docker build, push to registry
3. **Security** – gosec, Trivy scans, dependency audit
4. **Deploy** – ArgoCD sync to staging; manual approval for production
5. **Post-Deploy** – smoke tests, ledger reconciliation sanity checks

## Local Tooling

- Hot reload via [air](https://github.com/cosmtrek/air) (optional) with config in `.air.toml`
- `buf` for proto linting and codegen (proto config incoming)
- `Taskfile` (optional) for higher-level automation wrappers

## Observability in Dev

- Run `docker compose -f ops/observability.yml up` to spin up Prometheus + Grafana (to be added)
- Use `OTEL_EXPORTER_OTLP_ENDPOINT` to target local collector
- Inspect event streams via `nats cli` (subscribe to `treasury.*`)
