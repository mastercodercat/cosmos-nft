package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* --------------------------------------------------------------------------- */
// MsgTransferNFT
/* --------------------------------------------------------------------------- */

var _ sdk.Msg = &MsgTransferNFT{}

// Msg<Action> - struct for unjailing jailed validator
type MsgTransferNFT struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"recipient" yaml:"recipient"`
	Denom     string         `json:"denom" yaml:"denom"`
	ID        string         `json:"id" yaml:"id"`
}

// NewMsg<Action> creates a new Msg<Action> instance
func NewMsgTransferNFT(sender, recipient sdk.AccAddress, denom, id string) MsgTransferNFT {
	return MsgTransferNFT{
		Sender:    sender,
		Recipient: recipient,
		Denom:     strings.TrimSpace(denom),
		ID:        strings.TrimSpace(id),
	}
}

// Route Implements Msg
func (msg MsgTransferNFT) Route() string { return RouterKey }

const TransferNFTConst = "TransferNFT"

// Type Implements Msg
func (msg MsgTransferNFT) Type() string { return TransferNFTConst }

// TODO
func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}

// ValidateBasic Implements Msg.
func (msg MsgTransferNFT) ValidateBasic() error {
	if strings.TrimSpace(msg.Denom) == "" {
		return ErrInvalidCollection
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	if msg.Recipient.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}
	if strings.TrimSpace(msg.ID) == "" {
		return ErrInvalidCollection
	}

	return nil
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
