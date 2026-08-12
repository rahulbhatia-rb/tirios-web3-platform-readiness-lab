# Tirios Web3 Platform Readiness Lab

Independent proof-of-work inspired by the Tirios.ai MVP review and the public product direction around tokenized real estate, lifecycle tracking, staking, Web3 infrastructure, and production readiness.

The purpose is not to rewrite the product. It is to demonstrate how I would turn an MVP/demo surface into a safer, observable, state-consistent platform.

## What this demonstrates

- explicit property/token lifecycle state machine
- on-chain/off-chain event reconciliation
- idempotent blockchain event ingestion
- wallet/transaction state handling
- staking/reward state validation
- production-readiness gate
- no remote dynamic code execution
- secret/config validation
- persistence and API readiness checks
- indexer/RPC health controls
- observable transaction and lifecycle flows
- CI validation with safe and intentionally unsafe scenarios

## Architecture

```text
Property / admin action
        |
        v
 Backend workflow
        |
        +-----------> persistence
        |
        v
 Smart contract / chain
        |
        v
 blockchain events
        |
        v
 indexer / reconciler
        |
        +-----------> lifecycle state
        +-----------> investor dashboard
        +-----------> payout/reward views
        |
        v
 audit / observability
```

## Why this maps to the MVP review

The reviewed repository had a strong presentation layer, but the core production questions were around:
- where smart-contract and staking logic actually live;
- how lifecycle tracking is represented;
- how on-chain and off-chain state stay consistent;
- backend persistence and auth maturity;
- RPC/indexer strategy;
- deployment, observability, and secrets;
- removing unsafe dynamic execution paths.

This lab focuses on those seams.

## Run

```bash
go test ./...
go vet ./...
go run ./cmd/readiness -contract examples/property-prod.json
```

Expected:

```json
{
  "allowed": true
}
```

## Repository layout

```text
cmd/readiness/          executable platform gate
internal/policy/        production-readiness rules
internal/lifecycle/     lifecycle state model
examples/               safe + unsafe contracts
docs/architecture.md    platform design
docs/event-model.md     chain/off-chain reconciliation
docs/30-60-90.md        delivery roadmap
.github/workflows/      CI
```

## Disclaimer

Independent engineering prototype based on project context and code-review observations shared by the user.
