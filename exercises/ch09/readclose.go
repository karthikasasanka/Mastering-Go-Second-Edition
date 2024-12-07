package main

import (
	"fmt"
)

func main() {
	willClose := make(chan string, 10)
	// Write some data to the channel
	willClose <- "-1"
	willClose <- "0"
	willClose <- "2"
	// Empty channel
	<-willClose
	<-willClose
	<-willClose
	close(willClose)
	read := <-willClose
	fmt.Println("--", read, "--0")
}
