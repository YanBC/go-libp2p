package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("My hosts ID is %s\n", h.ID())
	for _, addr := range h.Addrs() {
		fmt.Println(addr.String())
	}

	select {}

}
