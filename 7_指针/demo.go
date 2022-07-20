package main

import "fmt"

func demo1(a, b int) (int, int) {
	return b, a
}

func demo2() {
	var  number=10 
	//获取str的地址
	var address= &number //取变量number的地址，将指针保存到b中
	fmt.Printf("a:%d ptr:%p\n", number, &number) // a:10 ptr:0xc00000a0a8
	fmt.Printf("b:%p type:%T\n", address, address) // b:0xc00000a0a8 type:*int
	fmt.Println(&address)//0xc000006028
}

func demo3() {
  //指针取值
  a := 10
  b := &a // 取变量a的地址，将指针保存到b中
  fmt.Printf("type of b:%T\n", b)
  c := *b // 指针取值（根据指针去内存取值）
  fmt.Printf("type of c:%T\n", c)
  fmt.Printf("value of c:%v\n", c)
}

func demo4(){
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", *b)
	fmt.Println(*b)
	// c := *b // 指针取值（根据指针去内存取值）
	// fmt.Printf("type of c:%T\n", c)
	// fmt.Printf("value of c:%v\n", c)
}

func main() {
	// a, b := 3, 4
	// a, b = demo1(a, b)
	// fmt.Println(a,b)
	// demo3()
	demo4()
}