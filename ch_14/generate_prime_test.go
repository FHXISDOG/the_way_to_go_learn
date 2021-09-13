package main

import (
	"fmt"
	"testing"
)

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 1000; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func service() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

//埃拉托色尼素数筛选法
func TestGeneratePrime(t *testing.T) {
	out := service()
	for {
		fmt.Print(<-out, " ")
	}
}
