package greeter

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	gtypes "github.com/cosmos/sdk-tutorials/hellochain/x/greeter/internal/types"
	"github.com/spf13/cobra"
)

// AppModuleBasic struct
type AppModuleBasic struct {
	starter.BlankModuleBasic
}

// AppModule struct
type AppModule struct {
	starter.BlankModule
	keeper     Keeper
	ModuleName string
}

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// RegisterCodec registers module Messages for encoding/decoding.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(gtypes.MsgGreet{}, "greeter/SayHello", nil)
}

// NewHandler returns a function for routing Messages to their appropriate handler functions.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// NewQuerierHandler returns a function for routing incoming Queries to the right querier.
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// QuerierRoute is used for routing Queries to this module.
func (am AppModule) QuerierRoute() string {
	return am.ModuleName
}

// GetQueryCmd assembles and returns all the clie query CLI commands supported by the module.
func (ab AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(gtypes.StoreKey, cdc)
}

// GetTxCmd assembles and returns all the clie query CLI commands supported by the module.
func (ab AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(gtypes.StoreKey, cdc)
}

// NewAppModule contstructs the full AppModule struct for this module.
func NewAppModule(keeper Keeper) AppModule {
	blank := starter.NewBlankModule(gtypes.ModuleName, keeper)
	return AppModule{blank, keeper, gtypes.ModuleName}
}
