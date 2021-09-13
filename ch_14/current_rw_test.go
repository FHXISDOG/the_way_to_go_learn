package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type person struct {
	name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string) *person {
	p := &person{name: name}
	p.chF = make(chan func()) //记得初始化这里,向一个为nil的chan中写入数据直接陷入死锁
	go p.backend()
	return p
}

func (p *person) backend() {
	for f := range p.chF {
		f()
	}
}

func (p *person) addSalary(i int, salary float64) {
	//fmt.Printf("begin %d times add salary %.3f \n",i,salary)
	fmt.Println("call add salary")
	p.chF <- func() {
		//p.setSalary(p.getSalary() + salary) //这样写会报错,因为backend执行这个函数的时候会调用getSalary,getSalary添加一个func到chF,但是chF此时被这个函数使用,所以会等待,陷入死锁
		p.salary = p.salary + salary //这里保证了同时只能有一个函数去访问salary
	}
}

func (p *person) setSalary(salary float64) {
	fmt.Println("call set salary")
	p.chF <- func() {
		p.salary = salary
	}
}

func (p *person) getSalary() float64 {
	fmt.Println("call get salary")
	salaryCh := make(chan float64)
	p.chF <- func() {
		salaryCh <- p.salary
	}
	return <-salaryCh
}
func (p *person) String() string {
	return fmt.Sprintf("%s salary is %.3f $", p.name, p.getSalary())
}

func TestPersonSalary(t *testing.T) {
	p := NewPerson("finger")
	p.setSalary(0)
	fmt.Println(p)
	const addTimes = 10
	var doneChArr [addTimes]chan int
	//并发涨薪
	var totalAdd float64 = 0
	for i := 0; i < addTimes; i++ {
		doneCh := make(chan int)
		doneChArr[i] = doneCh
		addSalary := rand.Float64() * 1000
		totalAdd += addSalary
		go func(i int) {
			p.addSalary(i, addSalary)
			//p.salary += addSalary
			doneCh <- 1
		}(i)
	}
	for _, v := range doneChArr {
		<-v
	}
	fmt.Printf("total add salary is %.3f \n", totalAdd)
	fmt.Println(p)
}

func TestNPE(t *testing.T) {
}
