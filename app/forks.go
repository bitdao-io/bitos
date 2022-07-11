package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ScheduleForkUpgrade executes any necessary fork logic for based upon the current
// block height and chain ID (mainnet or testnet). It sets an upgrade plan once
// the chain reaches the pre-defined upgrade height.
//
// CONTRACT: for this logic to work properly it is required to:
//
// 	1) Release a non-breaking patch version so that the chain can set the scheduled upgrade plan at upgrade-height.
// 	2) Release the software defined in the upgrade-info
func (app *Bitos) ScheduleForkUpgrade(ctx sdk.Context) {
	// todo: for version update
}
