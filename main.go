package main

import (
	"fmt"

	"github.com/tonkeeper/tongo"
)

func main() {
	fmt.Println("Hello, World!")
	// Add your code to use the tongo package here

	config := tongo.Config{
		Liteservers: []tongo.Liteserver{
			{
				IP:   "127.0.0.1",
				Port: 4924,
				Key:  "your-public-key",
			},
		},
	}

	client, err := tongo.NewClient(config)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	fmt.Println("Connected to TON liteserver:", client)
}
