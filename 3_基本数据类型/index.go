package main

import (
	"bytes"
	"fmt"
	"strings"
)

func demo1() {
    var j int8 = -128
	fmt.Println(j)
	
	var a int8 = 127
	fmt.Println(a)
}

func demo2(){
    // 测试unit8取值范围
	var a uint8 = 255
	fmt.Println(a)
	//a = 256 as uint8 value in assignment (overflows)
	
	var b uint8 = 0
	fmt.Println(b)
	//b = -1 as uint8 value in assignment (overflows)
	fmt.Println(b)

    str := "hello world"
	fmt.Println(len(str)) //获取字符串长度

	fmt.Println(strings.Split(str , " ")) //字符串分割

	fmt.Println(strings.Contains(str , "hello")) //判断是否包含字串

	fmt.Println(strings.HasPrefix(str , "hello")) //判断是否作为前缀

	fmt.Println(strings.HasSuffix(str , "world")) //判断是否作为后缀

	fmt.Println(strings.Index(str , "l")) //字串出现的位置

	fmt.Println(strings.LastIndex(str , "l")) //join拼接操作
}

func demo3(){
    var price float32 = 89.12 
	fmt.Println(price)

	// float32 和 float64都是有符号的
	var num1 float32 = -0.00
	var num2 float64 = -0.00
	fmt.Println("num1= " , num1 , "num2= " , num2)

	// 尾数部分可能丢失，造成精度损失
	// float64 比 float32精度更大
	var num3 float32 = -123.0000901 // num3= -123.00009
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3 , "num4" , num4)
    
    num5 := 89.001
	fmt.Printf("%T", num5) //float64
    
	str := "hello world"
	n := 3
	m := 5

	fmt.Println(str[n])		// 获取字符串索引位置为n的原始字节
	fmt.Println(str[n:m])	// 截取得字符串索引位置为n到m-1的字符串
	fmt.Println(str[n:])	// 截取得字符串索引位置为n到len(str) - 1的字符串
	fmt.Println(str[:m])	// 截取得字符串索引位置为0到m-1的字符串
}

func demo5(){
    var b1 bool = true
    var b2 bool = false
    var b3 = true
    var b4 = false
    b5 := true
    b6 := false

    fmt.Println(b1)
    fmt.Println(b2)
    fmt.Println(b3)
    fmt.Println(b4)
    fmt.Println(b5)
    fmt.Println(b6)
}
func demo4(){
    var buffer bytes.Buffer
	buffer.WriteString("ljy")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Println(buffer.String())
}

func main() {
    demo1()
}
