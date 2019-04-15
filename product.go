package liquid

import (
	"fmt"
)

type Product struct {
	ID                  string  `json:"id"`
	ProductType         string  `json:"product_type"`
	Code                string  `json:"code"`
	Name                string  `json:"name"`
	MarketAsk           float64 `json:"market_ask,string"`
	MarketBid           float64 `json:"market_bid,string"`
	Indicator           int     `json:"indicator"`
	Currency            string  `json:"currency"`
	CurrencyPairCode    string  `json:"currency_pair_code"`
	Symbol              string  `json:"symbol"`
	FiatMinimumWithdraw float64 `json:"fiat_minimum_withdraw,string"`
	PusherChannel       string  `json:"pusher_channel"`
	TakerFee            float64 `json:"taker_fee,string"`
	MakerFee            float64 `json:"maker_fee,string"`
	LowMarketBid        float64 `json:"low_market_bid,string"`
	HighMarketAsk       float64 `json:"high_market_ask,string"`
	Volume24H           float64 `json:"volume_24h,string"`
	LastPrice24H        float64 `json:"last_price_24h,string"`
	LastTradedPrice     float64 `json:"last_traded_price,string"`
	LastTradedQuantity  float64 `json:"last_traded_quantity,string"`
	QuotedCurrency      string  `json:"quoted_currency"`
	BaseCurrency        string  `json:"base_currency"`
	ExchangeRate        float64 `json:"exchange_rate,string"`
}

func (c *Client) GetProducts() ([]Product, error) {
	res, err := c.sendRequest("GET", "/products", nil, nil)
	if err != nil {
		return nil, err
	}

	var products []Product
	if err := decode(res, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (c *Client) GetProduct(productID int) (Product, error) {
	var product Product
	spath := fmt.Sprintf("/products/%d", productID)
	res, err := c.sendRequest("GET", spath, nil, nil)
	if err != nil {
		return product, err
	}

	if err := decode(res, &product); err != nil {
		return product, err
	}

	return product, nil
}
