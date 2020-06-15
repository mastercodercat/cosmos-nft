package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper /* TODO: Define what keepers the module needs */, data GenesisState) {
	k.SetOwners(ctx, data.Owners)

	for _, c := range data.Collections {
		// k.SetCollection(ctx, c.Denom, c)
		// VHX
		sortedCollection := NewCollection(c.Denom, c.NFTs.Sort())
		k.SetCollection(ctx, c.Denom, sortedCollection)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	return NewGenesisState(k.GetOwners(ctx), k.GetCollections(ctx))
	//return NewGenesisState()
}
