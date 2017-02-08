package main

import (
	"time"
	vapi "github.com/hashicorp/vault/api"
)

var (
	PollInterval = time.Second * 5
	VaultPrefix = "secret/"
)

func main() {
	config := vapi.DefaultConfig()
	config.Address = "https://vault.chronojam.co.uk:8200"
	client, err := vapi.NewClient(config)
	if err != nil {
		panic(err)
	}

	client.SetToken("c33a75d9-810b-bd6b-53a0-d9b8355f504c")
	c := client.Logical()

	calumStop := make(chan struct{})
	PollLoop(c, calumStop, "calum")
}

