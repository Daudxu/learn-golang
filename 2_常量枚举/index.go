package main

import "fmt"

func demo1() {
	//使用const来定义枚举类型
	const (
		MON = 1
		TUE = 2
		WEN = 3
		TUR = 4
		FRI = 5
		SAT = 6
		SUN = 7
	)
	fmt.Println(MON)

}

func demo2() {
	//在const()中添加一个关键字iota，每行的iota会加1，第一行默认为0
	const (
		MON = iota
		TUE
		WEN
		TUR
		FRI
		SAT
		SUN
	)
	fmt.Println(MON)
}

func main() {
	demo1()
	demo2()
}
