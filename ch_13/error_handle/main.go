package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func main() {
	//err := fmt.Errorf("hhhh %v",errors.New("ggg"))
	//fmt.Println(err)
	fmt.Println("main begin")
	testCheckUser()
	fmt.Println("main end")
}

func testCheckUser() {
	fmt.Println("begin check test check User")
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s \r\n", e)
		}
	}()
	checkUser()
	fmt.Println("begin check test check User")
}

func checkUser() {
	if user != "fingerfringss" {
		panic("this app only can run by finger")
	}
	fmt.Println("user check pass")
}

func generatePanic() {
	panic("fuck! this is  a panic")
}
