## Treasury Service Delivery Plan

### 1. Vision & Scope
- Build a world-class, API-first treasury and payments orchestration microservice supporting all BengoBox products with resilient collections, disbursements, settlements, reconciliation, and billing flows across multiple organisations and branches.
- Provide first-class integrations for MPesa (Daraja C2B, B2C, B2B, STK Push/MPesa Express), PayPal, Stripe, blockchain settlement rails, and pluggable bank APIs while maintaining strong compliance, auditability, and multi-currency ledgering.
- **Entity Ownership**: This service owns all financial entities: ledger & accounts, payment intents & transactions, invoices, expenses, bills (AP), payouts, settlements, budgets, mileage logs, and fuel purchases. Cafe-backend and POS services call treasury APIs for payment processing but store only payment references (not full transactions). Logistics creates expenses/bills via treasury APIs. See `docs/cross-service-entity-ownership.md` for complete ownership matrix.

### 2. Technical Foundations
**Technologies explained (plain English):**
- Go (Golang): A compiled programming language designed for reliability and performance.
- gRPC (Google Remote Procedure Call): A high‑performance binary protocol over HTTP/2 for service‑to‑service APIs; paired with a REST gateway for browser/mobile clients.
- PostgreSQL: A relational database; supports ACID transactions needed for financial accuracy.
- Event sourcing (EventStore/Kafka): Persisting system events to enable auditability and replay.
- Redis: In‑memory datastore for fast lookups, caching, idempotency keys, and rate limiting.
- S3‑compatible storage: Object storage for invoices, statements, and exports.
- Hyperledger Fabric / Polygon PoS: Distributed ledgers used to anchor proofs of settlement where required.
- OAuth2, mTLS, KMS/HSM: Standards for identity, secure transport, and key custody.
- **Language & Framework:** Go 1.22+, Clean Architecture with domain-driven design, gRPC + REST gateway (ConnectRPC or gRPC-Gateway), Cobra CLI for ops tasks.
- **Core Services:** Modular domains (`tenancy`, `ledger`, `payments`, `settlements`, `reconciliation`, `billing`, `kyc`, `integrations`, `notifications` bridge).
- **Data Layer:** PostgreSQL for double-entry ledger + transactional data, EventStore (or Kafka) for event sourcing, Redis for caching idempotency keys and rate limits, S3-compatible storage for statements and invoice artifacts, optional blockchain state anchored via Hyperledger Fabric or Polygon PoS for immutable settlement proofs.
- **Infrastructure:** Containerized deployment, Helm charts, GitHub Actions CI/CD, Terraform-managed cloud resources, HashiCorp Vault for secrets, OpenTelemetry instrumentation, Prometheus + Grafana dashboards.
- **Security:** OAuth2 service accounts, mTLS between services, HSM/KMS for signing, encrypted config, comprehensive RBAC, tenant-aware access control.
- **Auth-Service SSO Integration:** ✅ **COMPLETED** - Integrated `shared/auth-client` v0.1.0 library for production-ready JWT validation using JWKS from auth-service. All protected `/v1/{tenantID}` routes require valid Bearer tokens. Auth config added to config struct with JWKS caching and refresh settings. Swagger documentation updated with BearerAuth security definition. **Deployment:** Uses monorepo `replace` directives with versioned dependency (`v0.1.0`). Go workspace (`go.work`) handles local development automatically. Each service has independent DevOps workflows and can be deployed separately while sharing the auth library. See `shared/auth-client/DEPLOYMENT.md` and `shared/auth-client/TAGGING.md` for details.

### 3. Domain Capabilities
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
4. **Expenses, Bills & Accounts Payable (Priority 1)**
   - Expenses module with categories, merchants, receipt OCR (future), multi-currency, tax codes, approvals, and reimbursements.
   - Vendor bills (AP) with purchase orders, partial receipts, credits/debit notes, and due date workflows.
   - Dimensions/tags: project, cost center, vehicle, route, driver (for logistics), branch; robust COA mapping rules.
5. **Mileage & Fuel Management (Priority 1)**
   - Mileage logs (per-driver/vehicle), per‑km reimbursement rates by policy; automated expense creation from logs.
   - Fuel purchases with receipt capture; price/volume validation; per‑trip fuel allocation.
   - Integrations: ingest fuel/toll/parking/maintenance events from logistics-service; auto‑categorise and post to GL.
4. **Blockchain Settlement Rail (Priority 2)**
   - Optional blockchain ledger anchoring for high-value settlements (Hyperledger Fabric channels or Polygon smart contracts).
   - Smart-contract based escrow, cryptographic proof APIs, chain event listeners for reconciliation.
6. **Invoicing & Billing (Priority 2)**
   - Supports full invoice matrix: standard invoice, quote, proforma, tax invoice, credit memo, credit note, receipt, sales receipt, cash receipt, estimate, purchase order, delivery note, debit note.
   - Template customization, numbering rules, multi-currency tax handling, attachment management, payment links per invoice.
   - Workflow automation (draft → approval → send), recurring invoices, dunning schedules.
