package main

import "fmt"
func main(){
    a := "hesitate"
    for _,char := range a{
        fmt.Printf("%c",char)
        fmt.Printf("--")
    }

    fmt.Println()

    fmt.Printf(a[2:])

    fmt.Println()

    b:= "hello"
    c:= "world"
    b+=c
    fmt.Printf(b)

    fmt.Println(len(a))
}

func consts() {
    const (
        filename = "abc"
    )

    fmt.Println(filename)
}