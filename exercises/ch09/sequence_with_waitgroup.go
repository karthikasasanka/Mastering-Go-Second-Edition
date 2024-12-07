package main

import (
	"fmt"
	"sync"
	"time"
)

func write(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		fmt.Println("started process", i)
		c <- i
	}
	close(c)

}

func process(c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for k := range c {
		fmt.Println(k * k)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go write(c, &wg)
	go process(c, &wg)
	wg.Wait()
	fmt.Println("exiting")
}
