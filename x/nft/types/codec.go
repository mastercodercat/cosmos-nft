package types

import (
	"github.com/blockscapeLab/green/x/nft/exported"
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&BaseNFT{}, "nft/BaseNFT", nil)
	cdc.RegisterConcrete(&IDCollection{}, "nft/IDCollection", nil)
	cdc.RegisterConcrete(&Collection{}, "nft/Collection", nil)
	cdc.RegisterConcrete(&Owner{}, "nft/Owner", nil)

	cdc.RegisterConcrete(MsgTransferNFT{}, "nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(MsgEditNFTMetadata{}, "nft/MsgEditNFTMetadata", nil)
	cdc.RegisterConcrete(MsgMintNFT{}, "nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(MsgBurnNFT{}, "nft/MsgBurnNFT", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
