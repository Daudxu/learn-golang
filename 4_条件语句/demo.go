package main

import (
	"fmt"
	"io/ioutil"
)

func demo1() {
	const fileName = "abc.txt"
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n", contents)
	}
}

func demo2(score int) string{
	g :=""
	switch {
	case score < 60:
    	g = "F"
	default: //default case
     	panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func main() {
	demo1()
	fmt.Println(demo2(61))
	fmt.Println(123)
}