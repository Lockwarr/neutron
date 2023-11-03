package bindings

import (
	"encoding/json"
	incentivestypes "github.com/neutron-org/neutron/x/incentives/types"

	contractmanagertypes "github.com/neutron-org/neutron/x/contractmanager/types"

	feerefundertypes "github.com/neutron-org/neutron/x/feerefunder/types"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/neutron-org/neutron/x/interchainqueries/types"
)

// NeutronQuery contains neutron custom queries.
type NeutronQuery struct {
	// Registered Interchain Query Result for specified QueryID
	InterchainQueryResult *QueryRegisteredQueryResultRequest `json:"interchain_query_result,omitempty"`
	// Interchain account address for specified ConnectionID and OwnerAddress
	InterchainAccountAddress *QueryInterchainAccountAddressRequest `json:"interchain_account_address,omitempty"`
	// RegisteredInterchainQueries
	RegisteredInterchainQueries *QueryRegisteredQueriesRequest `json:"registered_interchain_queries,omitempty"`
	// RegisteredInterchainQuery
	RegisteredInterchainQuery *QueryRegisteredQueryRequest `json:"registered_interchain_query,omitempty"`
	// TotalBurnedNeutronsAmount
	TotalBurnedNeutronsAmount *QueryTotalBurnedNeutronsAmountRequest `json:"total_burned_neutrons_amount,omitempty"`
	// MinIbcFee
	MinIbcFee *QueryMinIbcFeeRequest `json:"min_ibc_fee,omitempty"`
	// Token Factory queries
	// Given a subdenom minted by a contract via `NeutronMsg::MintTokens`,
	// returns the full denom as used by `BankMsg::Send`.
	FullDenom *QueryFullDenom `json:"full_denom,omitempty"`
	// Returns the admin of a denom, if the denom is a Token Factory denom.
	DenomAdmin *QueryDenomAdmin `json:"denom_admin,omitempty"`
	// Returns the before send hook if it was set before
	BeforeSendHook *QueryBeforeSendHook `json:"before_send_hook,omitempty"`
	// Contractmanager queries
	// Query all failures for address
	Failures *QueryFailures `json:"failures,omitempty"`
	/// Incentives queries
	Incentives *QueryIncentives `json:"incentives,omitempty"`
}

/* Requests */

type QueryRegisteredQueryResultRequest struct {
	QueryID uint64 `json:"query_id,omitempty"`
}

type QueryInterchainAccountAddressRequest struct {
	// owner_address is the owner of the interchain account on the controller chain
	OwnerAddress string `json:"owner_address,omitempty"`
	// interchain_account_id is an identifier of your interchain account from which you want to execute msgs
	InterchainAccountID string `json:"interchain_account_id,omitempty"`
	// connection_id is an IBC connection identifier between Neutron and remote chain
	ConnectionID string `json:"connection_id,omitempty"`
}

type QueryRegisteredQueriesRequest struct {
	Owners       []string           `json:"owners,omitempty"`
	ConnectionID string             `json:"connection_id,omitempty"`
	Pagination   *query.PageRequest `json:"pagination,omitempty"`
}

type QueryRegisteredQueryRequest struct {
	QueryID uint64 `json:"query_id,omitempty"`
}

/* Responses */

type QueryRegisteredQueryResponse struct {
	RegisteredQuery *RegisteredQuery `json:"registered_query,omitempty"`
}

type QueryRegisteredQueriesResponse struct {
	RegisteredQueries []RegisteredQuery `json:"registered_queries"`
}

type RegisteredQuery struct {
	// The unique id of the registered query.
	ID uint64 `json:"id"`
	// The address that registered the query.
	Owner string `json:"owner"`
	// The KV-storage keys for which we want to get values from remote chain
	Keys []*types.KVKey `json:"keys"`
	// The filter for transaction search ICQ
	TransactionsFilter string `json:"transactions_filter"`
	// The query type identifier (i.e. 'kv' or 'tx' for now).
	QueryType string `json:"query_type"`
	// The IBC connection ID for getting ConsensusState to verify proofs.
	ConnectionID string `json:"connection_id"`
	// Parameter that defines how often the query must be updated.
	UpdatePeriod uint64 `json:"update_period"`
	// The local chain last block height when the query result was updated.
	LastSubmittedResultLocalHeight uint64 `json:"last_submitted_result_local_height"`
	// The remote chain last block height when the query result was updated.
	LastSubmittedResultRemoteHeight *ibcclienttypes.Height `json:"last_submitted_result_remote_height,omitempty"`
	// Amount of coins deposited for the query.
	Deposit sdktypes.Coins `json:"deposit"`
	// Timeout before query becomes available for everybody to remove.
	SubmitTimeout uint64 `json:"submit_timeout"`
	// The local chain height when the query was registered.
	RegisteredAtHeight uint64 `json:"registered_at_height"`
}

type QueryTotalBurnedNeutronsAmountRequest struct{}

type QueryTotalBurnedNeutronsAmountResponse struct {
	Coin sdktypes.Coin `json:"coin"`
}

