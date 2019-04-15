package liquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Orders struct {
	Models      []Order `json:"models"`
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
}

type Order struct {
	ID                   int             `json:"id"`
	OrderType            string          `json:"order_type"`
	Quantity             float64         `json:"quantity,string"`
	DiscQuantity         float64         `json:"disc_quantity,string"`
	IcebergTotalQuantity string          `json:"iceberg_total_quantity"`
	Side                 string          `json:"side"`
	FilledQuantity       string          `json:"filled_quantity"`
	Price                float64         `json:"price,string"`
	CreatedAt            int64           `json:"created_at"`
	UpdatedAt            int64           `json:"updated_at"`
	Status               string          `json:"status"`
	LeverageLevel        int             `json:"leverage_level"`
	SourceExchange       string          `json:"source_exchange"`
	ProductID            int             `json:"product_id"`
	ProductCode          string          `json:"product_code"`
	FundingCurrency      string          `json:"funding_currency"`
	CurrencyPairCode     string          `json:"currency_pair_code"`
	OrderFee             float64         `json:"order_fee,string"`
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

func (c *Client) GetOrders(productID, withDetails int, fundingCurrency, status string) (Orders, error) {
	params := &map[string]string{
		"product_id":       strconv.Itoa(productID),
		"with_details":     strconv.Itoa(withDetails),
		"status":           status,
		"funding_currency": fundingCurrency}

	var orders Orders
	res, err := c.sendRequest("GET", "/orders", nil, params)
	if err != nil {
		return orders, err
	}

	if err := decode(res, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

type RequestOrder struct {
	Order struct {
		OrderType  string `json:"order_type"`
		ProductID  int    `json:"product_id"`
		Side       string `json:"side"`
		Quantity   string `json:"quantity"`
		Price      string `json:"price"`
		PriceRange string `json:"price_range,omitempty"`
		// Margin trade
		LeverageLevel   int    `json:"leverage_level,omitempty"`
		FundingCurrency string `json:"funding_currency,omitempty"`
		OrderDirection  string `json:"order_direction,omitempty"`
	} `json:"order"`
}

// orderType, side, quantity, price, priceRange string, productID int
func (c *Client) CreateOrder(o RequestOrder) (Order, error) {
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

func (c *Client) EditLiveOrder(orderID int, quantity, price string) (Order, error) {
	spath := fmt.Sprintf("/orders/%d", orderID)
	bodyTemplate :=
		`{
			"order": {
				"quantity":"%s",
				"price":"%s",
			}
		}`
	body := fmt.Sprintf(bodyTemplate, quantity, price)

	var order Order
	res, err := c.sendRequest("PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return order, err
	}

	if err := decode(res, &order); err != nil {
		return order, err
	}

	return order, nil
}
