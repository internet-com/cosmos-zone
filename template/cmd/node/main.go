package main

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/wire"
	"_REMOTE_PROJECT_PATH_/app"
)

// init parameters
var CoolAppInit = server.AppInit{
	AppGenState: CoolAppGenState,
	AppGenTx:    server.SimpleAppGenTx,
}

// coolGenAppParams sets up the app_state and appends the cool app state
func CoolAppGenState(cdc *wire.Codec, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {
	appState, err = server.SimpleAppGenState(cdc, appGenTxs)
	if err != nil {
		return
	}

	key := "cool"
	value := json.RawMessage(`{
        "trend": "ice-cold"
      }`)

	appState, err = server.InsertKeyJSON(cdc, appState, key, value)
	if err != nil {
		return
	}

	key = "pow"
	value = json.RawMessage(`{
        "difficulty": "1",
        "count": "0"
      }`)

	appState, err = server.InsertKeyJSON(cdc, appState, key, value)
	return
}

func newApp(logger log.Logger, db dbm.DB) abci.Application {
	return app.NewDemocoinApp(logger, db)
}

func exportAppStateAndTMValidators(logger log.Logger, db dbm.DB) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	dapp := app.NewDemocoinApp(logger, db)
	return dapp.ExportAppStateAndValidators()
}

func main() {
	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "_PROJECT_SHORT_NAME_d",
		Short:             "_CAPITALIZED_PROJECT_SHORT_NAME_ Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	server.AddCommands(ctx, cdc, rootCmd, CoolAppInit,
		server.ConstructAppCreator(newApp, "_PROJECT_SHORT_NAME_"),
		server.ConstructAppExporter(exportAppStateAndTMValidators, "_PROJECT_SHORT_NAME_"))

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/._PROJECT_SHORT_NAME_d")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}
