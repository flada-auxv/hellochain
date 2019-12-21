package hellochain

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/sdk-tutorials/hellochain/starter"
)

const appName = "hellochain"

var (
	// ModuleBasics variable
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
}

// NewHelloChainApp function
// NOTE: 01 の時点だと abci.Application が定義されてないとエラーになる
func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Applciation {
	appStarter := starter.NewAppStarter(appName, logger, db)

	var app = &helloChainApp{appStarter}
	app.InitializeStarter()

	return app
}
