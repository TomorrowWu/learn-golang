package main

import (
	"flag"
	"fmt"
)

var (
	i int
	b bool
)

func init() {
	flag.IntVar(&i, "i", 0, "a int value")
	flag.BoolVar(&b, "b", true, "a bool value")

	flag.Parse()
}

func main() {
	fmt.Printf("i=%d b=%t", i, b)
}
