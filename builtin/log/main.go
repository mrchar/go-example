package main

import (
	"log"
	"os"
)

func main() {
	// 执行普通打印
	println()
	// 执行异常打印
	// panicln()
	// 执行致命错误打印
	// fatalln()
	// new 新建logger并打印
	// new()
	// 设置Flags并打印
	// flags()
	// 控制打印的调用深度
	// output()
}

func println() {
	log.Println("这是一条普通打印")
}

func panicln() {
	log.Panicln("这是一条异常打印")
}

func fatalln() {
	log.Fatalln("这是一条致命错误打印")
}

func new() {
	logger := log.New(os.Stdout, "[new]", log.LstdFlags)
	logger.Println("这是一条普通打印")
}

func flags() {
	log.SetPrefix("[prefix] ")
	// log.SetFlags(log.Ldate)
	// log.SetFlags(log.Ltime)
	// log.SetFlags(log.Ldate | log.Ltime)
	// log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	// log.SetFlags(log.Llongfile)
	// log.SetFlags(log.Lshortfile)
	// log.SetFlags(log.Lmsgprefix)
	log.SetFlags(log.LstdFlags)
	// log.SetFlags(log.LstdFlags | log.Llongfile)
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	flags := log.Flags()
	log.Printf("这是一条普通打印, Flags是: %d", flags)
}

func output() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	print := func() {
		log.Output(1, "这是一条普通打印, 调用深度为：1")
		log.Output(2, "这是一条普通打印, 调用深度为：2")
		log.Output(3, "这是一条普通打印, 调用深度为：3")
	}

	print()
}
