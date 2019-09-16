package liquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
)

type Orders struct {
	Models      []Order `json:"models"`
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
}

type Order struct {
	ID                   int             `json:"id"`
	OrderType            string          `json:"order_type"`
	Quantity             json.Number     `json:"quantity"`
	DiscQuantity         json.Number     `json:"disc_quantity"`
	IcebergTotalQuantity string          `json:"iceberg_total_quantity"`
	Side                 string          `json:"side"`
	FilledQuantity       string          `json:"filled_quantity"`
	Price                json.Number     `json:"price"`
	CreatedAt            int64           `json:"created_at"`
	UpdatedAt            int64           `json:"updated_at"`
	Status               string          `json:"status"`
	LeverageLevel        int             `json:"leverage_level"`
	SourceExchange       string          `json:"source_exchange"`
	ProductID            int             `json:"product_id"`
	ProductCode          string          `json:"product_code"`
	FundingCurrency      string          `json:"funding_currency"`
	CurrencyPairCode     string          `json:"currency_pair_code"`
	OrderFee             json.Number     `json:"order_fee"`
	Executions           OrderExecutions `json:"executions"`
}

type OrderExecutions []struct {
	ID        int     `json:"id"`
	Quantity  float64 `json:"quantity,string"`
	Price     float64 `json:"price,string"`
	TakerSide string  `json:"taker_side"`
	MySide    string  `json:"my_side"`
	CreatedAt int64   `json:"created_at"`
}

func (c *Client) GetOrder(orderID int) (Order, error) {
	spath := fmt.Sprintf("/orders/%d", orderID)

	var order Order
	res, err := c.sendRequest("GET", spath, nil, nil)
	if err != nil {
		return order, err
	}

	if err := decode(res, &order); err != nil {
		return order, err
	}

	return order, nil
}

type OrdersFilter struct {
	ProductID       string `json:"product_id,omitempty"`
	WithDetails     string `json:"with_details,omitempty"`
	Status          string `json:"status,omitempty"`
	FundingCurrency string `json:"funding_currency,omitempty"`

	// 下記Doc非公開フィルター
	// WebAPI: page=1&limit=24&currency_pair_code=BTCJPY&status=live&trading_type=cfd
	Page             string `json:"page,omitempty"`
	Limit            string `json:"limit,omitempty"`
	CurrencyPairCode string `json:"currency_pair_code,omitempty"`
	TradingType      string `json:"trading_type,omitempty"`
}

func (c *Client) GetOrders(filters OrdersFilter) (Orders, error) {
	var orders Orders

	j, err := json.Marshal(filters)
	if err != nil {
		return orders, err
	}

	res, err := c.sendRequest("GET", "/orders", bytes.NewReader(j), nil)
	if err != nil {
		return orders, err
	}

	if err := decode(res, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

type RequestOrder struct {
	Order OrderParams `json:"order"`
}

type OrderParams struct {
	OrderType  string `json:"order_type"`
	ProductID  int    `json:"product_id"`
	Side       string `json:"side"`
	Quantity   string `json:"quantity"`
	Price      string `json:"price,omitempty"`
	PriceRange string `json:"price_range,omitempty"`
	// Margin trade
	LeverageLevel   int    `json:"leverage_level,omitempty"`
	FundingCurrency string `json:"funding_currency,omitempty"`
	OrderDirection  string `json:"order_direction,omitempty"`
}

// orderType, side, quantity, price, priceRange string, productID int
func (c *Client) CreateOrder(o *RequestOrder) (Order, error) {
	var order Order

	body, err := json.Marshal(o)
	if err != nil {
		return order, err
	}

	res, err := c.sendRequest("POST", "/orders/", bytes.NewReader(body), nil)
	if err != nil {
		return order, err
	}

	if err := decode(res, &order); err != nil {
		return order, err
	}

	return order, nil
}

func (c *Client) CancelOrder(orderID int) (Order, error) {
	spath := fmt.Sprintf("/orders/%d/cancel", orderID)

	var order Order
	res, err := c.sendRequest("PUT", spath, nil, nil)
	if err != nil {
		return order, err
	}

	if err := decode(res, &order); err != nil {
		return order, err
	}

	return order, nil
}

type EditOrderParams struct {
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

// func (c *Client) EditLiveOrder(orderID int, quantity, price string) (Order, error) {
func (c *Client) EditLiveOrder(id int, e *EditOrderParams) (Order, error) {
	spath := fmt.Sprintf("/orders/%d", id)
	// bodyTemplate :=
	// 	`{
	// 		"order": { // Orderセクションは実はいらない
	// 			"quantity":	"%s",
	// 			"price":	"%s"
	// 		}
	// 	}`
	// body := fmt.Sprintf(bodyTemplate, quantity, price)

	var order Order

	body, err := json.Marshal(e)
	if err != nil {
		return order, err
	}

	res, err := c.sendRequest("PUT", spath, bytes.NewReader(body), nil)
	if err != nil {
		return order, err
	}

	if err := decode(res, &order); err != nil {
		return order, err
	}

	return order, nil
}

// https://api.liquid.com/trades/close_all?funding_currency=JPY&product_id=5
// OrderAllClose
func (p *Client) OrderAllClose() {

}

// ToPriceString is float to string price.00001
func ToPriceString(price float64) string {
	price = math.RoundToEven(price*100000) / 100000

	return fmt.Sprintf("%.5f", price)
}

// ToQtyString is float to string size.001
func ToQtyString(size float64) string {
	size = math.RoundToEven(size*1000) / 1000

	return fmt.Sprintf("%.3f", size)
}
