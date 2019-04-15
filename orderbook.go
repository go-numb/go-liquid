package liquid

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Depth struct {
	BuyPriceLevels  [][]json.Number `json:"buy_price_levels"`
	SellPriceLevels [][]json.Number `json:"sell_price_levels"`
}

func (c *Client) GetOrderBook(productID int, full bool) (Depth, error) {
	spath := fmt.Sprintf("/products/%d/price_levels", productID)

	var params *map[string]string
	if full {
		params = &map[string]string{
			"full": "1",
		}
	} else {
		params = nil
	}

	var depth Depth
	res, err := c.sendRequest("GET", spath, nil, params)
	if err != nil {
		return depth, err
	}

	if err := decode(res, &depth); err != nil {
		return depth, err
	}

	return depth, nil
}

func (p *Depth) GetSellDepthFloat64() [][]float64 {
	var sellFloat64 [][]float64
	for _, s := range p.SellPriceLevels {
		a, _ := s[0].Float64()
		b, _ := s[1].Float64()
		sellFloat64 = append(sellFloat64, []float64{a, b})
	}
	return sellFloat64
}

func (p *Depth) GetBuyDepthFloat64() [][]float64 {
	var buyFloat64 [][]float64
	for _, buy := range p.BuyPriceLevels {
		a, _ := buy[0].Float64()
		b, _ := buy[1].Float64()
		buyFloat64 = append(buyFloat64, []float64{a, b})
	}
	return buyFloat64
}

func (p *Depth) SortSellDepthByPrice(order string) [][]float64 {
	sortSell := p.GetSellDepthFloat64()
	if order == "asc" {
		sort.Slice(sortSell, func(i, j int) bool {
			return sortSell[i][0] < sortSell[j][0]
		})
	} else {
		sort.Slice(sortSell, func(i, j int) bool {
			return sortSell[i][0] > sortSell[j][0]
		})
	}
	return sortSell
}

func (p *Depth) SortBuyDepthByPrice(order string) [][]float64 {
	sortBuy := p.GetBuyDepthFloat64()

	if order == "asc" {
		sort.Slice(sortBuy, func(i, j int) bool {
			return sortBuy[i][0] < sortBuy[j][0]
		})
	} else {
		sort.Slice(sortBuy, func(i, j int) bool {
			return sortBuy[i][0] > sortBuy[j][0]
		})
	}
	return sortBuy
}

func (p *Depth) SortSellDepthByQuontity() [][]float64 {
	sortSell := p.GetSellDepthFloat64()
	sort.Slice(sortSell, func(i, j int) bool {
		return sortSell[i][1] > sortSell[j][1]
	})
	return sortSell
}

func (p *Depth) SortBuyDepthByQuontity() [][]float64 {
	sortBuy := p.GetBuyDepthFloat64()
	sort.Slice(sortBuy, func(i, j int) bool {
		return sortBuy[i][1] > sortBuy[j][1]
	})
	return sortBuy
}
