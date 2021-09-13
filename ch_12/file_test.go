package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("secret.txt")
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("close file fail", err)
		}
	}()
	if err != nil {
		fmt.Println("open file have an error", err)
	}
	reader := bufio.NewReader(file)
	//var results []rune
	var results bytes.Buffer
	i := 0
	for {
		fmt.Println(i)
		//一次读一个utf-8字符
		//r,_,err := reader.ReadRune()
		//results = append(results,r)

		//一次读一行
		readString, err := reader.ReadString('\n')
		results.WriteString(readString)

		if err != nil && err == io.EOF {
			break
		}
		i++
	}
	//finalResult := string(results)
	finalResult := results.String()

	fmt.Println(finalResult)
}

// TestBuffRead 缓冲读取骚操作,抄的 ioutuil.ReadFile,这个例子很有助于理解切片和数组的关系
func TestBuffRead(t *testing.T) {
	//bb,_ := ioutil.ReadFile("bb.txt")
	//fmt.Println(bb)
	//基操,打开文件
	file, err := os.Open("secret.txt")
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("close file fail", err)
		}
	}()
	if err != nil {
		fmt.Println("open file have an error", err)
	}
	reader := bufio.NewReader(file)
	//搞一个buffer去读取文件
	//这个 0 和512就是骚操作,其实size多少都没问题,原来的方法设置了成512有注释,最少要读512
	//我这里为了测试,把size的长度小一点,因为我的文件也才231字节,ioutil.ReadFile方法直接读取整个文件的size+1,+1是为了取到eof
	//0 是buf的len,也就是现在的长度,size 是底层数组的大小
	//size := 512
	size := 10
	var buf = make([]byte, 0, size)
	for {
		//如果cap用满了,也就是说底层数组没有空间了,那就进行下扩容
		if len(buf) >= cap(buf) {
			//这里进行了扩容,d的底层数组长度就变成了原来的2倍
			d := append(buf[:cap(buf)], 0)
			// 这里的将d传给buf,也就是让buf指向的d的底层数组,注意这个len(buf),是因为后面加了个0,所以要保证原来的程度
			buf = d[:len(buf)]
		}
		count, err := reader.Read(buf[len(buf):cap(buf)])
		//使buf的len增长,而且有错误也不怕,因为有错误这里count == 0
		buf = buf[:len(buf)+count]
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	fmt.Println(string(buf))

}
