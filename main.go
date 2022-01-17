package main

import "fmt"
//p3_变量和常量

//全局变量
var x = 100//var可省略
var y = "test"
func main() {
	var name string
	var age int
	var isOK bool
	fmt.Println("hello naZha",name,age,isOK)
	//批量声明
	var (
		a string
		b bool
		c int
		d float32
	)
	fmt.Println(a,b,c,d)
	var name1 = "小王子"//类型推导
	fmt.Println(name1)
	//短变量声明。定义在函数内部
	m :=10
	fmt.Println(m)
	fmt.Println(x,y)
	//常量
	const n  = 3.14
	const (
		n1 = 10
		n2 = 20
		n3 //n3等于n2
	)
	fmt.Println(n,n1,n2,n3)
	//iota，常量计数器，只能用于常量表达式
	//使用iota能简化定义，定义枚举时很有用
	const (
		n4 = iota//0
		_//跳过这个值
		n5//2
		n6 //n3等于n2
		n7 = 17
		n8
		n9 = iota//6
	)
	const n10  = iota//0
	fmt.Println(n4,n5,n6,n7,n8,n9)
	fmt.Println(n10)
	//用iota定义数量级
	const (
		_ = iota
		KB = 1<<(10*iota)//1左移10位
		MB = 1<<(10*iota)//1左移20位
		GB = 1<<(10*iota)//1左移30位
	)
	fmt.Println(KB,MB,GB)
	//iota是每新增一行，值加一
	const (
		a1, a2 = iota+1,iota+2//iota=0
		a3, a4 = iota+1,iota+2//iota=1
	)
	fmt.Println(a1,a2,a3,a4)
}