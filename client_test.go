package liquid

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ID    string
	Token string
}

const (
	KEYPATH = ""
)

func TestGetFiat(t *testing.T) {
	var c Config
	toml.DecodeFile(KEYPATH, &c)

	client, err := New(c.ID, c.Token, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", client)

	fiat, err := client.GetFiatAccounts()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", fiat)
}
