# API & Event Contracts

## REST

- Base path: `/v1/{tenantId}` where `{tenantId}` maps to organisation/branch context.
- Use resource-oriented endpoints: `/ledger/accounts`, `/payments/intents`, `/settlements/batches`.
- Timestamp fields ISO8601 UTC; monetary values expressed as string decimals with currency code.
- Error responses follow RFC 7807 Problem Details with correlation/request IDs.
- OpenAPI specs generated via `buf`/`oapi-codegen` (specs stored in `api/openapi/`).

## Webhooks

- Signed using HMAC-SHA256 with rotating secrets stored in Vault.
- Event types: `payment.intent.created`, `payment.intent.completed`, `settlement.batch.settled`, `ledger.journal.posted`, `invoice.due`, `invoice.paid`.
- Retries use exponential backoff with dead-letter queue (NATS JetStream).

## Events (NATS JetStream)

- Subject pattern: `treasury.{domain}.{event}` e.g. `treasury.payments.intent.created`.
- Payload envelope aligns with CloudEvents 1.0 for interoperability.
- Consumers require durable subscriptions with manual acks to support replay.

## gRPC / ConnectRPC (Planned)

- Services: `LedgerService`, `PaymentService`, `SettlementService` under package `treasury.v1`.
- Use Buf for linting (`buf.yaml`) and codegen to `internal/gen/`.
- Auth handled via mTLS + JWT service accounts, enforced via interceptors.

## Versioning

- SemVer across APIs; non-breaking additions allowed within `v1`.
- Breaking changes require new major version path and migration plan.
- Deprecation notices documented in `docs/api-contracts.md` and communicated to consumers.
