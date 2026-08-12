# Architecture

## Boundaries

The platform should make four boundaries explicit:

1. **Chain state** — ownership, transfers, staking/reward events where contract-driven.
2. **Backend state** — KYC, legal/property metadata, workflows, payout operations.
3. **Indexed state** — normalized event ledger powering APIs and dashboards.
4. **Presentation state** — derived views only; not a hidden source of truth.

## Production path

- API service with validated auth/config
- durable database
- managed secrets
- primary + secondary RPC providers
- event indexer with checkpoints
- idempotent writes
- reconciliation against chain head
- transaction confirmation tracking
- observability around lag, failures, reorgs and RPC health
