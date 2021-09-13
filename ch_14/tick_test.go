package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	chRate := time.Tick(1e10 / 10)
	for {
		<-chRate
		fmt.Println("tick")
	}
}
