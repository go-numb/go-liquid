# go-liquid

Liquid by Quoine API version2

## Description

go-liquid is a go client library for [Liquid by Quoine API](https://developers.quoine.com/v2).

## Installation

```
$ go get -u github.com/go-numb/go-liquid
```

## Usage
``` 
package main

import (
 "fmt"
 "github.com/go-numb/go-liquid"
)


func main() {
	client := liquid.New("<tokenID number>", "<secretkey>")

	fiat, err := client.GetFiatAccounts()
	if err != nil {
	    fmt.Println(err)
	}

	fmt.Printf("%v\n", fiat)

	doSomething()
}
```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-liquid/blob/master/LICENSE)