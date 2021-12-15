package main

import (
	"context"
	"src/golang_testWork2/api"
	"src/golang_testWork2/vault"
	"time"
)

const tickTime = time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(tickTime)
	cache := vault.NewVault(ctx, *ticker)
	go cache.ProcessTimer()
	defer cancel()

	api.Init(cache)
}
