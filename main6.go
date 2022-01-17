package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

//p9_map映射
func main() {
	//和切片一样，是引用类型，需要初始化
	//声明
	var a map[string]int//key,value
	fmt.Println(a)//a为nil
	//初始化
	a = make(map[string]int,8)
	fmt.Println(a)

	//添加键值对
	a["号码"] = 120
	fmt.Printf("%T %v %#v",a,a,a)

	//声明并初始化
	b := map[int]string{
		1:"hh",
		2:"pp",
	}
	fmt.Println(b)

	//判断某个键是否存在
	var score = make(map[string]int,8)
	score["h"] = 100
	score["p"] = 99
	v,ok :=score["x"]
	if ok{
		fmt.Println("have",v)
	}else {
		fmt.Println("not have",v)
	}

	//遍历
	for i, i2 := range score {
		//注意，添加键值对的顺序和遍历的顺序是无关的
		fmt.Println(i,i2)
	}

	//删除
	delete(score,"h")
	fmt.Println(score)

	//按照某个顺序遍历
	var test = make(map[string]int,100)
	for i:=0;i<50 ;i++  {
		key := fmt.Sprintf("stu%02d",i)
		value :=rand.Intn(100)
		test[key] = value
	}
	//打印是有序的
	fmt.Println(test)
	for i, i2 := range test {
		//发现遍历是无序的
		fmt.Println(i,i2)
	}
	//按从小到大遍历
	//先取出key，再排序key，再遍历
	keys:=make([]string,0,100)//切片
	for i := range test {
		keys = append(keys,i)
	}
	sort.Strings(keys)
	for _, i2 := range keys {
		fmt.Println(i2,test[i2])
	}

	//元素类型为map的slice
	var mapSlice = make([]map[string]int,8,8)//完成切片初始化
	//里面的map也需要初始化
	mapSlice[0] = make(map[string]int,8)//完成第一个map初始化
	mapSlice[0]["hp"] = 99
	mapSlice[0]["xyq"] = 100
	fmt.Println(mapSlice)

	//值为切片的map
	var sliceMap = make(map[string][]int,8)

	//里外都要初始化
	sliceMap["China"] = make([]int,8)
	sliceMap["China"][0] = 100
	sliceMap["China"][1] = 90
	fmt.Println(sliceMap)

	//统计单词出现次数
	var s = "hello thank you sir thank you too"
	var wordCount = make(map[string]int,10)
	words:=strings.Split(s," ")
	for _, word := range words {
		v,ok:=wordCount[word]
		if ok{
			wordCount[word] +=v
		}else {
			wordCount[word] = 1
		}
	}
	fmt.Println(wordCount)

}
