package main

import "fmt"

var (
    aa = 3
    ss = "kkk"
    bb= true
)

func variableZeroValue() {
    var a int
    var s string
    fmt.Printf("%d %q\n", a, s)
}

func variableInitiaValue() {
    var a , b int = 3,4
    var s string = "abc"
    fmt.Println(a, b, s)
}

func variabTypeDeduction() {
    var a, b, c, s = 3,4, true,"6"
    fmt.Println(a,b,c,s)
}

func varialbeShorter() {
    a, b, c, s := 3,4, true,"6"
//    b = 5,
   fmt.Println(a,b,c,s)
}

func main () {
    variableZeroValue()
    variableInitiaValue()
    variabTypeDeduction()
    varialbeShorter()
    // fmt.Println("hello world")
}