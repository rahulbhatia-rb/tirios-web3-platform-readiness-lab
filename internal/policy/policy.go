package policy

type Contract struct {
	Service                    string `json:"service"`
	Environment                string `json:"environment"`
	PersistenceEnabled         bool   `json:"persistence_enabled"`
	AuthConfigured             bool   `json:"auth_configured"`
	DynamicRemoteCodeExecution bool   `json:"dynamic_remote_code_execution"`
	SecretsManaged             bool   `json:"secrets_managed"`
	RPCFailover                bool   `json:"rpc_failover"`
	IndexerEnabled             bool   `json:"indexer_enabled"`
	IdempotentEventIngestion   bool   `json:"idempotent_event_ingestion"`
	ReconciliationJob          bool   `json:"reconciliation_job"`
	WalletTransactionTracking  bool   `json:"wallet_transaction_tracking"`
	ContractEventsMapped       bool   `json:"contract_events_mapped"`
	StakingSourceOfTruth       string `json:"staking_source_of_truth"`
	Observability              bool   `json:"observability"`
	HealthChecks               bool   `json:"health_checks"`
	BackupRestore              bool   `json:"backup_restore"`
}

type Finding struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

type Result struct {
	Allowed  bool      `json:"allowed"`
	Findings []Finding `json:"findings,omitempty"`
}

func Evaluate(c Contract) Result {
	var f []Finding
	if c.Environment == "production" {
		if c.DynamicRemoteCodeExecution { f = append(f, Finding{"critical","remote_code_execution","remote/dynamic code execution is prohibited"}) }
		if !c.PersistenceEnabled { f = append(f, Finding{"critical","persistence_disabled","production persistence must be enabled"}) }
		if !c.AuthConfigured { f = append(f, Finding{"critical","auth_misconfigured","authentication configuration must validate successfully"}) }
		if !c.SecretsManaged { f = append(f, Finding{"high","secrets_unmanaged","secrets must come from managed secret storage"}) }
		if !c.RPCFailover { f = append(f, Finding{"high","rpc_single_point","RPC/provider failover is required"}) }
		if !c.IndexerEnabled || !c.IdempotentEventIngestion { f = append(f, Finding{"high","indexer_unreliable","chain events require an idempotent indexer"}) }
		if !c.ReconciliationJob { f = append(f, Finding{"high","reconciliation_missing","on-chain/off-chain reconciliation is required"}) }
		if !c.WalletTransactionTracking { f = append(f, Finding{"high","tx_tracking_missing","wallet transaction lifecycle tracking is required"}) }
		if !c.ContractEventsMapped { f = append(f, Finding{"high","contract_event_mapping_missing","business lifecycle changes must map to contract/backend events"}) }
		if c.StakingSourceOfTruth == "" { f = append(f, Finding{"high","staking_truth_undefined","staking/reward source of truth must be explicit"}) }
		if !c.Observability || !c.HealthChecks { f = append(f, Finding{"high","operability_missing","metrics/logs/traces and health checks are required"}) }
		if !c.BackupRestore { f = append(f, Finding{"high","recovery_missing","backup and restore path must be tested"}) }
	}
	allowed := true
	for _, x := range f { if x.Severity == "high" || x.Severity == "critical" { allowed = false } }
	return Result{Allowed: allowed, Findings: f}
}
