package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input some thing")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input is %s", input)
	}
}
