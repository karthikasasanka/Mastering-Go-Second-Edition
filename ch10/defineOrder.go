package main

import (
	"fmt"
	"time"
)

func A(x, y chan struct{}) {
	<-x
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(y)
}

func B(y, z chan struct{}) {
	<-y
	fmt.Println("B()!")
	close(z)
}

func C(z chan struct{}) {
	<-z
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
}
