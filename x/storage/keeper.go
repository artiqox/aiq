package storage

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	coinKeeper bank.Keeper
	storeKey   types.StoreKey

	cdc *codec.Codec
}

func (k Keeper) StoreReport(ctx types.Context, name string, report []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), report)
}

func (k Keeper) RetrieveReport(ctx types.Context, name string) []byte {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(name)) {
		return nil
	}
	report := store.Get([]byte(name))
	return report
}

func NewKeeper(coinKeeper bank.Keeper, key types.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
}
