package liquid

import (
	"encoding/json"
	"fmt"

	"strings"
)

type TradingAccount struct {
	ID               int         `json:"id"`
	LeverageLevel    int         `json:"leverage_level"`
	MaxLeverageLevel int         `json:"max_leverage_level"`
	Pnl              json.Number `json:"pnl"`
	Equity           json.Number `json:"equity"`
	Margin           json.Number `json:"margin"`
	FreeMargin       json.Number `json:"free_margin"`
	TraderID         int         `json:"trader_id"`
	Status           string      `json:"status"`
	ProductCode      string      `json:"product_code"`
	CurrencyPairCode string      `json:"currency_pair_code"`
	Position         json.Number `json:"position"`
	Balance          json.Number `json:"balance"`
	CreatedAt        int64       `json:"created_at"`
	UpdatedAt        int64       `json:"updated_at"`
	PusherChannel    string      `json:"pusher_channel"`
	MarginPercent    json.Number `json:"margin_percent"`
	ProductID        int         `json:"product_id"`
	FundingCurrency  string      `json:"funding_currency"`
}

func (c *Client) GetTradingAccounts() ([]TradingAccount, error) {
	res, err := c.sendRequest("GET", "/trading_accounts", nil, nil)
	if err != nil {
		return nil, err
	}

	var tradingAccounts []TradingAccount
	if err := decode(res, &tradingAccounts); err != nil {
		return nil, err
	}

	return tradingAccounts, nil
}

func (c *Client) GetATradingAccount(tradingAccountID int) (TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts/%d", tradingAccountID)

	var tradingAccount TradingAccount
	res, err := c.sendRequest("GET", spath, nil, nil)
	if err != nil {
		return tradingAccount, err
	}

	if err := decode(res, &tradingAccount); err != nil {
		return tradingAccount, err
	}

	return tradingAccount, nil
}

func (c *Client) UpdateLeverageLevel(tradeAccountID, leverageLevel int) (TradingAccount, error) {
	spath := fmt.Sprintf("/trading_accounts/%d", tradeAccountID)
	bodyTemplate :=
		`{
			"trading_account": {
				"leverage_level": %d
			}
		}`
	body := fmt.Sprintf(bodyTemplate, leverageLevel)

	var tradingAccount TradingAccount
	res, err := c.sendRequest("PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return tradingAccount, err
	}
	defer res.Body.Close()

	if err := decode(res, &tradingAccount); err != nil {
		return tradingAccount, err
	}

	return tradingAccount, nil
}
