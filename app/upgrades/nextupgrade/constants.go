package nextupgrade

import (
	"github.com/cosmos/cosmos-sdk/store/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/neutron-org/neutron/app/upgrades"
	auctiontypes "github.com/skip-mev/block-sdk/x/auction/types"
	dextypes "github.com/neutron-org/neutron/x/dex/types"
	epochstypes "github.com/neutron-org/neutron/x/epochs/types"
	incentivestypes "github.com/neutron-org/neutron/x/incentives/types"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "Next-Upgrade"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: types.StoreUpgrades{
		Added: []string{
			consensusparamtypes.ModuleName,
			crisistypes.ModuleName,
			auctiontypes.ModuleName,
			dextypes.ModuleName,
			incentivestypes.ModuleName,
			epochstypes.ModuleName,
		},
	},
}
