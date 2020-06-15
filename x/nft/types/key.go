package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

const (
	// ModuleName is the name of the module
	ModuleName = "nft"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	// CollectionsKeyPrefix key for NFT collections
	CollectionsKeyPrefix = "ck-"
	// OwnersKeyPrefix key for balance of NFTs held by an address
	OwnersKeyPrefix = "ok-"
)

// GetCollectionKey gets the key of a collection
func GetCollectionKey(denom string) []byte {
	h := tmhash.New()
	_, err := h.Write([]byte(denom))
	if err != nil {
		panic(err)
	}
	bs := h.Sum(nil)

	return append([]byte(CollectionsKeyPrefix), bs...)
}

// SplitOwnerKey gets an address and denom from an owner key
func SplitOwnerKey(key []byte) (sdk.AccAddress, []byte) {
	if len(key) != 53 {
		panic(fmt.Sprintf("unexpected key length %d", len(key)))
	}
	address := key[1 : sdk.AddrLen+1]
	denomHashBz := key[sdk.AddrLen+1:]
	return sdk.AccAddress(address), denomHashBz
}

// GetOwnersKey gets the key prefix for all the collections owned by an account address
func GetOwnersKey(address sdk.AccAddress) []byte {
	return append([]byte(OwnersKeyPrefix), address.Bytes()...)
}

// GetOwnerKey gets the key of a collection owned by an account address
func GetOwnerKey(address sdk.AccAddress, denom string) []byte {
	h := tmhash.New()
	_, err := h.Write([]byte(denom))
	if err != nil {
		panic(err)
	}
	bs := h.Sum(nil)

	return append(GetOwnersKey(address), bs...)
}
