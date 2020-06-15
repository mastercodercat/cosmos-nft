package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// GenesisState - all nft state that must be provided at genesis
type GenesisState struct {
	Owners      []Owner     `json:"owners"`
	Collections Collections `json:"collections"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(owners []Owner, collections Collections) GenesisState {
	return GenesisState{
		Owners:      owners,
		Collections: collections,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return NewGenesisState([]Owner{}, NewCollections())
}

// ValidateGenesis validates the nft genesis parameters
func ValidateGenesis(data GenesisState) error {
	for _, Owner := range data.Owners {
		if Owner.Address.Empty() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address cannot be empty")
		}
	}
	return nil
}