7. **Disbursements & Settlements (Priority 2)**
   - MPesa B2C/B2B payouts, scheduled batch settlements, bulk disbursement API, negative balance guardrails.
   - Bank payout adapters (REST/SOAP) with configurable mapping, SWIFT/ACH support roadmap.
8. **Reconciliation & Reporting (Priority 2)**
   - Automated statement ingestion (MPesa, Stripe, PayPal, banks) with matching engine, suspense account handling.
   - Dispute management, refund flows, chargeback workflows, audit trails.
9. **Budgets & Cost Centers (Priority 3)**
   - Cost center and project budgets; budget vs actuals reporting; alerts on threshold breaches.
   - Allocations: split costs across projects/routes/vehicles.
10. **Compliance & Risk (Priority 3)**
   - KYC/KYB onboarding, sanctions screening (OFAC, EU), transaction monitoring rules, AML suspicious activity reports.
   - Limits & velocity checks, fraud scoring hooks, anomaly alerts.
11. **Treasury Operations Portal (Priority 3)**
   - User roles for finance team, manual adjustments with approval workflow, settlement calendar, cash position dashboard.
12. **Developer Experience & Ecosystem (Priority 4)**
   - SDKs (Go/TypeScript), webhook sandbox, Postman collections, observability dashboards, feature toggles.
13. **Notifications Orchestration (Priority 4)**
   - Deep integration with `notifications-app` to deliver invoice emails/SMS with secure pay links, payment reminders, dunning sequences, and provider redundancy.

### 4. Integration Strategy
- **API-First Design:** REST and gRPC endpoints exposing multi-tenant aware operations (`/v1/{tenantId}/payments`, `/v1/{tenantId}/invoices`) with service discovery and API key/ OAuth2 flows for partner apps.
- **External APIs:** MPesa Daraja (OAuth token mgmt, short codes, callback URLs), Stripe Connect, PayPal REST, blockchain nodes (Hyperledger Fabric SDK / Polygon RPC), custom bank connectors via interface adapters.
- **Internal Consumers:** Logistics (expenses/bills for fuel/toll/parking/maintenance; per‑km reimbursements; driver payouts), Food Delivery backend (orders/payments), ERP, notifications (status events & invoice delivery), billing systems, marketplace apps.
- **Notifications Integration:** Webhook and event topics (`invoice_created`, `invoice_due`, `payment_failed`, `payment_success`) consumed by `notifications-app` to trigger templated emails/SMS with secure payment links and attachments.
- **Onboarding Flows:** Tenancy APIs enable external apps to create/link organisations and branches, manage API credentials, and sync metadata.
- **Message Contracts:** Payment and invoicing events (`payment_initiated`, `payment_captured`, `invoice_sent`, `payout_requested`, `settlement_completed`) published to event bus.
- **Idempotency:** Idempotency keys per request, outbox pattern for reliable event delivery, blockchain transaction hash correlation.

#### 4.1 Logistics Expense Ingestion (Design)
- Endpoints for expense ingestion: `/v1/{tenantId}/expenses`, `/v1/{tenantId}/bills` (vendor), `/v1/{tenantId}/journals`.
- Required dimensions: `route_id`, `vehicle_id`, `driver_id`, `category`, `amount`, `currency`, `cost_center`, `project`, `receipt_url`.
- Category → GL mapping rules (admin-configurable) determine target account/tax code; defaults per tenant.
- Mileage API: `/v1/{tenantId}/mileage-logs` + policy engine (per‑km rate by vehicle class/region).
- Reimbursements: approved expenses flow into payouts (driver/vendor) with approval steps and notifications.

### 5. Non‑Functional Goals
- Availability 99.95%, financial accuracy with ACID guarantees, P99 latency < 800ms for payment initiation.
- Regulatory compliance (PCI DSS scope reduction, Kenya CBK guidelines, GDPR/DPA for personal data), audit-ready logs, immutability guarantees via blockchain anchoring.
- Disaster recovery (RPO < 5 minutes, RTO < 1 hour), active monitoring & alerting, chaos testing for provider outages.
- Tenant isolation verified through automated security tests, configurable data residency by organisation, and branch-level segregation policies.

### 6. Roadmap & Sprint Plan (Priority Order)
1. **Sprint 0 – Foundations & Compliance Readiness (Week 1)**
   - Repository, CI/CD, multi-tenant architecture blueprint, environment configs, secret management, audit logging baseline, documentation scaffolding.
2. **Sprint 1 – Tenancy & Ledger Core (Weeks 2-3)**
   - Organisation/branch onboarding APIs, tenant RBAC, double-entry ledger engine, balance queries, FX rate service stubs, initial migrations & tests.
