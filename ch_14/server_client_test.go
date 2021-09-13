package main

import (
	"fmt"
	"testing"
)

type opFunc func(int, int) int

type request struct {
	a   int
	b   int
	res chan int
}

func (req request) String() string {
	return fmt.Sprintf("a is :%d,b is:%d", req.a, req.b)
}

func run(op opFunc, req *request) {
	r := op(req.a, req.b)
	req.res <- r
}

func server(op opFunc, req chan *request, quit chan bool) {
	for { //这里的for不能少,select只遍历一次,不加for,server就退出了造成死锁
		select {
		case reqChan := <-req:
			fmt.Println("I'm going to run ", reqChan)
			go run(op, reqChan)
		case <-quit:
			return
		default:
			//fmt.Println("oh fuck ,no channel can use")
		}
	}
}

func startServer(op opFunc) (reqChan chan *request, quitChan chan bool) {
	reqChan = make(chan *request)
	quitChan = make(chan bool)
	go server(op, reqChan, quitChan)
	return
}

const N = 100

func TestServer(t *testing.T) {
	reqChan, quit := startServer(func(a, b int) int { return a + b })
	var reqArr [N]request
	for i := N - 1; i >= 0; i-- {
		req := &reqArr[i]
		req.a = i
		req.b = i + N
		req.res = make(chan int)
		reqChan <- req
	}
	quit <- true
	// test result
	for i := 0; i < N; i++ {
		req := &reqArr[i]
		if res := <-req.res; res == 2*i+N {
			fmt.Println("this result is ok")
		} else {
			fmt.Println("this result is fault")
		}
	}
	fmt.Println("done!")
}
