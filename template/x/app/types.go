package _PROJECT_SHORT_NAME_

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message struct {
	Sender  sdk.Address
	Message string
	Fee     uint64 // Optional
}

var _ sdk.Msg = Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message{}

// New cool message
func NewSample_CAPITALIZED_PROJECT_SHORT_NAME_Message(sender sdk.Address, message string, fee uint64) ClaimDomainMessage {
	return Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message{
		Sender:  sender,
		Message: message,
		Fee:     fee,
	}
}

// enforce the msg type at compile time

// nolint
func (msg Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message) Type() string { return "_PROJECT_SHORT_NAME_" }
func (msg Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message) Get(key interface{}) (value interface{}) {
	return nil
}
func (msg Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message) GetSigners() []sdk.Address {
	return []sdk.Address{msg.Sender}
}
func (msg Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message) String() string {
	return fmt.Sprintf("Sample_CAPITALIZED_PROJECT_SHORT_NAME_Message{Sender: %v, Message: %v}", msg.Sender, msg.Message)
}
