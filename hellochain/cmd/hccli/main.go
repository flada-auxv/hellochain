package main

import (
	app "github.com/cosmos/sdk-tutorials/hellochain"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	"github.com/cosmos/sdk-tutorials/hellochain/x/greeter"
	"github.com/tendermint/tendermint/libs/cli"
)

func main() {
	starter.BuildModuleBasics(greeter.AppModuleBasic{})

	rootCmd := starter.NewCLICommand()

	txCmd := starter.TxCmd(starter.Cdc)
	queryCmd := starter.QueryCmd(starter.Cdc)

	app.ModuleBasics.AddTxCommands(txCmd, starter.Cdc)
	app.ModuleBasics.AddQueryCommands(queryCmd, starter.Cdc)
	rootCmd.AddCommand(txCmd, queryCmd)

	executor := cli.PrepareMainCmd(rootCmd, "HC", starter.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
