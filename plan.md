# Treasury Service Delivery Plan

## Vision & Scope
- Build a world-class, API-first treasury and payments orchestration microservice supporting all BengoBox products with resilient collections, disbursements, settlements, reconciliation, and billing flows across multiple organisations and branches.
- Provide first-class integrations for MPesa (Daraja C2B, B2C, B2B, STK Push/MPesa Express), PayPal, Stripe, blockchain settlement rails, and pluggable bank APIs while maintaining strong compliance, auditability, and multi-currency ledgering.

## Technical Foundations
- **Language & Framework:** Go 1.22+, Clean Architecture with domain-driven design, gRPC + REST gateway (ConnectRPC or gRPC-Gateway), Cobra CLI for ops tasks.
- **Core Services:** Modular domains (`tenancy`, `ledger`, `payments`, `settlements`, `reconciliation`, `billing`, `kyc`, `integrations`, `notifications` bridge).
- **Data Layer:** PostgreSQL for double-entry ledger + transactional data, EventStore (or Kafka) for event sourcing, Redis for caching idempotency keys and rate limits, S3-compatible storage for statements and invoice artifacts, optional blockchain state anchored via Hyperledger Fabric or Polygon PoS for immutable settlement proofs.
- **Infrastructure:** Containerized deployment, Helm charts, GitHub Actions CI/CD, Terraform-managed cloud resources, HashiCorp Vault for secrets, OpenTelemetry instrumentation, Prometheus + Grafana dashboards.
- **Security:** OAuth2 service accounts, mTLS between services, HSM/KMS for signing, encrypted config, comprehensive RBAC, tenant-aware access control.

## Domain Capabilities
1. **Multi-Organisation & Branch Tenancy (Priority 1)**
   - Hierarchical tenant model (organisation → branch → application) with strict data isolation, configurable branding, and onboarding APIs.
   - Self-service provisioning for consuming apps to link existing organisations or create new ones via invite/approval workflows.
2. **Core Ledger & Accounts (Priority 1)**
   - Multi-tenant chart of accounts, currency-aware double-entry ledger, hierarchical accounts (platform, merchant, rider, customer).
   - Real-time balances, ledger snapshots, period closures, FX rate integration.
3. **Collections (Priority 1)**
   - MPesa C2B registration, webhook processing, STK Push workflows, duplicate payment detection.
   - Stripe & PayPal checkout sessions, card tokenization (PCI scope minimization), 3DS and SCA support.
   - Payment intent API with status lifecycle, retry & fallback logic.
4. **Blockchain Settlement Rail (Priority 2)**
   - Optional blockchain ledger anchoring for high-value settlements (Hyperledger Fabric channels or Polygon smart contracts).
   - Smart-contract based escrow, cryptographic proof APIs, chain event listeners for reconciliation.
5. **Invoicing & Billing (Priority 2)**
   - Supports full invoice matrix: standard invoice, quote, proforma, tax invoice, credit memo, credit note, receipt, sales receipt, cash receipt, estimate, purchase order, delivery note, debit note.
   - Template customization, numbering rules, multi-currency tax handling, attachment management, payment links per invoice.
   - Workflow automation (draft → approval → send), recurring invoices, dunning schedules.
6. **Disbursements & Settlements (Priority 2)**
   - MPesa B2C/B2B payouts, scheduled batch settlements, bulk disbursement API, negative balance guardrails.
   - Bank payout adapters (REST/SOAP) with configurable mapping, SWIFT/ACH support roadmap.
7. **Reconciliation & Reporting (Priority 2)**
   - Automated statement ingestion (MPesa, Stripe, PayPal, banks) with matching engine, suspense account handling.
   - Dispute management, refund flows, chargeback workflows, audit trails.
8. **Compliance & Risk (Priority 3)**
   - KYC/KYB onboarding, sanctions screening (OFAC, EU), transaction monitoring rules, AML suspicious activity reports.
   - Limits & velocity checks, fraud scoring hooks, anomaly alerts.
9. **Treasury Operations Portal (Priority 3)**
   - User roles for finance team, manual adjustments with approval workflow, settlement calendar, cash position dashboard.
10. **Developer Experience & Ecosystem (Priority 4)**
   - SDKs (Go/TypeScript), webhook sandbox, Postman collections, observability dashboards, feature toggles.
11. **Notifications Orchestration (Priority 4)**
   - Deep integration with `notifications-app` to deliver invoice emails/SMS with secure pay links, payment reminders, dunning sequences, and provider redundancy.

