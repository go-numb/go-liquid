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
	KEYPATH = "/Users/<user_name>/.keys/liquid.toml"
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

func TestGetExecutions(t *testing.T) {
	var c Config
	toml.DecodeFile(KEYPATH, &c)

	client, err := New(c.ID, c.Token, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", client)

	ex, err := client.GetExecutions(5, 10, 1)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", ex)
}
