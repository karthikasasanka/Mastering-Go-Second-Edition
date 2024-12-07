package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args[1:]
	sum := 0.0
	for _, arg := range arguments {
		f, _ := strconv.ParseFloat(arg, 64)
		fmt.Println(f)
		sum = sum + f
	}
	fmt.Println("--------------------------------------")
	fmt.Println(sum)
}
