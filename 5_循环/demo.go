package main

import (
	"bufio"
	"fmt"
	"os"
)

func convertToBin() string {
	for {
		fmt.Println("abc")
	}
}

func demo1(){
	for n := 1; n <= 3; n++ {
		fmt.Println("n")
	}
}
func demo2(){
	var count int = 10 
	for i := 0; i < count; i++ {
		fmt.Println("a")
	}
}

func demo3(){
	file, err :=os.Open("a.txt")
	if(err != nil) {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println((scanner.Text()))
	}
}

func main() {
	// convertToBin()
	demo1()
	demo2()
	demo3()
	fmt.Println((123))
}