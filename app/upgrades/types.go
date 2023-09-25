package upgrades

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	adminmodulekeeper "github.com/cosmos/admin-module/x/adminmodule/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	builderkeeper "github.com/skip-mev/pob/x/builder/keeper"

	contractmanagerkeeper "github.com/neutron-org/neutron/x/contractmanager/keeper"
	cronkeeper "github.com/neutron-org/neutron/x/cron/keeper"
	feeburnerkeeper "github.com/neutron-org/neutron/x/feeburner/keeper"
	icqkeeper "github.com/neutron-org/neutron/x/interchainqueries/keeper"
	tokenfactorykeeper "github.com/neutron-org/neutron/x/tokenfactory/keeper"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(*module.Manager, module.Configurator, *UpgradeKeepers, StoreKeys, codec.Codec) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades store.StoreUpgrades
}

type UpgradeKeepers struct {
	// keepers
	IcqKeeper             icqkeeper.Keeper
	CronKeeper            cronkeeper.Keeper
	TokenFactoryKeeper    *tokenfactorykeeper.Keeper
	FeeBurnerKeeper       *feeburnerkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	BuilderKeeper         builderkeeper.Keeper
	ContractManagerKeeper contractmanagerkeeper.Keeper
	AdminModuleKeeper     adminmodulekeeper.Keeper
	WasmMsgServer         wasmtypes.MsgServer
	// subspaces
	GlobalFeeSubspace   paramtypes.Subspace
	CcvConsumerSubspace paramtypes.Subspace
}

type StoreKeys interface {
	GetKey(string) *store.KVStoreKey
}
