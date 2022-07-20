package main

import "fmt"

func demo1() {
	var arr [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arr)
}

func demo2(){
	q := [...]int{1, 2, 3}
	fmt.Println(q)
}

func demo3(){
	var arr [4]int = [4]int{1,2,3,4}
	for i, v :=range arr {
		fmt.Printf("数组中的第%v项, 值是%v\n", i, v)
	}
}

func  demo4(){
	var arr [4]int = [4]int{1,2,3,4}
	for i := 0; i<len(arr); i++ {
		fmt.Printf("数组中的第%v项, 值是%v\n", i, arr[i])
	}
}

func demo5(){
	slice := make([]int, 5)
	fmt.Println(slice)
}

func demo6(){
	myStr := []string{"Jack", "Mark", "Nick"}
    myNum := []int{10, 20, 30, 40}
	fmt.Println(myStr)
	fmt.Println(myNum)
	fmt.Println(myNum[2:3])
	fmt.Println("len", len(myNum))
	fmt.Println("cap", cap(myNum))
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
}

