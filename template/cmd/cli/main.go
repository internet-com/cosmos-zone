package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/tendermint/tmlibs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"

	_PROJECT_SHORT_NAME_cmd "_REMOTE_PROJECT_PATH_/x/_PROJECT_SHORT_NAME_/commands"

	"github.com/cosmos/cosmos-sdk/version"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/commands"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/commands"

	"_REMOTE_PROJECT_PATH_/app"
	"_REMOTE_PROJECT_PATH_/types"
)

// gaiacliCmd is the entry point for this binary
var (
	_PROJECT_SHORT_NAME_clicmd = &cobra.Command{
		Use:   "_PROJECT_SHORT_NAME_cli",
		Short: "_PROJECT_SHORT_NAME_ light-client",
	}
)

func todoNotImplemented(_ *cobra.Command, _ []string) error {
	return errors.New("TODO: Command not yet implemented")
}

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := app.MakeCodec()

	rpc.AddCommands(_PROJECT_SHORT_NAME_clicmd)
	_PROJECT_SHORT_NAME_clicmd.AddCommand(client.LineBreak)
	tx.AddCommands(_PROJECT_SHORT_NAME_clicmd, cdc)
	_PROJECT_SHORT_NAME_clicmd.AddCommand(client.LineBreak)

	// add query/post commands (custom to binary)
	_PROJECT_SHORT_NAME_clicmd.AddCommand(
		client.GetCommands(
			authcmd.GetAccountCmd("main", cdc, types.GetAccountDecoder(cdc)),
		)...)
	_PROJECT_SHORT_NAME_clicmd.AddCommand(
		client.PostCommands(
			bankcmd.SendTxCmd(cdc),
		)...)
	_PROJECT_SHORT_NAME_clicmd.AddCommand(
		client.PostCommands(
			_PROJECT_SHORT_NAME_cmd.Sample_CAPITALIZED_PROJECT_SHORT_NAME_Command(cdc),
		)...)

	// add proxy, version and key info
	_PROJECT_SHORT_NAME_clicmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(_PROJECT_SHORT_NAME_clicmd, "BC", os.ExpandEnv("$HOME/._PROJECT_SHORT_NAME_cli"))
	executor.Execute()
}
