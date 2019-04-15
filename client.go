package liquid

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	BASEURL = "https://api.liquid.com"
	VERSION = "2"
)

type Client struct {
	URL         *url.URL
	TokenNumber string
	Secret      string
	HTTPClient  *http.Client

	Logger *log.Logger
}

func New(tokenNum string, secret string, logger *log.Logger) (*Client, error) {
	if tokenNum == "" {
		return nil, fmt.Errorf("API token number isnot set")
	}

	if secret == "" {
		return nil, fmt.Errorf("API secret key isnot set")
	}

	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	return &Client{URL: u, TokenNumber: tokenNum, Secret: secret, HTTPClient: client, Logger: logger}, nil

}

type InterestRates struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

func (c *Client) GetInterestRates(currency string) (InterestRates, error) {
	var rate InterestRates
	spath := fmt.Sprintf("/ir_ladders/%s", currency)
	res, err := c.sendRequest("GET", spath, nil, nil)
	if err != nil {
		return rate, err
	}

	if err := decode(res, &rate); err != nil {
		return rate, err
	}

	return rate, nil
}

func (c *Client) newRequest(method, spath string, body io.Reader, params *map[string]string) (*http.Request, error) {
	// avoid Pointer's butting
	u := *c.URL
	u.Path = c.URL.Path + spath

	if params != nil {
		q := u.Query()
		for k, v := range *params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"path":     spath,
		"nonce":    time.Now().UnixNano(),
		"token_id": c.TokenNumber,
	})

	tokenString, err := token.SignedString([]byte(c.Secret))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Quoine-API-Version", VERSION)
	req.Header.Set("X-Quoine-Auth", tokenString)

	return req, nil
}

func (c *Client) sendRequest(method, spath string, body io.Reader, params *map[string]string) (*http.Response, error) {
	req, err := c.newRequest(method, spath, body, params)
	c.Logger.Printf("Request:  %s \n", requestLog(req))
	if err != nil {
		c.Logger.Printf("err: %#v \n", err)
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	c.Logger.Printf("Response: %s \n", responseLog(res))
	if err != nil {
		c.Logger.Printf("err: %#v \n", err)
		return nil, err
	}

	if res.StatusCode != 200 {
		c.Logger.Printf("err: %#v \n", err)
		return nil, fmt.Errorf("faild to get data. status: %s", res.Status)
	}
	return res, nil
}

func decode(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

func responseLog(res *http.Response) string {
	b, _ := httputil.DumpResponse(res, true)
	return string(b)
}
func requestLog(req *http.Request) string {
	b, _ := httputil.DumpRequest(req, true)
	return string(b)
}
