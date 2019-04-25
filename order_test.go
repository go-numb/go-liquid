package liquid

import (
	"fmt"
	"strings"
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

	var tradingAccountID int
	accounts, err := c.GetTradingAccounts()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Trading accounts")
	for i, v := range accounts {
		if strings.HasPrefix(v.CurrencyPairCode, "BTC") {
			fmt.Printf("accout %d: %+v\n", i, v)
			tradingAccountID = v.ID
		}
	}

	res, err := c.UpdateLeverageLevel(tradingAccountID, 25)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("leverage level: %+v\n", res)

	o, err := c.CreateOrder(&RequestOrder{
		OrderParams{
			OrderType:       "market",
			ProductID:       5,
			Side:            "sell",
			Quantity:        "0.001",
			LeverageLevel:   25,
			FundingCurrency: "JPY",
			OrderDirection:  "one_direction",
		},
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", o)
}
