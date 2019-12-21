package types

import (
	"fmt"
	"strings"

	codec "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName constant
	ModuleName = "greeter"

	// StoreKey constant
	StoreKey = ModuleName
)

var (
	// ModuleCdc variable
	ModuleCdc = codec.New()
)

// Greeting struct
type Greeting struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Body      string         `json:"body" yaml:"body"`
}

// GreetingList list of Greeting
type GreetingList []Greeting

// NewGreeting function
func NewGreeting(sender sdk.AccAddress, body string, receiver sdk.AccAddress) Greeting {
	return Greeting{
		Recipient: receiver,
		Sender:    sender,
		Body:      body,
	}
}

func (g Greeting) String() string {
	return strings.TrimSpace(
		fmt.Sprintf(`Sender: %s Recipient: %s Body: %s`, g.Sender.String(), g.Recipient.String(),
			g.Body),
	)
}

// QueryResGreetings map[string][]Greeting
type QueryResGreetings map[string][]Greeting

func (q QueryResGreetings) String() string {
	b := ModuleCdc.MustMarshalJSON(q)
	return string(b)
}

// NewQueryResGreetings function
func NewQueryResGreetings() QueryResGreetings {
	return make(map[string][]Greeting)
}
