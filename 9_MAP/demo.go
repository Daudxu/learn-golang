package main

import "fmt"

func demo1() {
	var m map[string]int
	m = make(map[string]int)
	m["张三"] = 10
	m["李四"] = 20
	fmt.Println(m)
	fmt.Println(m["李四"])
}

func demo2() {
	m := make(map[string]int)
    m["张三"] = 10
    m["李四"] = 20
	for key,val := range m {
        fmt.Println(key, val)
    }
 
    for key := range m {
        fmt.Println(key)
    }

}

func demo3() {
	m := make(map[string]int)
    m["小明"] = 50
    m["张三"] = 10
    m["李四"] = 20
    for key,val := range m {
        fmt.Println(key, val)
    }
     
    delete(m,"张三")
    fmt.Println("删除后的map:")
    for key,val := range m {
        fmt.Println(key, val)
    }
}

func demo4(){
	m:=make(map[string][]string)
	fmt.Println(m)
	val := []string {"北京,上海"}
	m["中国"] = val
	fmt.Println(m)
	m["中国"] = append(m["中国"], "广州", "深圳")
	fmt.Println(m)
}

func demo5(){
	var mapSlice = make([]map[string]string, 3)
    for index, value := range mapSlice {
        fmt.Printf("index:%d value:%v\n", index, value)
    }
    fmt.Println("初始化元素：")
	
    mapSlice[0] = make(map[string]string)
    mapSlice[0]["name"] = "小明"
    mapSlice[0]["password"] = "123456"
    mapSlice[0]["address"] = "TBD云集中心"
    for index, value := range mapSlice {
        fmt.Printf("index:%d value:%v\n", index, value)
    }
}

func demo6(){
	
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}