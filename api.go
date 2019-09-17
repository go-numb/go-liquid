package liquid

import "sync"

// Limit is managed API Limit
type Limit struct {
	int
}

// Check is api limit
func (p *Limit) Check() bool {
	if p.int <= 0 {
		return false
	}

	p.int--
	return true
}

// ResetAPI is resets api limit at every 5minutes
func (p *Limit) ResetAPI() {
	var mux sync.Mutex
	mux.Lock()
	defer mux.Unlock()

	p.int = APILIMIT
}
