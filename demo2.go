package main

import "fmt"

//全局常量
const a = "hello"
const b = 1
func main() {
    //局部常量
    const c = true
    fmt.Println(a)
    test()
    fmt.Println(c)
}

func test(){
    fmt.Println(b)
}