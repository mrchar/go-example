package main

import (
	"flag"
	"fmt"
)

func main() {
	var arg1 = flag.String("arg1", "hello", "第一个参数")
	var arg2 string
	flag.StringVar(&arg2, "arg2", "world", "第二个参数")

	flag.Parse()

	fmt.Printf("arg1: %s, arg2:%s\n", *arg1, arg2)
}
