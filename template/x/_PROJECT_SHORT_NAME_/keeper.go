package _PROJECT_SHORT_NAME_

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper - handlers sets/gets of custom variables for your module
type Keeper struct {
	ck       bank.CoinKeeper
	storeKey sdk.StoreKey // The (unexposed) key used to access the store from the Context.
}

// NewKeeper - Returns the Keeper
func NewKeeper(key sdk.StoreKey, bankKeeper bank.CoinKeeper) Keeper {
	return Keeper{ck: bankKeeper, storeKey: key}
}

func (k Keeper) DeductFee(fee Fee) {
	// TODO: Subtract fee from the total available coins
}

// InitGenesis - store the genesis trend
func (k Keeper) InitGenesis(ctx sdk.Context, data json.RawMessage) error {
	var state GenesisState
	if err := json.Unmarshal(data, &state); err != nil {
		return err
	}
	return nil
}
