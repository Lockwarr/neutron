package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	dextypes "github.com/neutron-org/neutron/x/dex/types"
	epochstypes "github.com/neutron-org/neutron/x/epochs/types"
	feeburnertypes "github.com/neutron-org/neutron/x/feeburner/types"
	incentivestypes "github.com/neutron-org/neutron/x/incentives/types"
	interchainqueriestypes "github.com/neutron-org/neutron/x/interchainqueries/types"
	interchaintxstypes "github.com/neutron-org/neutron/x/interchaintxs/types"
	tokenfactorytypes "github.com/neutron-org/neutron/x/tokenfactory/types"
)

func AcceptedStargateQueries() wasmkeeper.AcceptedStargateQueries {
	return wasmkeeper.AcceptedStargateQueries{
		// ibc
		"/ibc.core.client.v1.Query/ClientState":    &ibcclienttypes.QueryClientStateResponse{},
		"/ibc.core.client.v1.Query/ConsensusState": &ibcclienttypes.QueryConsensusStateResponse{},
		"/ibc.core.connection.v1.Query/Connection": &ibcconnectiontypes.QueryConnectionResponse{},

		// token factory
		"/osmosis.tokenfactory.v1beta1.Query/Params":                 &tokenfactorytypes.QueryParamsResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomAuthorityMetadata": &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/DenomsFromCreator":      &tokenfactorytypes.QueryDenomsFromCreatorResponse{},
		"/osmosis.tokenfactory.v1beta1.Query/BeforeSendHookAddress":  &tokenfactorytypes.QueryBeforeSendHookAddressResponse{},

		// transfer
		"/ibc.applications.transfer.v1.Query/DenomTrace": &ibctransfertypes.QueryDenomTraceResponse{},

		// auth
		"/cosmos.auth.v1beta1.Query/Account": &authtypes.QueryAccountResponse{},
		"/cosmos.auth.v1beta1.Query/Params":  &authtypes.QueryParamsResponse{},

		// bank
		"/cosmos.bank.v1beta1.Query/Balance":       &banktypes.QueryBalanceResponse{},
		"/cosmos.bank.v1beta1.Query/DenomMetadata": &banktypes.QueryDenomsMetadataResponse{},
		"/cosmos.bank.v1beta1.Query/Params":        &banktypes.QueryParamsResponse{},
		"/cosmos.bank.v1beta1.Query/SupplyOf":      &banktypes.QuerySupplyOfResponse{},

		// interchaintxs
		"/neutron.interchaintxs.Query/Params": &interchaintxstypes.QueryParamsResponse{},

		// interchainqueries
		"/neutron.interchainqueries.Query/Params": &interchainqueriestypes.QueryParamsResponse{},

		// feeburner
		"/neutron.feeburner.Query/Params": &feeburnertypes.QueryParamsResponse{},

		// dex
		"/neutron.dex.Query/Params":                    &dextypes.QueryParamsResponse{},
		"/neutron.dex.Query/LimitOrderTrancheUser":     &dextypes.QueryGetLimitOrderTrancheUserResponse{},
		"/neutron.dex.Query/LimitOrderTranche":         &dextypes.QueryGetLimitOrderTrancheResponse{},
		"/neutron.dex.Query/UserDepositsAll":           &dextypes.QueryAllUserDepositsResponse{},
		"/neutron.dex.Query/UserLimitOrdersAll":        &dextypes.QueryAllUserLimitOrdersResponse{},
		"/neutron.dex.Query/InactiveLimitOrderTranche": &dextypes.QueryGetInactiveLimitOrderTrancheResponse{},
		"/neutron.dex.Query/PoolReserves":              &dextypes.QueryGetPoolReservesResponse{},
		"/neutron.dex.Query/EstimateMultiHopSwap":      &dextypes.QueryEstimateMultiHopSwapResponse{},
		"/neutron.dex.Query/EstimatePlaceLimitOrder":   &dextypes.QueryEstimatePlaceLimitOrderResponse{},

		// incentives
		"/neutron.incentives.Query/GetModuleStatus":         &incentivestypes.GetModuleStatusResponse{},
		"/neutron.incentives.Query/GetGaugeByID":            &incentivestypes.GetGaugeByIDResponse{},
		"/neutron.incentives.Query/GetGauges":               &incentivestypes.GetGaugesResponse{},
		"/neutron.incentives.Query/GetStakeByID":            &incentivestypes.GetStakeByIDResponse{},
		"/neutron.incentives.Query/GetStakes":               &incentivestypes.GetStakesResponse{},
		"/neutron.incentives.Query/GetFutureRewardEstimate": &incentivestypes.GetFutureRewardEstimateResponse{},
		"/neutron.incentives.Query/GetAccountHistory":       &incentivestypes.GetAccountHistoryResponse{},
		"/neutron.incentives.Query/GetGaugeQualifyingValue": &incentivestypes.GetGaugeQualifyingValueResponse{},

		// epochs
		"/neutron.epochs.Query/EpochInfos":   &epochstypes.QueryEpochsInfoResponse{},
		"/neutron.epochs.Query/CurrentEpoch": &epochstypes.QueryCurrentEpochResponse{},
	}
}
