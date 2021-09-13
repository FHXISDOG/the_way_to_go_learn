package ch_7

import (
	"bytes"
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	i := []int32("你好")
	b := []byte("hello")
	for _, ii := range i {
		fmt.Printf("%d\n", ii)
		fmt.Printf("%c\n", ii)
	}
	fmt.Println("========================")
	for _, bb := range b {
		fmt.Printf("%d\n", bb)
		fmt.Printf("%c\n", bb)
	}
	str := "我赵日天最牛逼"
	fmt.Println("========================")
	for _, v := range str {
		switch interface{}(v).(type) {
		case int:
			fmt.Println("v is int")
		case byte:
			fmt.Println("v is byte")
		case rune:
			fmt.Println("v is rune")
		default:
			fmt.Println("I don't now v's type")
		}
		fmt.Printf("%d\n", v)
	}
	fmt.Println("========================")
	for i := 0; i < len(str); i++ {
		v := str[i]
		switch interface{}(v).(type) {
		case int:
			fmt.Println("v is int")
		case byte:
			fmt.Println("v is byte")
		case rune:
			fmt.Println("v is rune")
		default:
			fmt.Println("I don't now v's type")
		}
		fmt.Printf("%d\n", v)
	}
	fmt.Println(len([]rune(str)))
}

func TestStringJoin(t *testing.T) {
	strs := []string{"我", "赵", "天", "最", "牛", "逼"}
	buffer := bytes.Buffer{}
	for _, v := range strs {
		i := []int32(v)
		fmt.Println("=============")
		fmt.Println(len(i))
		for _, ii := range i {
			fmt.Println(ii)
		}
		fmt.Println("=============")
		fmt.Printf("v's len is %d \n", len(v))
		if count, err := buffer.WriteString(v); err == nil {
			fmt.Printf("have append %d byte \n", count)
		}
	}
	str := buffer.String()
	fmt.Printf("result is %s", str)
}

func TestSlice(t *testing.T) {
	//数组是可以new不可以make
	//b := make([5]int)
	//b[0]  = 1
	//fmt.Println(b)

	//此处要理解,分片操作,操作的是原来的数组
	s := make([]int, 0, 10)

	ss := s[len(s):cap(s)]
	ss[0] = 1
	s = s[0:2]
	fmt.Println(s)
}
