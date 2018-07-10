package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/tmlibs/cli"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	"_GITHUB_PROJECT_PATH_/app"

	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// democoindCmd is the entry point for this binary
var (
	context                 = server.NewDefaultContext()
	_PROJECT_SHORT_NAME_Cmd = &cobra.Command{
		Use:               "_PROJECT_SHORT_NAME_",
		Short:             "_CAPITALIZED_PROJECT_NAME_ Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(context),
	}
)

// defaultAppState sets up the app_state for the
// default genesis file
func defaultAppState(args []string, addr sdk.Address, coinDenom string) (json.RawMessage, error) {
	baseJSON, err := server.DefaultGenAppState(args, addr, coinDenom)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]json.RawMessage
	err = json.Unmarshal(baseJSON, &jsonMap)
	if err != nil {
		return nil, err
	}
	bz, err := json.Marshal(jsonMap)
	return json.RawMessage(bz), err
}

func generateApp(rootDir string, logger log.Logger) (abci.Application, error) {
	dbMain, err := dbm.NewGoLevelDB("_PROJECT_SHORT_NAME_", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbAcc, err := dbm.NewGoLevelDB("_PROJECT_SHORT_NAME_-acc", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}
	dbs := map[string]dbm.DB{
		"main": dbMain,
		"acc":  dbAcc,
	}
	bapp := app.New_CAPITALIZED_PROJECT_SHORT_NAME_App(logger, dbs)
	return bapp, nil
}

func main() {

	server.AddCommands(_PROJECT_SHORT_NAME_Cmd, defaultAppState, generateApp, context)

	// prepare and add flags
	rootDir := os.ExpandEnv("$HOME/._PROJECT_SHORT_NAME_")
	executor := cli.PrepareBaseCmd(_PROJECT_SHORT_NAME_Cmd, "BC", rootDir)
	executor.Execute()
}
