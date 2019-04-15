package liquid

import (
	"fmt"
	"strings"
)

type LoanBid struct {
	ID             int     `json:"id"`
	BidaskType     string  `json:"bidask_type"`
	Quantity       float64 `json:"quantity,string"`
	Currency       string  `json:"currency"`
	Side           string  `json:"side"`
	FilledQuantity float64 `json:"filled_quantity,string"`
	Status         string  `json:"status"`
	Rate           float64 `json:"rate,string"`
	UserID         int     `json:"user_id"`
}

type LoanBids struct {
	Models      []LoanBid `json:"models"`
	CurrentPage int       `json:"current_page"`
	TotalPages  int       `json:"total_pages"`
}

func (c *Client) CreateLoanBid(quantity, currency, rate string) (LoanBid, error) {
	bodyTemplate :=
		`{
			"loan_bid": {
				"quantity":"%s",
				"currency":"%s",
				"rate":"%s"
			}
		}`
	body := fmt.Sprintf(bodyTemplate, quantity, currency, rate)

	var loanBid LoanBid
	res, err := c.sendRequest("POST", "/loan_bids", strings.NewReader(body), nil)
	if err != nil {
		return loanBid, err
	}

	if err := decode(res, &loanBid); err != nil {
		return loanBid, err
	}

	return loanBid, nil
}

func (c *Client) GetLoanBids(currency string) (LoanBids, error) {
	queryParam := &map[string]string{
		"currency": currency}

	var loanBids LoanBids
	res, err := c.sendRequest("GET", "/loan_bids", nil, queryParam)
	if err != nil {
		return loanBids, err
	}

	if err := decode(res, &loanBids); err != nil {
		return loanBids, err
	}

	return loanBids, nil
}

func (c *Client) CloseLoanBid(loanBidID int) (LoanBid, error) {
	spath := fmt.Sprintf("/loan_bids/%d/close", loanBidID)

	var loanBid LoanBid
	res, err := c.sendRequest("PUT", spath, nil, nil)
	if err != nil {
		return loanBid, err
	}

	if err := decode(res, &loanBid); err != nil {
		return loanBid, err
	}

	return loanBid, nil
}
