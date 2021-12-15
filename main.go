package main

import (
	"context"
	"time"
)
import "./vault"

const tickTime = time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(tickTime)
	cache := vault.NewVault(ctx, *ticker)
	defer cancel()

}
