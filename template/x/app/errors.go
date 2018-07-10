package _PROJECT_SHORT_NAME_

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Cool module reserves error 400-499 lawl
	CodeInvalid_CAPITALIZED_PROJECT_SHORT_NAME_Request sdk.CodeType = 400
	CodeParameterMissing                               sdk.CodeType = 401
)

// ErrIncorrectCoolAnswer - Error returned upon an incorrect guess
func ErrInvalidRequest(answer string) sdk.Error {
	return sdk.NewError(CodeInvalid_CAPITALIZED_PROJECT_SHORT_NAME_Request, fmt.Sprintf("Invalid Message Request: %v", answer))
}

func ErrParameterMissing(parameter string) sdk.Error {
	return sdk.NewError(CodeParameterMissing, fmt.Sprintf("Parameter %v is missing.", parameter))
}
