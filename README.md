# go-liquid

Liquid by Quoine API version2

## Description

go-liquid is a go client library for [Liquid by Quoine API](https://developers.quoine.com/v2).

## Installation

```
$ go get -u github.com/go-numb/go-liquid
```

## Usage
``` golang
package main

import (
 "fmt"
 "github.com/go-numb/go-liquid"
)


func main() {
	client := liquid.New("<tokenID number>", "<secretkey>", nil)

	fiat, err := client.GetFiatAccounts()
	if err != nil {
	    fmt.Println(err)
	}

	fmt.Printf("%v\n", fiat)

	doSomething()
}
```

## Usage for websocket(pusher)
``` golang
package main

import (
	"fmt"
	"github.com/go-numb/go-liquid"
)


func Method(t *testing.T) {
	done := make(chan struct{})

	// for recive response data
	realtime := NewConnect()
	// Connect is reading channels in for loop
	go realtime.Connect([]string{PUSHERchEXECUTION, PUSHERchASK, PUSHERchBID}, []string{BTCJPY, ETHJPY})

	// use result data
	// data in struct
	for {
		select {
		case result := <-realtime.Results:
			fmt.Printf("result: %+v\n", result)
		}
	}

	<-done
}
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-liquid/blob/master/LICENSE)