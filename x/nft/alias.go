package nft

import (
	"github.com/blockscapeLab/green/x/nft/keeper"
	"github.com/blockscapeLab/green/x/nft/types"
)

const (
	QuerySupply       = keeper.QuerySupply
	QueryOwner        = keeper.QueryOwner
	QueryOwnerByDenom = keeper.QueryOwnerByDenom
	QueryCollection   = keeper.QueryCollection
	QueryDenoms       = keeper.QueryDenoms
	QueryNFT          = keeper.QueryNFT

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	RegisterInvariants = keeper.RegisterInvariants
	AllInvariants      = keeper.AllInvariants
	SupplyInvariant    = keeper.SupplyInvariant
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	RegisterCodec      = types.RegisterCodec
	NewCollection      = types.NewCollection
	EmptyCollection    = types.EmptyCollection
	NewCollections     = types.NewCollections

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	ErrInvalidCollection = types.ErrInvalidCollection
	ErrUnknownCollection = types.ErrUnknownCollection
	ErrInvalidNFT        = types.ErrInvalidNFT
	ErrNFTAlreadyExists  = types.ErrNFTAlreadyExists
	ErrUnknownNFT        = types.ErrUnknownNFT
	ErrEmptyMetadata     = types.ErrEmptyMetadata

	GetCollectionKey = types.GetCollectionKey
	SplitOwnerKey    = types.SplitOwnerKey

	GetOwnersKey          = types.GetOwnersKey
	GetOwnerKey           = types.GetOwnerKey
	NewMsgTransferNFT     = types.NewMsgTransferNFT
	NewMsgEditNFTMetadata = types.NewMsgEditNFTMetadata
	NewMsgMintNFT         = types.NewMsgMintNFT
	NewMsgBurnNFT         = types.NewMsgBurnNFT
	NewBaseNFT            = types.NewBaseNFT
	NewNFTs               = types.NewNFTs

	NewIDCollection          = types.NewIDCollection
	NewOwner                 = types.NewOwner
	NewQueryCollectionParams = types.NewQueryCollectionParams
	NewQueryBalanceParams    = types.NewQueryBalanceParams
	NewQueryNFTParams        = types.NewQueryNFTParams

	// variable aliases
	ModuleCdc = types.ModuleCdc

	EventTypeTransfer        = types.EventTypeTransfer
	EventTypeEditNFTMetadata = types.EventTypeEditNFTMetadata
	EventTypeMintNFT         = types.EventTypeMintNFT
	EventTypeBurnNFT         = types.EventTypeBurnNFT
	AttributeValueCategory   = types.AttributeValueCategory
	AttributeKeySender       = types.AttributeKeySender
	AttributeKeyRecipient    = types.AttributeKeyRecipient
	AttributeKeyOwner        = types.AttributeKeyOwner
	AttributeKeyNFTID        = types.AttributeKeyNFTID
	AttributeKeyNFTTokenURI  = types.AttributeKeyNFTTokenURI
	AttributeKeyDenom        = types.AttributeKeyDenom
	CollectionsKeyPrefix     = types.CollectionsKeyPrefix
	OwnersKeyPrefix          = types.OwnersKeyPrefix
)

type (
	Keeper         = keeper.Keeper
	Collection     = types.Collection
	Collections    = types.Collections
	CollectionJSON = types.CollectionJSON
	GenesisState   = types.GenesisState

	Params = types.Params

	MsgTransferNFT     = types.MsgTransferNFT
	MsgEditNFTMetadata = types.MsgEditNFTMetadata
	MsgMintNFT         = types.MsgMintNFT
	MsgBurnNFT         = types.MsgBurnNFT
	BaseNFT            = types.BaseNFT
	NFTs               = types.NFTs

	IDCollection          = types.IDCollection
	IDCollections         = types.IDCollections
	Owner                 = types.Owner
	QueryCollectionParams = types.QueryCollectionParams
	QueryBalanceParams    = types.QueryBalanceParams
	QueryNFTParams        = types.QueryNFTParams
)
