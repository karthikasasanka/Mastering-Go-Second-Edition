package main

import (
	"fmt"
	"time"
)

func write(c chan<- int) {
	for i := range 10 {
		c <- i
		fmt.Println("started process", i)
	}
	close(c)
}
func process(c <-chan int, c2 chan<- bool) {
	for k := range c {
		fmt.Println(k * k)
		time.Sleep(time.Second)
	}
	c2 <- true
}

func main() {
	c := make(chan int)
	c2 := make(chan bool)
	go write(c)
	go process(c, c2)
	// time.Sleep(3 * time.Second)
	<-c2
	fmt.Println("exiting")
}
