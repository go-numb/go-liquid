package liquid

import (
	"errors"
	"sync"
)

// Limit is managed API Limit
type Limit struct {
	int
}

func NewLimit(isPrivate bool) *Limit {
	if isPrivate {
		return &Limit{
			int: APILIMIT,
		}
	}
	// 201904: Private/Public same limit number
	return &Limit{
		int: APILIMIT,
	}
}

// Check is api limit
func (p *Limit) Check() error {
	if p.int <= 0 {
		return errors.New("api limit, too much")
	}

	p.int--
	return nil
}

// ResetAPI is resets api limit at every 5minutes
func (p *Limit) ResetAPI() {
	var mux sync.Mutex
	mux.Lock()
	defer mux.Unlock()

	p.int = APILIMIT
}

// Remain is remain number of API limit
func (p *Limit) Remain() int {
	return p.int
}
