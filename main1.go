package main
//p4_基本数据类型
import (
	"fmt"
	"math"
	"strings"
)

func main()  {
	//int uint取决于操作系统位数

	n := 10
	//进制
	fmt.Printf("%b \n",n)
	fmt.Printf("%d \n",n)

	m := 075//八进制以0开头，%o输出
	fmt.Printf("%o \n",m)
	fmt.Printf("%0x \n",m)

	var age uint8//无符号型
	age = 255//max
	fmt.Println(age)

	//浮点数范围
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	//bool默认为false
	fmt.Println(true,"2333")

	//字符串，utf-8
	s1 := "HEAT"
	fmt.Println(s1)
	//转义符\
	fmt.Println("\\-\\")
	//多行字符串,原样输出，包括空格tab等
	s2 := `qwe\n
		\n
	test def
		adidas\\nike
	`
	fmt.Println(s2)
	s3 := "较好的"//utf-8中文为3个字节
	fmt.Println(len(s1),len(s3))

	//合并字符串
	fmt.Println(s1+s3)
	s4 := fmt.Sprintf("%s-%s",s1,s3)
	fmt.Println(s4)

	//切割字符串,返回字符串切片
	fmt.Println(strings.Split(s4,"-"))
	//%T打印类型
	fmt.Printf("%T\n",strings.Split(s4,"-"))

	//判断是否包含
	fmt.Println(strings.Contains(s4,"好"))
	//判断前缀后缀
	fmt.Println(strings.HasPrefix(s4,"好"))
	//子串位置
	//下面输出为6而不是2
	fmt.Println(strings.LastIndex(s3,"的"))
	//join
	s5 := []string{"how","do","you","do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5,"+"))

	//字符
	//Go的字符有两种：
	//1、uint8，或叫byte，代表一个ASCII字符；
	//2、rune，int32，代表一个utf-8字符
	var c uint8 = 'g'
	var d = 'f'
	e := 'k'
	fmt.Printf("%c,%d,%T",c,d,e)

	//下面两种的区别
	s := "english中文"
	for i:=0;i<len(s) ;i++  {
		fmt.Printf("%d:%c\n",i,s[i])
	}
	for _, r := range s {
		fmt.Printf("%c\n",r)
	}
}
