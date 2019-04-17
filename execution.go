package liquid

import (
	"fmt"
	"strconv"
)

type Executions struct {
	Models      []ExecutionsModels `json:"models"`
	CurrentPage int                `json:"current_page"`
	TotalPages  int                `json:"total_pages"`
}

type ExecutionsModels struct {
	ID        int     `json:"id"`
	Quantity  float64 `json:"quantity,string"`
	Price     float64 `json:"price,string"`
	TakerSide string  `json:"taker_side"`
	MySide    string  `json:"my_side"`
	CreatedAt int64   `json:"created_at"`
}

func (c *Client) GetExecutionsByTimestamp(
	productID int,
	limit int,
	timestamp int) ([]ExecutionsModels, error) {
	req, err := c.newRequest("GET", "/executions", nil,
		&map[string]string{
			"product_id": strconv.Itoa(productID),
			"limit":      strconv.Itoa(limit),
			"timestamp":  strconv.Itoa(timestamp)})
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get data. status: %s", res.Status)
	}

	var executions []ExecutionsModels
	if err := decode(res, &executions); err != nil {
		return nil, err
	}

	return executions, nil
}

func (c *Client) GetExecutions(productID int, limit int, page int) (Executions, error) {
	var executions Executions
	req, err := c.newRequest("GET", "/executions", nil,
		&map[string]string{
			"product_id": strconv.Itoa(productID),
			"limit":      strconv.Itoa(limit),
			"page":       strconv.Itoa(page)})
	if err != nil {
		return executions, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return executions, err
	}

	if res.StatusCode != 200 {
		return executions, fmt.Errorf("failed to get data. status: %s", res.Status)
	}

	if err := decode(res, &executions); err != nil {
		return executions, err
	}

	return executions, nil
}

func (c *Client) GetOwnExecutions(productID int) (Executions, error) {
	var executions Executions
	req, err := c.newRequest("GET", "/executions/me", nil,
		&map[string]string{
			"product_id": strconv.Itoa(productID),
		})
	if err != nil {
		return executions, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return executions, err
	}

	if res.StatusCode != 200 {
		return executions, fmt.Errorf("failed to get data. status: %s", res.Status)
	}

	if err := decode(res, &executions); err != nil {
		return executions, err
	}

	return executions, nil
}
