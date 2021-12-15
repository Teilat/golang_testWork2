package main

import (
	"fmt"
	"time"
)
import "./cache"

const tickTime = time.Second

func main()  {
	ticker := time.NewTicker(tickTime)
	go cache.ProcessTimer(ticker)
	fmt.Scanln()
}