type QueryMinIbcFeeRequest struct{}

type QueryMinIbcFeeResponse struct {
	MinFee feerefundertypes.Fee `json:"min_fee"`
}

func (rq RegisteredQuery) MarshalJSON() ([]byte, error) {
	type AliasRQ RegisteredQuery

	a := struct {
		AliasRQ
	}{
		AliasRQ: (AliasRQ)(rq),
	}

	// We want keys be as empty array in Json ('[]'), not 'null'
	// It's easier to work with on smart-contracts side
	if a.Keys == nil {
		a.Keys = make([]*types.KVKey, 0)
	}

	return json.Marshal(a)
}

// Query response for an interchain account address
type QueryInterchainAccountAddressResponse struct {
	// The corresponding interchain account address on the host chain
	InterchainAccountAddress string `json:"interchain_account_address,omitempty"`
}

type QueryRegisteredQueryResultResponse struct {
	Result *QueryResult `json:"result,omitempty"`
}

type QueryResult struct {
	KvResults []*StorageValue `json:"kv_results,omitempty"`
	Height    uint64          `json:"height,omitempty"`
	Revision  uint64          `json:"revision,omitempty"`
}

type StorageValue struct {
	StoragePrefix string `json:"storage_prefix,omitempty"`
	Key           []byte `json:"key"`
	Value         []byte `json:"value"`
}

func (sv StorageValue) MarshalJSON() ([]byte, error) {
	type AliasSV StorageValue

	a := struct {
		AliasSV
	}{
		AliasSV: (AliasSV)(sv),
	}

	// We want Key and Value be as empty arrays in Json ('[]'), not 'null'
	// It's easier to work with on smart-contracts side
	if a.Key == nil {
		a.Key = make([]byte, 0)
	}
	if a.Value == nil {
		a.Value = make([]byte, 0)
	}

	return json.Marshal(a)
}

type QueryFullDenom struct {
	CreatorAddr string `json:"creator_addr"`
	Subdenom    string `json:"subdenom"`
}

type QueryDenomAdmin struct {
	Subdenom string `json:"subdenom"`
}

type QueryBeforeSendHook struct {
	Denom string `json:"denom"`
}

type BeforeSendHookResponse struct {
	ContractAddr string `json:"contract_addr"`
}

type DenomAdminResponse struct {
	Admin string `json:"admin"`
}

type FullDenomResponse struct {
	Denom string `json:"denom"`
}

type QueryFailures struct {
	Address    string             `json:"address"`
	Pagination *query.PageRequest `json:"pagination,omitempty"`
}

type FailuresResponse struct {
	Failures []contractmanagertypes.Failure `json:"failures"`
}

type QueryIncentives struct {
	ModuleStatus         *ModuleStatus         `json:"module_status,omitempty"`
	GaugeByID            *GaugeByID            `json:"gauge_by_id,omitempty"`
	Gauges               *Gauges               `json:"gauges,omitempty"`
	StakeByID            *StakeByID            `json:"stake_by_id,omitempty"`
	Stakes               *Stakes               `json:"stakes,omitempty"`
	FutureRewardEstimate *FutureRewardEstimate `json:"future_reward_estimate,omitempty"`
	AccountHistory       *AccountHistory       `json:"account_history,omitempty"`
	GaugeQualifyingValue *GaugeQualifyingValue `json:"gauge_qualifying_value,omitempty"`
}

type ModuleStatus struct {
}

type ModuleStatusResponse struct {
	RewardCoins sdktypes.Coins         `json:"reward_coins"`
	Params      incentivestypes.Params `json:"params"`
}

type GaugeByID struct {
	ID uint64 `json:"id"`
}

type GaugeByIDResponse struct {
	Gauge *incentivestypes.Gauge `json:"gauge"`
}

type Gauges struct {
	Status string `json:"status"`
	Denom  string `json:"denom"`
}

type GaugesResponse struct {
	Gauges []*incentivestypes.Gauge `json:"gauges"`
}

type StakeByID struct {
	StakeID uint64 `json:"stake_id"`
}

type StakeByIDResponse struct {
	Stake *incentivestypes.Stake `json:"stake"`
}

type Stakes struct {
	Owner string `json:"owner"`
}

type StakesResponse struct {
	Stakes []*incentivestypes.Stake `json:"stakes"`
}

type FutureRewardEstimate struct {
	Owner     string   `json:"owner"`
	StakeIDs  []uint64 `json:"stake_ids"`
	NumEpochs int64    `json:"num_epochs"`
}

type FutureRewardEstimateResponse struct {
	Coins sdktypes.Coins `json:"coins"`
}

type AccountHistory struct {
	Account string `json:"account"`
}

type AccountHistoryResponse struct {
	Coins sdktypes.Coins `json:"coins"`
}

type GaugeQualifyingValue struct {
	ID uint64 `json:"id"`
}

type GaugeQualifyingValueResponse struct {
	QualifyingValue uint64 `json:"qualifying_value"`
}
