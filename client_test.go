package liquid

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ID        string
	SecretKey string
}

const (
	KEYPATH = ""
)

func TestGetFiat(t *testing.T) {
	var c Config
	toml.DecodeFile(KEYPATH, &c)

	client, err := New(c.ID, c.SecretKey, nil)

	fiat, err := client.GetFiatAccounts()
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", fiat)
}
