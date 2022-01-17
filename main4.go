package main

import "fmt"

//p7_数组
func main() {
	var a[3]int
	var b[4]int
	fmt.Println(a,b)
	//初始化,数组的类型包括数组长度
	//方法1
	var city = [4]string{"cq","bj","tj"}
	fmt.Println(city)
	//方法2,不指定长度（去掉三个点是切片）
	var Array = [...]bool{true,false,false}
	fmt.Println(Array)
	fmt.Println(len(Array))
	//方法3，使用索引值初始化
	var test = [...]string{1:"one",4:"four"}
	test[2]="two"
	test[0]="zero"
	fmt.Println(test)
	fmt.Printf("%T\n",test)

	//数组遍历
	for i, i2 := range test {
		//索引和值
		fmt.Println(i,i2)
	}
	for i:= range test {
		//索引
		fmt.Println(i)
	}
	for _, i2 := range test {
		//值
		fmt.Println(i2)
	}

	//多维数组
	//下面是一个长度为3的[2]string类型数组的数组
	aa := [3][2]string{
		{"1","2"},
		{"3","4"},
		{"5","6"},//末尾逗号不能省
	}
	fmt.Println(aa)
	//二维数组遍历
	for _, i2 := range aa {
		fmt.Println(i2)
		for i, _ := range aa {
			fmt.Println(aa[i])
		}
	}

	//数组是值类型,拷贝值，不改变数组的引用
	x :=[3]int{1,2,3}
	f1(x)//发现x未变化
	fmt.Println(x)
}
func f1(a [3]int) {
	a[0] = 100
}
