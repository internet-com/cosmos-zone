package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/commands"

	"_GITHUB_PROJECT_PATH_/x/_PROJECT_SHORT_NAME_"
)

// take the coolness quiz transaction
func Sample_CAPITALIZED_PROJECT_SHORT_NAME_Command(cdc *wire.Codec) *cobra.Command {
	command := &cobra.Command{
		Use:   "sample",
		Short: "Sample command",
		RunE: func(cmd *cobra.Command, args []string) error {

			message := viper.GetString("message")
			if message == "" {
				return _PROJECT_SHORT_NAME_.ErrParameterMissing("message")
			}
			// get the from address from the name flag
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			from, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}

			fee, err := strconv.ParseUint(viper.GetString("fee"), 10, 64)

			//fee, err := strconv.ParseUint(args[1], 10, 64)
			// create the message
			if err != nil {
				return err
			}
			msg := _PROJECT_SHORT_NAME_.NewSample_CAPITALIZED_PROJECT_SHORT_NAME_Message(from, message, fee)

			// get account name
			name := viper.GetString(client.FlagName)

			// build and sign the transaction, then broadcast to Tendermint
			res, err := ctx.SignBuildBroadcast(name, msg, cdc)
			if err != nil {
				return err
			}

			fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.Hash.String())
			return nil
		},
	}
	command.Flags().StringP("message", "m", "", "message")
	//command.Flags().Int64("fees", 0, "fee")
	return command
}
