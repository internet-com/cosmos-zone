package _PROJECT_SHORT_NAME_

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message:
			return handleSample_CAPITALIZED_PROJECT_SHORT_NAME_Message(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized message type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleSample_CAPITALIZED_PROJECT_SHORT_NAME_Message(ctx sdk.Context, k Keeper, msg Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message) sdk.Result {
	// Process the message here.
	return sdk.Result{
		Code: sdk.CodeOK,
	}
}
