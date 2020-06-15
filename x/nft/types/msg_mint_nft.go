package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* --------------------------------------------------------------------------- */
// MsgMintNFT
/* --------------------------------------------------------------------------- */

var _ sdk.Msg = &MsgMintNFT{}

// MsgMintNFT defines a MintNFT message
type MsgMintNFT struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
	TokenURI  string         `json:"token_uri" yaml:"token_uri"`
}

// NewMsgMintNFT is a constructor function for MsgMintNFT
func NewMsgMintNFT(sender, recipient sdk.AccAddress, id, denom, tokenURI string) MsgMintNFT {
	return MsgMintNFT{
		Sender:    sender,
		Recipient: recipient,
		ID:        strings.TrimSpace(id),
		Denom:     strings.TrimSpace(denom),
		TokenURI:  strings.TrimSpace(tokenURI),
	}
}


// Route Implements Msg
func (msg MsgMintNFT) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgMintNFT) Type() string { return "MintNFT" }

// ValidateBasic Implements Msg.
func (msg MsgMintNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidNFT
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidNFT
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

