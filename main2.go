package main

import "fmt"

//p5_运算符
func main() {
	//算数运算符
	var a = 10
	var b = 20
	fmt.Println(a+b)

	//||是逻辑运算符：或
	//|是位运算符：参与运算的两数各自对应的二进位相或
	c := 1// 001
	d := 5// 101
	fmt.Printf("%b \n",c|d)
	fmt.Printf("%b \n",d<<2)//乘以2的n次方
	fmt.Printf("%b \n",d>>2)//除以2的n次方

	//赋值
	a+=10
	fmt.Println(a)


}
