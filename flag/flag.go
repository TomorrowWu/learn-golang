package main

import (
	"flag"
	"fmt"
)

var (
	i      string
	isTest bool
)

// ./flag -i=1,2,3,4 -b=false

func init() {
	flag.StringVar(&i, "i", "", "a string value")
	flag.BoolVar(&isTest, "istest", true, "a bool value")

	flag.Parse()
}

func main() {
	fmt.Printf("i=%s b=%t", i, isTest)
}
