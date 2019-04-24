package liquid

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestOrder(t *testing.T) {
	var k Config
	toml.DecodeFile(KEYPATH, &k)

	fmt.Printf("config load: %+v\n", k)

	c, err := New(k.ID, k.Token, nil)
	if err != nil {
		t.Error(err)
	}

	accounts, err := c.GetTradingAccounts()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("accouts: %+v\n", accounts)

	// o, err := c.CreateOrder(&RequestOrder{
	// 	Params{
	// 		OrderType:       "market",
	// 		ProductID:       5,
	// 		Side:            "sell",
	// 		Quantity:        "0.1",
	// 		LeverageLevel:   25,
	// 		FundingCurrency: "JPY",
	// 		OrderDirection:  "two_direction",
	// 	},
	// })
	// if err != nil {
	// 	t.Error(err)
	// }

	// fmt.Printf("%+v\n", o)
}
