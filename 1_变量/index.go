package basics

import "fmt"
//函数体外的全局变量要遵循闭包规则,不可以使用 := 或者 var c
var a string
var b int
var c [] int
var d [5] int
// e := 12  golang具有闭包规则。 := 其实是两步操作, var e int +  e = 12  而 e = 12 是不能在函数体外执行的
func test() {
    a = "hello"
    e := "world"
    var f = "just"
    
    fmt.Println(a)
    fmt.Println(e)
    fmt.Println(f)
}