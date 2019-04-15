package liquid

import (
	"fmt"
	"strconv"
	"strings"
)

type Loans struct {
	Models      []*Loan `json:"models"`
	CurrentPage int     `json:"current_page"`
	TotalPages  int     `json:"total_pages"`
}

type Loan struct {
	ID           int    `json:"id"`
	Quantity     string `json:"quantity"`
	Rate         string `json:"rate"`
	CreatedAt    int64  `json:"created_at"`
	LenderID     int    `json:"lender_id"`
	BorrowerID   int    `json:"borrower_id"`
	Status       string `json:"status"`
	Currency     string `json:"currency"`
	FundReloaned bool   `json:"fund_reloaned"`
}

func (c *Client) GetTradesLoans(tradeID int) ([]Loan, error) {
	spath := fmt.Sprintf("/trades/%d/loans", tradeID)
	res, err := c.sendRequest("GET", spath, nil, nil)
	if err != nil {
		return nil, err
	}

	var loans []Loan
	if err := decode(res, &loans); err != nil {
		return nil, err
	}

	return loans, nil
}

func (c *Client) GetLoans(currency string) (Loans, error) {
	queryParam := &map[string]string{
		"currency": currency}

	var loans Loans
	res, err := c.sendRequest("GET", "/loans", nil, queryParam)
	if err != nil {
		return loans, err
	}

	if err := decode(res, &loans); err != nil {
		return loans, err
	}

	return loans, nil
}

func (c *Client) UpdateALoan(loanID int, fundReloaned bool) (Loan, error) {
	spath := fmt.Sprintf("/loans/%d", loanID)
	bodyTemplate :=
		`{
			"loan": {
				"fund_reloaned":%s
			}
		}`
	body := fmt.Sprintf(bodyTemplate, strconv.FormatBool(fundReloaned))

	var loan Loan
	res, err := c.sendRequest("PUT", spath, strings.NewReader(body), nil)
	if err != nil {
		return loan, err
	}

	if err := decode(res, &loan); err != nil {
		return loan, err
	}

	return loan, nil
}
