package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// errorf()
	// print()
	printf()
	// fprint()
	// sprint()
	// scan()
	// sscan()
	// fscan()
}

func errorf() {
	err := fmt.Errorf("这是一条错误, 使用%s创建", "fmt.Errorf")
	fmt.Println(err)
}

func print() {
	fmt.Print("这是一条普通打印\n")
	fmt.Printf("这是%d条%s打印\n", 1, "格式化")
	fmt.Println("这是一条普通打印")
}

func printf() {
	fmt.Printf("普通格式化打印\n")

	fmt.Printf("默认格式打印：%v\n", struct {
		A string
		B int
	}{A: "Hello", B: 1})

	fmt.Printf("带字段名的默认格式打印：%+v\n", struct {
		A string
		B int
	}{A: "Hello", B: 1})

	fmt.Printf("带类型和字段名的默认格式打印：%#v\n", struct {
		A string
		B int
	}{A: "Hello", B: 1})

	fmt.Printf("打印数据类型： %T\n", struct {
		A string
		B int
	}{A: "Hello", B: 1})

	fmt.Printf("打印%%： %%\n")

	fmt.Printf("布尔类型：%t\n", true)

	fmt.Printf("二进制：%b\n", 99)
	fmt.Printf("Unicode对应的字符:%c\n", 99)
	fmt.Printf("十进制:%d\n", 99)
	fmt.Printf("八进制：%o\n", 99)
	fmt.Printf("八进制：%O\n", 99)
	fmt.Printf("单引号包裹的字符：%q\n", 99)
	fmt.Printf("十六进制：%x\n", 110)
	fmt.Printf("十六进制：%X\n", 110)
	fmt.Printf("Unicode格式：%U\n", 99)

}

func fprint() {
	out, err := os.OpenFile("out.log", os.O_CREATE, 0655)
	if err != nil {
		fmt.Println("打开out.log时发生错误", err)
	}
	defer out.Close()

	fmt.Fprint(out, "这是一条普通打印\n")
	fmt.Fprintf(out, "这是第%d条打印\n", 2)
	fmt.Fprintln(out, "这是第三条打印")
}

func sprint() {
	str1 := fmt.Sprint("这是一条普通打印\n")
	str2 := fmt.Sprintf("这是%d条普通打印\n", 1)
	str3 := fmt.Sprintln("这是一条普通打印")

	fmt.Print(str1, str2, str3)
}

func scan() {
	var (
		a string
		b int
		c float32
	)

	count, err := fmt.Scan(&a, &b, &c)
	// count, err := fmt.Scanf("%s %d %f", &a, &b, &c)
	// count, err := fmt.Scanln(&a, &b, &c)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("共获取到%d个参数, a: %s, b: %d, c: %f", count, a, b, c)
}

func sscan() {
	var (
		a string
		b int
		c float32
	)

	count, err := fmt.Sscan("hello 1 3.14", &a, &b, &c)
	// count, err := fmt.Sscanf("hello 1 3.14", "%s %d %f", &a, &b, &c)
	// count, err := fmt.Sscanln("hello 1 3.14", &a, &b, &c)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("共获取到%d个参数, a: %s, b: %d, c: %f", count, a, b, c)
}

func fscan() {
	var (
		a string
		b int
		c float32
	)

	r := strings.NewReader("hello 1 3.14")

	count, err := fmt.Fscan(r, &a, &b, &c)
	// count, err := fmt.Fscanf(r, "%s %d %f", &a, &b, &c)
	// count, err := fmt.Fscanln(r, &a, &b, &c)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("共获取到%d个参数, a: %s, b: %d, c: %f", count, a, b, c)
}
