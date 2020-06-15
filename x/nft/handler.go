package nft

import (
	"fmt"

	"github.com/blockscapeLab/green/x/nft/keeper"
	"github.com/blockscapeLab/green/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the nft type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgTransferNFT:
			return HandleMsgTransferNFT(ctx, msg, k)
		case types.MsgEditNFTMetadata:
			return HandleMsgEditNFTMetadata(ctx, msg, k)
		case types.MsgMintNFT:
			return HandleMsgMintNFT(ctx, msg, k)
		case types.MsgBurnNFT:
			return HandleMsgBurnNFT(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// HandleMsgTransferNFT handler for MsgTransferNFT
func HandleMsgTransferNFT(ctx sdk.Context, msg types.MsgTransferNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}
	// update NFT owner
	nft.SetOwner(msg.Recipient)
	// update the NFT (owners are updated within the keeper)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgEditNFTMetadata handler for MsgEditNFTMetadata
func HandleMsgEditNFTMetadata(ctx sdk.Context, msg types.MsgEditNFTMetadata, k keeper.Keeper,
) (*sdk.Result, error) {
	nft, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	// update NFT
	nft.EditMetadata(msg.TokenURI)
	err = k.UpdateNFT(ctx, msg.Denom, nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeEditNFTMetadata,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
			sdk.NewAttribute(types.AttributeKeyNFTTokenURI, msg.TokenURI),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgMintNFT handles MsgMintNFT
func HandleMsgMintNFT(ctx sdk.Context, msg types.MsgMintNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	nft := types.NewBaseNFT(msg.ID, msg.Recipient, msg.TokenURI)
	err := k.MintNFT(ctx, msg.Denom, &nft)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintNFT,
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
			sdk.NewAttribute(types.AttributeKeyNFTTokenURI, msg.TokenURI),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// HandleMsgBurnNFT handles MsgBurnNFT
func HandleMsgBurnNFT(ctx sdk.Context, msg types.MsgBurnNFT, k keeper.Keeper,
) (*sdk.Result, error) {
	_, err := k.GetNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	// remove  NFT
	err = k.DeleteNFT(ctx, msg.Denom, msg.ID)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBurnNFT,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyNFTID, msg.ID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
