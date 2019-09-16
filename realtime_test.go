package liquid

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	done := make(chan struct{})
	realtime := NewConnect()
	go realtime.Connect([]string{PUSHERchEXECUTION, PUSHERchASK, PUSHERchBID}, []string{BTCJPY, ETHJPY})

	for {
		select {
		case result := <-realtime.Results:
			fmt.Printf("result: %+v\n", result)
		}
	}

	<-done
}
