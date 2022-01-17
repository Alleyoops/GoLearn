package main

import "fmt"

//p10_函数
//Go语言中函数没有默认参数
func main() {
	hello()
	what("you")
	what(and(3,7))
	what(andMore(1,3,6))
	what(andMore(1,3))
	_,c :=and2(9,9)
	fmt.Println(c)
	fmt.Println(and2(8,8))

	//defer延迟逆序执行
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)//最后执行
	fmt.Println("end")
	//defer的特性，方便处理资源释放问题，如：资源清理、文件关闭、解锁、记录时间等

	//函数可以当做变量赋值，也可以传入其它函数当参数
	a :=and
	fmt.Printf("%T",a)//a的类型是函数func
	fmt.Println(a(1,5))
	test(a)
}
func hello()  {
	fmt.Println("hello")
}
func what(string2 string)  {
	fmt.Println(string2)
}
func and(a int,b int)(c string)  {
	c = fmt.Sprintf("%d",a+b)
	return c
}
//函数接收可变参数
func andMore(a int,b int, d...int)(c string)  {
	e := a+b
	//d是可以遍历的,d是[]int切片
	fmt.Printf("%T\n",d)
	for _, i2 := range d {
		e+=i2
	}
	c = fmt.Sprintf("%d",e)
	return c//c可省略
}

//函数参数同类型可以简写
//函数具有多个返回值
func and2(a,b int)(sum,sum1 int)  {
	sum = a+b
	sum1 = a+b+1
	return sum,sum1
}

//函数进阶之：变量作用域
//定义全局变量
var num = 10
//传入一个函数当参数
func test(f func(int,int)string){
	c:=f(3,5)
	//可以在函数中访问全局变量
	//num:=100//函数内部有同名变量则优先使用内部变量
	fmt.Println(num,c)
}

