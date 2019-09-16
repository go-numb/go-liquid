package liquid

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-numb/pusher-websocket-go"
)

const (
	PUSHERAPPNAME           = ""
	PUSHERKEY               = "LIQUID"
	PUSHERTAPHOST           = "tap.liquid.com"
	PUSHERchEXECUTION       = "executions_cash_%s"
	PUSHERchDETAILEXECUTION = "execution_details_cash_%s"
	PUSHERchUSEREXECUTION   = "user_executions_cash_%s"
	PUSHERchASK             = "price_ladders_cash_%s_sell"
	PUSHERchBID             = "price_ladders_cash_%s_buy"
	PUSHERCREATED           = "created"
	PUSHERUPDATE            = "updated"

	BTCJPY  = "btcjpy"
	BTCUSD  = "btcusd"
	ETHJPY  = "ethjpy"
	ETHUSD  = "ethusd"
	XRPJPY  = "xrpjpy"
	XRPUSD  = "xrpusd"
	QASHJPY = "qashjpy"
)

type tap struct {
	c *pusher.Client
}

type Realtime struct {
	Results chan interface{}
}

// NewConnect is gets notif in results
func NewConnect() *Realtime {
	return &Realtime{
		Results: make(chan interface{}),
	}
}

// Connect is connected pusher for liquid websocket
func (p *Realtime) Connect(channels, products []string) {
	done := make(chan struct{})
	ws := &tap{c: pusher.New(PUSHERKEY, PUSHERTAPHOST)}

	m := make(map[string]*pusher.Channel, 0)
	for _, channel := range channels {
		for _, product := range products {
			subscribe := fmt.Sprintf(channel, product)
			m[subscribe] = ws.c.Subscribe(subscribe)
		}
	}

	for key, channel := range m {
		switch {
		case strings.Contains(key, "executions"):
			go p.handle(PUSHERCREATED, key, channel)
		default:
			go p.handle(PUSHERUPDATE, key, channel)
		}

	}

	<-done
}

// PusherExecution for Pusher
type PusherExecutionMe struct {
	ID        int   `json:"id"`
	OrderID   int   `json:"order_id"`
	CreatedAt int64 `json:"created_at"`

	Quantity      float64 `json:"quantity,string"`
	Price         float64 `json:"price,string"`
	TakerSide     string  `json:"taker_side"`
	MySide        string  `json:"my_side"`
	ClientOrderID string  `json:"client_order_id"`

	Delay time.Duration
}

// PusherExecution for Pusher
type PusherExecution struct {
	CreatedAt int64   `json:"created_at"`
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	TakerSide string  `json:"taker_side"`

	Delay time.Duration
}

// PusherBooks for Pusher
type PusherBooks struct {
	IsAsk bool
	Books Books
}

func (p *Realtime) handle(method, key string, channel *pusher.Channel) {
	switch { // Bind内部はgoroutine処理、故にhandle稼働時に処理を分けきる
	case strings.Contains(key, "user_executions_cash"): // 自己約定
		channel.Bind(method, func(data interface{}) { // Use CREATE
			b, ok := data.(string)
			if !ok {
				return
			}
			var s PusherExecutionMe
			json.Unmarshal([]byte(b), &s)
			s.Delay = toDelay(s.CreatedAt)
			p.Results <- s
		})

	case strings.Contains(key, "executions_cash"): // 約定
		channel.Bind(method, func(data interface{}) { // Use CREATE
			b, ok := data.(string)
			if !ok {
				return
			}
			var s PusherExecution
			json.Unmarshal([]byte(b), &s)
			s.Delay = toDelay(s.CreatedAt)
			p.Results <- s
		})

	case strings.Contains(key, "price_ladders"): // 板
		channel.Bind(method, func(data interface{}) { // Use CREATE
			b, ok := data.(string)
			if !ok {
				return
			}
			var s PusherBooks
			json.Unmarshal([]byte(b), &s.Books)
			s.IsAsk = isAsk(key)
			p.Results <- s
		})
	}
}

// toDelay is sub server time, api time.
func toDelay(apiTime int64) time.Duration {
	return time.Now().Sub(time.Unix(apiTime, 10))
}

func isAsk(key string) bool {
	if !strings.HasSuffix(key, "buy") {
		return true
	}
	return false
}
