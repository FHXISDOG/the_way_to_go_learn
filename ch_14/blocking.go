package main

import (
	"fmt"
)

func main() {
	right4()
}

var ch = make(chan string, 2)
var intCh = make(chan int)

func right1() {
	go receiveData(ch)
	sendData(ch)
	fmt.Println("end right 1")
}

func err1() {
	sendData(ch)
	getData3(ch)
	fmt.Println("I'm stop!")
}

func right2() {
	go receiveData(ch)
	go sendData(ch)
	fmt.Println("end right 2")
}

func right3() {
	go sendData2()
	getData2()
	fmt.Println("I'm doing something orther")
	fmt.Println("I'm stop!")
}

func right4() {
	go sendData(ch)
	getDataForRange(ch)
	fmt.Println("I'm stop!")
}
func sendDataClose(ch chan string) {
	defer fmt.Println("stop send")
	ch <- "fuck"
	ch <- "u"
	ch <- "golang"
	ch <- "hahahah"
	close(ch)
}
func sendData(ch chan string) {
	defer fmt.Println("stop send")
	ch <- "fuck"
	ch <- "u"
	ch <- "golang"
	ch <- "hahahah"
}

func receiveData(ch chan string) {
	var res string
	for {
		res = <-ch
		fmt.Println(res)
	}
}

func sendData2() {
	intCh <- 1
}

func getData2() {
	fmt.Println(<-intCh)
	fmt.Println("end get data")
}

func getData3(ch chan string) {
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}

func getDataForRange(ch chan string) {
	for i := range ch {
		fmt.Println(i)
	}
}
