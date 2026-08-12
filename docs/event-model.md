# Event and Lifecycle Model

## Example events
- PropertyListed
- PropertyTokenized
- OfferingOpened
- OfferingFunded
- StakeCreated
- RewardsAccrued
- RewardsClaimed
- OwnershipTransferred
- PropertyExited

Each event should carry:
- deterministic event id
- chain id
- block number/hash
- transaction hash
- contract address
- entity/property id
- actor
- timestamp
- indexed/reconciled status

## Reorg safety

Do not treat the first observed block as final truth. Track confirmation depth and be able to invalidate/replay indexed records after a reorg.

## Reconciliation

A periodic job should compare indexed state against authoritative chain state and emit discrepancies.
