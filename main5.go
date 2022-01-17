package main

import (
	"fmt"
	"sort"
)

//p8_切片Slice
func main() {
	//数组长度是固定的，且长度属于类型的一部分
	//切片底层是数组，是一个拥有相同类型元素的
	//可变长度的序列

	//定义时跟数组不一样，方括号不指定长度
	var a []string
	var b []int
	var c  = []bool{false,true,true}
	fmt.Println(a,b,c)

	//基于数组来定义切片
	var d = [5]int{12,13,14,15,16}//数组
	e :=d[1:4]//切下标为1~3（前闭后开），容量是切的位置到数组的长度
	fmt.Printf("%T\n",e)
	fmt.Println(e)//[]int类型的切片

	//切片再切片
	f :=e[0:2]
	fmt.Println(f)

	//make函数构造切片
	g := make([]int,5,10)//个数为5，容量为10
	h := make([]string,5)//个数为5，容量为5
	fmt.Println(len(g), len(h))//长度都为5
	fmt.Println(cap(g), cap(h))//容量，前者要扩容

	//ps:容量是切片的容量，长度是底层数组的长度
	//容量是切片不扩容情况下能存入元素的个数
	//长度是切片中已有元素的个数，

	var aa []int//声明了，为nil，还未分配内存空间
	if aa==nil{
		fmt.Println(aa,len(aa), cap(aa))
	}
	var aaa  = []int{}//声明并初始化，不是nil,分配了内存
	if aaa!=nil {
		fmt.Println(aaa,len(aaa), cap(aaa))
	}
	aaaa :=make([]int,0)//不为nil
	if aaaa!=nil {
		fmt.Println(aaaa, len(aaaa), cap(aaaa))
	}

	//切片的赋值拷贝
	t := make([]int,3)//[0,0,0]
	tt := t//tt和t共同指向内存里同一个数组
	fmt.Println(t,tt)

	//切片的遍历
	for i, i2 := range e {
		fmt.Println(i,i2)
	}

	//append方法动态添加元素
	//底层数组不够时扩容会更换底层数组
	var s []int//声明
	//s[0] = 100//会报错，因为没有初始化分配内存
	//fmt.Println(s)
	s = append(s,100,200)
	fmt.Println(s)
	for i:=0;i<16 ;i++  {
		s = append(s,i)
		fmt.Println(len(s),cap(s), &s)//这里的容量double增加，按切片扩容机制不止于此
	}
	fmt.Println(s)
	//用切片追加切片,...表示追加e里所有元素
	ss := append(s,e...)
	fmt.Println(ss)

	//数组是值类型，切片是引用类型，直接a赋值给b相当于复制指针地址，修改b也会修改a的值
	//所以用copy，复制值
	var sss[]int
	sss = make([]int,len(ss))
	copy(sss,ss)
	sss[0]=10
	fmt.Println(sss)
	//也就是说，copy是传值，:=或=是传指针引用

	//切片的删除（使用append）
	sli := []string{"beijing","shanghai","tianjin","chongqing"}
	sli = append(sli[0:2],sli[3:]...)//删除tianjin
	fmt.Println(sli)

	//使用内置sort包对数组进行排序
	var array = [...]int{3,7,9,1,4,2}
	sort.Ints(array[:])//传参是一个切片，所以从头切到尾得到切片
	fmt.Println(array)

}
