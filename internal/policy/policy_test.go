package policy

import "testing"

func safe() Contract {
	return Contract{Service:"goldencity-platform",Environment:"production",PersistenceEnabled:true,AuthConfigured:true,DynamicRemoteCodeExecution:false,SecretsManaged:true,RPCFailover:true,IndexerEnabled:true,IdempotentEventIngestion:true,ReconciliationJob:true,WalletTransactionTracking:true,ContractEventsMapped:true,StakingSourceOfTruth:"contract-events+indexed-ledger",Observability:true,HealthChecks:true,BackupRestore:true}
}

func TestSafe(t *testing.T) { if !Evaluate(safe()).Allowed { t.Fatal("expected safe contract") } }
func TestRemoteExecRejected(t *testing.T) { c:=safe(); c.DynamicRemoteCodeExecution=true; if Evaluate(c).Allowed { t.Fatal("expected rejection") } }
func TestMissingIndexerRejected(t *testing.T) { c:=safe(); c.IndexerEnabled=false; if Evaluate(c).Allowed { t.Fatal("expected rejection") } }
