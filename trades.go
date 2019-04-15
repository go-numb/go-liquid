package liquid

import (
	"fmt"
	"strings"
)

type Trades struct {
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
	Models      []Trade `json:"models"`
}

type Trade struct {
	ID               int     `json:"id"`
	CurrencyPairCode string  `json:"currency_pair_code"`
	Status           string  `json:"status"`
	Side             string  `json:"side"`
	MarginUsed       float64 `json:"margin_used,string"`
	OpenQuantity     float64 `json:"open_quantity,string"`
	CloseQuantity    float64 `json:"close_quantity,string"`
	Quantity         float64 `json:"quantity,string"`
	LeverageLevel    int     `json:"leverage_level"`
	ProductCode      string  `json:"product_code"`
	ProductID        int     `json:"product_id"`
	OpenPrice        float64 `json:"open_price,string"`
	ClosePrice       float64 `json:"close_price,string"`
	TraderID         int     `json:"trader_id"`
	OpenPnl          float64 `json:"open_pnl,string"`
	ClosePnl         float64 `json:"close_pnl,string"`
	Pnl              float64 `json:"pnl,string"`
	StopLoss         float64 `json:"stop_loss,string"`
	TakeProfit       float64 `json:"take_profit,string"`
	FundingCurrency  string  `json:"funding_currency"`
	CreatedAt        int64   `json:"created_at"`
	UpdatedAt        int64   `json:"updated_at"`
	CloseFee         float64 `json:"close_fee,string"`
	TotalInterest    float64 `json:"total_interest,string"`
	DailyInterest    float64 `json:"daily_interest,string"`
}

func (c *Client) GetOrderTrades(orderID int) ([]Trade, error) {
	spath := fmt.Sprintf("/orders/%d/trades", orderID)
	res, err := c.sendRequest("GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err := decode(res, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func (c *Client) GetTrades(fundingCurrency, status string) (Trades, error) {
	queryParam := &map[string]string{
		"funding_currency": fundingCurrency,
		"status":           status}

	var trades Trades
	res, err := c.sendRequest("GET", "/trades", nil, queryParam)
	if err != nil {
		return trades, err
	}

	if err := decode(res, &trades); err != nil {
		return trades, err
	}

	return trades, nil
}

func (c *Client) CloseTrade(tradeID int, closedQuantity float64) (Trade, error) {
	spath := fmt.Sprintf("/trades/%d/close", tradeID)
	bodyTemplate := `{"closed_quantity":%f}`
	body := fmt.Sprintf(bodyTemplate, closedQuantity)

	var trade Trade
	res, err := c.sendRequest("PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return trade, err
	}

	if err := decode(res, &trade); err != nil {
		return trade, err
	}

	return trade, nil
}

func (c *Client) CloseAllTrade(side string) ([]Trade, error) {
	body := fmt.Sprintf(`{"side":"%s"}`, side)
	res, err := c.sendRequest("PUT", "/trades/close_all", strings.NewReader(body), nil)
	if err != nil {
		return nil, err
	}

	var trades []Trade
	if err := decode(res, &trades); err != nil {
		return nil, err
	}

	return trades, nil
}

func (c *Client) UpdateTrade(tradeID, stopLoss, takeProfit int) (Trade, error) {
	spath := fmt.Sprintf("/trades/%d", tradeID)
	bodyTemplate :=
		`{
			"trade": {
				"stop_loss":"%d",
				"take_profit":"%d"
			}
		}`
	body := fmt.Sprintf(bodyTemplate, stopLoss, takeProfit)

	var trade Trade
	res, err := c.sendRequest("PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return trade, err
	}

	if err := decode(res, &trade); err != nil {
		return trade, err
	}

	return trade, nil
}
