package main

import "fmt"

//p6_流程控制：if else，for，switch，goto（不推荐）
func main() {
	//第一个花括号必须跟关键字在一行
	var score = 65
	if score>=90 {
		fmt.Println("A")
	}else if score>=75 {
		fmt.Println("B")
	}else if score>=60 {
		fmt.Println("C")
	}else {
		fmt.Println("D")
	}

	//另一种写法,score1只在语句内能够被使用
	if score1 := 95;score1>=90 {
		fmt.Println("A")
	}else if score1>=75 {
		fmt.Println("B")
	}else if score1>=60 {
		fmt.Println("C")
	}else {
		fmt.Println("D")
	}

	//for
	//1：基本写法
	for i:=0;i<10 ;i++  {
		fmt.Println(i)
	}
	//2：省略初始语句，不能省略分号
	var i  = 0
	for ;i<10;i++{
		fmt.Printf("%d",i)
	}
	//3：省略初始语句和结束语句
	for i>5 {
		//这里的i是前面i++后的结果
		fmt.Printf("%d",i)
		i--
	}
	//4：死循环
	//for{
	//	fmt.Println(1)
	//}
	//5：break
	for i:=0;i<10 ;i++  {
		if i>5 {
			break
		}
		fmt.Println("break")
	}
	//6：continue，结束本次循环，继续下一次循环
	for i:=0;i<3 ;i++  {
		if i==1 {
			continue
		}
		fmt.Println("continue")
	}

	//for range,遍历数组、切片、字符串、map、通道channel
	//数组、切片、字符串返回索引和值
	//map返回键和值
	//通道channel只返回通道内的值

	//switch
	v := 3
	switch v {
	case 1:fmt.Println(1)
	case 2:fmt.Println(2)
	case 3:fmt.Println(3)
	default:
		fmt.Println(0)
	}
	//case一次判断多个值
	num := 5
	switch num {
	case 1,3,5,7:
		fmt.Println("奇数")
	default:
		fmt.Println("偶数")
	}
	//case中使用表达式,省略switch后面的表达式
	age :=30
	switch {
	case age>18:
		fmt.Println("yes")
	case age<18:
		fmt.Println("no")
	}
}
