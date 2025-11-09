# Ledger Principles

## Core Concepts

- **Double-entry accounting**: every transaction posts balanced debits and credits across ledger accounts.
- **Chart of accounts**: hierarchical structure (Assets, Liabilities, Equity, Income, Expenses) scoped per tenant with optional shared platform accounts.
- **Journals**: immutable records capturing postings, metadata, FX rates, correlation IDs.
- **Currencies**: monetary values stored as atomic units using `shopspring/decimal`; FX conversions captured via ledger rate tables.
- **Ledger periods**: monthly closures with status workflow (open → pending → closed) to support audit controls.

## Posting Rules

1. Validate tenant permissions and account states before posting.
2. Ensure balance between debit and credit legs (sum(debits) == sum(credits)).
3. Apply idempotency keys to prevent duplicate postings.
4. Use outbox events to emit domain notifications (`ledger.journal.posted`) after commit.
5. Support multi-currency by tracking base + reporting currency conversions.

## Account Lifecycle

- Accounts created via provisioning API with default state `pending`.
- Activation requires approval workflow; accounts can be soft-deleted but journal history is immutable.
- Accounts may be tagged with compliance metadata (e.g., restricted funds, escrow).

## Compliance & Audit

- Immutable event log stored in Postgres + optional blockchain anchoring for high-value settlements.
- All adjustments require supervisor approval and audit trail (who, when, why).
- Maintain audit exports (CSV/Parquet) with checksum verification.

## Reconciliation

- Automated ingestion of statements via `settlements` module.
- Suspense accounts used for unmatched transactions with SLA-driven workflows.
- Exception queues exposed via API/UI for manual reconciliation.

## Future Enhancements

- Smart contracts integration for on-chain settlement proofs.
- Automated FX revaluation entries at period close.
- Machine learning scoring to flag anomalous transactions.
