package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//声明把s和sep隐式地初始化成空字符串
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)

	//用foreach试一试

	s = ""
	seq = ""
	for index, arg := range os.Args[1:] {
		s += (string(index)) + seq + arg
		seq = " "
	}
	//以下方式等价
	//s := ""				只能用在函数内部，而不能用于包变量
	//var s string			依赖于字符串的默认初始化零值机制，被初始化为""
	//var s = ""			用得很少，除非同时声明多个变量
	//var s string = ""		显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了
	fmt.Println(s)

	//使用join
	fmt.Println(strings.Join(os.Args[1:], " "))

	//直接打印
	fmt.Println(os.Args[1:])

}
