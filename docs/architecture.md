# Architecture Overview

## Layers

- **Transport** (`internal/http`, future `internal/grpc`): chi-based REST API with tenant-aware routing, ConnectRPC gateway planned for high-throughput internal consumers.
- **Application** (`internal/app`): Bootstraps configuration, logging, infrastructure adapters, lifecycle management for HTTP/grpc servers and workers.
- **Domain Modules** (`internal/modules`):
  - `ledger`: double-entry posting engine, accounts, journals, FX conversions
  - `payments`: payment intents, MPesa/Stripe connectors, idempotency safeguards
  - `settlements`: disbursement scheduling, reconciliation, treasury float management
- **Infrastructure** (`internal/platform`): PostgreSQL (pgx), Redis, NATS JetStream, S3-compatible storage, Vault-based secrets provider, telemetry clients.
- **Shared** (`internal/shared`): Logging, middleware, error handling, cross-cutting utilities.

## Data Flow

1. API request hits REST endpoint with tenant context (header or URL param).
2. Handler validates payload, enriches context (trace, request ID, tenant), invokes use case service.
3. Domain service orchestrates repositories (Postgres, Redis), invokes integrations (treasury providers, notifications).
4. Events emitted via NATS JetStream with outbox pattern to ensure reliability.
5. Reconciliation jobs and workers consume streams to update ledgers, notify services, issue payouts.

## Key Integrations

- **Food Delivery backend** – order/payment events, wallet balances, loyalty accruals.
- **Notifications app** – invoice lifecycle events, dunning campaigns, payment confirmations.
- **External Providers** – MPesa Daraja, Stripe Connect, PayPal, bank APIs via pluggable adapters.
- **Compliance Tooling** – AML/KYC services, sanction screening, regulatory reporting.

## Deployment

- Container images built via multi-stage Dockerfile, Helm chart per environment.
- GitOps with ArgoCD; secrets delivered via Vault/Kubernetes secrets.
- Horizontal scaling with HPA, Pod Disruption Budgets, zero-downtime rollout with canary/circuit breakers.

## Observability

- Structured JSON logging with request/tenant IDs.
- Prometheus metrics (`/metrics`) capturing request latency, ledger posting throughput, event lag.
- OTEL exporters for traces and metrics; spans propagate across services with W3C trace context.
- SLO dashboards (Grafana) & alerting (PagerDuty) for collections latency, settlement success, reconciliation backlog.
