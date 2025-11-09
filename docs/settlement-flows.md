# Settlement Flows

## Overview

1. **Aggregation** – Collect payable transactions per tenant/channel (MPesa, card, bank) for a settlement window.
2. **Approval** – Treasury operator reviews batch, applies adjustments, and approves for disbursement.
3. **Execution** – Disbursement processors initiate payouts (MPesa B2C/B2B, bank API) and capture transaction IDs.
4. **Reconciliation** – Incoming statements matched against expected disbursements, differences routed to suspense.
5. **Reporting** – Summary metrics pushed to analytics dashboards and exported to finance systems.

## Batch Lifecycle

| Status | Description |
| ------ | ----------- |
| `draft` | Batch created, awaiting funding confirmation |
| `pending_approval` | Batch ready for treasury review |
| `approved` | Approved for execution, queued for payout worker |
| `processing` | Disbursement in flight |
| `completed` | All payouts confirmed |
| `failed` | Batch halted; requires intervention |

## Failure Handling

- Automatic retry (configurable) for transient provider errors.
- Permanent failures routed to support queue with reason codes.
- Escalation triggers notifications to operations Slack channel + PagerDuty if SLA breached.

## Integration Points

- **Treasury Workers** – JetStream consumer `treasury.settlement.execute`
- **Notifications App** – Alerts stakeholders on approval/completion, triggers dunning workflows if funds fail
- **Food Delivery Backend** – Updates wallet balances and order payment status via events

## Future Enhancements

- Predictive cash forecasting using historical settlement data.
- Intelligent batching optimising for provider fees and FX rates.
- Real-time status dashboards (Grafana/Observe) with drill-down into batch health.
