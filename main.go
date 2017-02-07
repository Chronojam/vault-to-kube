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
	config.Address = "https://localhost:8200"
	client, err := vapi.NewClient(config)
	if err != nil {
		panic(err)
	}

	client.SetToken("ab7c85c7-40bc-100b-d3ad-90575c6b192f")
	c := client.Logical()

	calumStop := make(chan struct{})
	PollLoop(c, calumStop, "calum")
}

