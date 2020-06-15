package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* --------------------------------------------------------------------------- */
// MsgEditNFTMetadata
/* --------------------------------------------------------------------------- */

var _ sdk.Msg = &MsgEditNFTMetadata{}

// MsgEditNFTMetadata edits an NFT's metadata
type MsgEditNFTMetadata struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	ID       string         `json:"id" yaml:"id"`
	Denom    string         `json:"denom" yaml:"denom"`
	TokenURI string         `json:"token_uri" yaml:"token_uri"`
}

// NewMsgEditNFTMetadata is a constructor function for MsgSetName
func NewMsgEditNFTMetadata(sender sdk.AccAddress, id,
	denom, tokenURI string,
) MsgEditNFTMetadata {
	return MsgEditNFTMetadata{
		Sender:   sender,
		ID:       strings.TrimSpace(id),
		Denom:    strings.TrimSpace(denom),
		TokenURI: strings.TrimSpace(tokenURI),
	}
}

// Route Implements Msg
func (msg MsgEditNFTMetadata) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgEditNFTMetadata) Type() string { return "EditNFTMetadata" }

// ValidateBasic Implements Msg.
func (msg MsgEditNFTMetadata) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgEditNFTMetadata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgEditNFTMetadata) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