## Integration Strategy
- **API-First Design:** REST and gRPC endpoints exposing multi-tenant aware operations (`/v1/{tenantId}/payments`, `/v1/{tenantId}/invoices`) with service discovery and API key/ OAuth2 flows for partner apps.
- **External APIs:** MPesa Daraja (OAuth token mgmt, short codes, callback URLs), Stripe Connect, PayPal REST, blockchain nodes (Hyperledger Fabric SDK / Polygon RPC), custom bank connectors via interface adapters.
- **Internal Consumers:** Food Delivery backend (orders/payments), ERP, notifications (status events & invoice delivery), billing systems, marketplace apps.
- **Notifications Integration:** Webhook and event topics (`invoice_created`, `invoice_due`, `payment_failed`, `payment_success`) consumed by `notifications-app` to trigger templated emails/SMS with secure payment links and attachments.
- **Onboarding Flows:** Tenancy APIs enable external apps to create/link organisations and branches, manage API credentials, and sync metadata.
- **Message Contracts:** Payment and invoicing events (`payment_initiated`, `payment_captured`, `invoice_sent`, `payout_requested`, `settlement_completed`) published to event bus.
- **Idempotency:** Idempotency keys per request, outbox pattern for reliable event delivery, blockchain transaction hash correlation.

## Non-Functional Goals
- Availability 99.95%, financial accuracy with ACID guarantees, P99 latency < 800ms for payment initiation.
- Regulatory compliance (PCI DSS scope reduction, Kenya CBK guidelines, GDPR/DPA for personal data), audit-ready logs, immutability guarantees via blockchain anchoring.
- Disaster recovery (RPO < 5 minutes, RTO < 1 hour), active monitoring & alerting, chaos testing for provider outages.
- Tenant isolation verified through automated security tests, configurable data residency by organisation, and branch-level segregation policies.

## Roadmap & Sprint Plan (Priority Order)
1. **Sprint 0 – Foundations & Compliance Readiness (Week 1)**
   - Repository, CI/CD, multi-tenant architecture blueprint, environment configs, secret management, audit logging baseline, documentation scaffolding.
2. **Sprint 1 – Tenancy & Ledger Core (Weeks 2-3)**
   - Organisation/branch onboarding APIs, tenant RBAC, double-entry ledger engine, balance queries, FX rate service stubs, initial migrations & tests.
3. **Sprint 2 – MPesa Collections Core (Weeks 4-5)**
   - Daraja authentication, C2B registration, STK Push flow, callback handling, idempotent payment intents, error taxonomy.
4. **Sprint 3 – Blockchain Settlement Rail (Weeks 6-7)**
   - Deploy baseline Fabric channel or Polygon contracts, chain anchoring service, settlement proof APIs, reconciliation hooks.
5. **Sprint 4 – Card & Wallet Collections (Weeks 8-9)**
   - Stripe Checkout integration, PayPal orders API, tokenized payments, webhooks, multi-currency support.
6. **Sprint 5 – Invoicing & Document Engine (Weeks 10-11)**
   - Invoice lifecycle management, template designer, numbering rules, PDF generation, payment link issuance, multi-document support (quote, PO, delivery note, credit memo, etc.).
7. **Sprint 6 – Notifications & Dunning Automation (Weeks 12-13)**
   - Event contracts with `notifications-app`, email/SMS invoice delivery, reminder schedules, portal links, dunning workflows, customer communication preferences.
8. **Sprint 7 – Disbursements Engine (Weeks 14-15)**
   - MPesa B2C/B2B payouts, payout scheduling, approval workflow, treasury-held float management, branch-level settlement rules.
9. **Sprint 8 – Reconciliation & Reporting (Weeks 16-17)**
   - Statement ingestion, auto-matching, suspense queues, manual reconcile UI/API, financial reporting endpoints, blockchain verification.
10. **Sprint 9 – Compliance & Risk Controls (Weeks 18-19)**
   - KYC/KYB module, AML rules, sanctions screening integration, fraud monitoring hooks, audit exports.
11. **Sprint 10 – Bank Integrations & Treasury Portal (Weeks 20-21)**
   - Generic bank API adapter framework, first bank integration, treasury operations dashboard, approval workflows, branch cash positioning.
12. **Sprint 11 – Hardening, Performance & Launch (Weeks 22-23)**
   - Load testing, failover drills, penetration testing, documentation, production go-live checklist, blockchain/node resilience testing.

## Backlog & Future Enhancements
- Instant payout cards, virtual IBAN issuance, smart-contract escrow market, cash-flow forecasting, dynamic currency conversion, crypto on/off ramps, machine learning risk scoring, ISO20022 messaging support, embedded finance marketplace for partner banks.

## Deployment

- Containerised (Docker multi-stage) → Helm chart → ArgoCD GitOps
- Horizontal Pod Autoscaling (CPU/RPS), Pod Disruption Budgets
- Feature flags handled via treasury-managed config service (future)

## Runtime Ports
- **Local development:** HTTP API listens on **4001** to avoid conflicts with the Food Delivery backend (4000) and Notifications service (4002).
- **Cloud deployment:** `TREASURY_HTTP_PORT` is overridden to **4000** so all backend services share a common service port behind ingress controllers.

---
**Next Steps:** Confirm regulatory requirements with compliance counsel, finalize MPesa credentials, align on shared event schemas with dependent services (including notifications), select blockchain rail, and schedule joint tenant-onboarding and invoicing integration tests with partner applications.