3. **Sprint 2 – MPesa Collections Core (Weeks 4-5)**
   - Daraja authentication, C2B registration, STK Push flow, callback handling, idempotent payment intents, error taxonomy.
4. **Sprint 3 – Expenses & AP (Weeks 6-7)**
   - Expenses module (categories, receipts, approvals, reimbursements); vendor bills (AP) & credits; COA mapping rules.
   - Dimensions/tags (project/cost center/vehicle/route/driver/branch); tax handling; multi‑currency.
5. **Sprint 4 – Mileage & Fuel (Weeks 8-9)**
   - Mileage logs & per‑km policy engine; fuel purchase ingestion; auto expense creation; validation and allocations.
   - Logistics-service ingestion endpoints; receipts storage; notifications on approvals.
6. **Sprint 5 – Invoicing & Document Engine (Weeks 10-11)**
   - Deploy baseline Fabric channel or Polygon contracts, chain anchoring service, settlement proof APIs, reconciliation hooks.
7. **Sprint 6 – Card & Wallet Collections (Weeks 12-13)**
   - Invoice lifecycle management, template designer, numbering rules, PDF generation, payment link issuance, multi-document support (quote, PO, delivery note, credit memo, etc.).
8. **Sprint 7 – Notifications & Dunning Automation (Weeks 14-15)**
   - Event contracts with `notifications-app`, email/SMS invoice delivery, reminder schedules, portal links, dunning workflows, customer communication preferences.
9. **Sprint 8 – Disbursements Engine (Weeks 16-17)**
   - MPesa B2C/B2B payouts, payout scheduling, approval workflow, treasury-held float management, branch-level settlement rules.
10. **Sprint 9 – Reconciliation & Reporting (Weeks 18-19)**
   - Statement ingestion, auto-matching, suspense queues, manual reconcile UI/API, financial reporting endpoints, blockchain verification.
11. **Sprint 10 – Compliance & Risk Controls (Weeks 20-21)**
   - KYC/KYB module, AML rules, sanctions screening integration, fraud monitoring hooks, audit exports.
12. **Sprint 11 – Budgets & Cost Centers (Weeks 22-23)**
   - Budgets for cost centers/projects; alerts; budget vs actuals; allocations.
13. **Sprint 12 – Bank Integrations & Treasury Portal (Weeks 24-25)**
   - Generic bank API adapter framework, first bank integration, treasury operations dashboard, approval workflows, branch cash positioning.
14. **Sprint 13 – Hardening, Performance & Launch (Weeks 26-27)**
   - Load testing, failover drills, penetration testing, documentation, production go-live checklist, blockchain/node resilience testing.

### 7. Backlog & Future Enhancements
- Instant payout cards, virtual IBAN issuance, smart-contract escrow market, cash-flow forecasting, dynamic currency conversion, crypto on/off ramps, machine learning risk scoring, ISO20022 messaging support, embedded finance marketplace for partner banks.

### 8. Deployment

- Containerised (Docker multi-stage) → Helm chart → ArgoCD GitOps
- Horizontal Pod Autoscaling (CPU/RPS), Pod Disruption Budgets
- Feature flags handled via treasury-managed config service (future)

### 9. Runtime Ports
- **Local development:** HTTP API listens on **4001** to avoid conflicts with the Food Delivery backend (4000) and Notifications service (4002).
- **Cloud deployment:** `TREASURY_HTTP_PORT` is overridden to **4000** so all backend services share a common service port behind ingress controllers.

### 10. Glossary & Acronyms (Plain‑English Reference)
- API / REST / gRPC / OpenAPI: Programmatic interfaces and protocols; REST uses HTTP verbs; gRPC is a high‑performance binary protocol; OpenAPI documents REST endpoints.
- OAuth2: An authorization standard for granting limited access (tokens) to APIs.
- mTLS (Mutual TLS): Both client and server authenticate each other using certificates.
- KMS / HSM: Key Management Service and Hardware Security Module for secure cryptographic key storage and signing.
- PCI DSS: Payment Card Industry Data Security Standard for cardholder data protection.
- MPesa (Daraja): Mobile money platform and its API; C2B (Customer‑to‑Business), B2C/B2B (Business‑to‑Customer/Business); STK Push prompts payment on a user’s handset.
- ACH / SWIFT: Bank transfer networks; ACH for domestic batch transfers, SWIFT for international messaging.
- KYC / KYB / AML / OFAC: Know Your Customer/Business, Anti‑Money Laundering, and sanctions screening (e.g., U.S. Treasury’s Office of Foreign Assets Control).
- IBAN: International Bank Account Number format for cross‑border payments.
- Ledger (double‑entry): Accounting method where every transaction has equal debits and credits.

---
**Next Steps:** Confirm regulatory requirements with compliance counsel, finalize MPesa credentials, align on shared event schemas with dependent services (including notifications), select blockchain rail, and schedule joint tenant‑onboarding and invoicing integration tests with partner applications.

