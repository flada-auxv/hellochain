package hellochain

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/sdk-tutorials/hellochain/starter"
)

const appName = "hellochain"

var (
	// ModuleBasics holds the AppModuleBasic struct of all modules included in the app
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter // helloChainApp extends starter.AppStarter
}

// NewHelloChainApp returns a fully constructed SDK application
func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {

	// construct our starter to extend
	appStarter := starter.NewAppStarter(appName, logger, db)

	// compose our app with starter
	var app = &helloChainApp{
		appStarter,
	}

	// do some final configuration...
	app.InitializeStarter()

	return app
}
