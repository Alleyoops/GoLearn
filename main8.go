package main

import "fmt"

//p11_函数进阶
//匿名函数，多用于实现回调函数和闭包
func main() {
	//匿名函数定义，和普通函数区别就是没有函数名
	//方法1，用变量接收函数
	hello := func() {
		fmt.Println("hello")
	}
	hello()
	//方法2，直接执行
	func() {
		fmt.Println("匿名函数")
	}() //匿名函数定义完成加()直接执行，其实就相当于前面的hello加了个()

	//闭包，函数加上外层的应用，r此时就是一个闭包
	r := a(18)
	r() //相当于执行了a函数内部的匿名函数

	//panic
	bb()
	aa()

}

//闭包：一个函数与其相关的引用环境组合而成的实体
//闭包=函数+引用环境
//定义一个函数，它的返回值是函数
func a(age int) func() {
	name := "我"
	return func() {
		fmt.Println("hello people", name, age)
	}
}

//内置函数
//close关闭channel
//len求长度
//new分配内存，主要是分配值类型
//make分配内存，主要分配引用类型
//append追加元素到数组、切片中
//panic、recover用来做错误处理

//panic和recover，有点像try catch
func aa() {
	fmt.Println("panic?")
}
func bb() {
	//defer recover语句要定义在panic之前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%T\n", err)
			fmt.Println(err, "what")
			fmt.Println("func bb error")
		}
	}()
	panic(404) //觉得err的类型
}
