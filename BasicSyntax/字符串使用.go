package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello哈"
	fmt.Println("len(hello)=", len(str))
	str2 := "xixi我是老北京"
	s := []rune(str2)
	for i := 0; i < len(s); i++ {
		fmt.Printf("str[%d]=%c\n", i, s[i])
	}

	n, err := strconv.Atoi("123") //字符串转整数
	if err != nil {
		println("转换失败", err)
	} else {
		fmt.Println("转换成功", n)
	}

	str = strconv.Itoa(123) //整数转字符串
	fmt.Printf("str=%v,str type is %T\n", str, str)

	var bytes = []byte("hello") //字符串转byte
	fmt.Printf("byte=%v\n", bytes)

	str = string([]byte{97, 98, 99}) //整数转byte
	fmt.Println(str)

	str1 := strconv.FormatInt(123, 2) //十进制转二进制(其他进制转换也是如此)
	fmt.Println(str1)

	str3 := strings.Contains("helloorld", "world") //查找字符串中是否有子串
	fmt.Printf("str2=%v\n", str3)

	num := strings.Count("hello", "LL") //统计字符串中出现几次字符
	println(num)

	str3 = strings.EqualFold("hello", "Hello") //不区分大小写
	println(str3)

	fmt.Println("结果为:", "abc" == "Abc") //区分大小写

	println(strings.Index("aaaabc", "cbc")) //找到子串第一次出现在的index值,找不到返回-1

	println(strings.LastIndex("hhhhhhhh", "h")) //找到子串最后一次出现的index值

	println(strings.Replace("my my name is hh", "my", "ME", 1))
	//替换字符串中的字符 -1代表全部替换，  s:传进变量也可以str:=my my name is hh 把s传进去 str本身不会变化

	str4 := strings.Split("hello,my name is,jack", ",")
	//把一个字符串拆分成多个字符串
	fmt.Printf("str4=%v\n", str4)

	println(strings.ToLower("HHHHHHHHHHH")) //把字符串转换成小写
	println(strings.ToUpper("jhhhhhhhhh"))  //把字符串转换成大写

	println(("   去黑夜   "))
	println(strings.TrimSpace("  去黑夜  ")) //去除字符串两边的空格

	println(strings.Trim("Hajimi", "Hi")) //去除字符串左右两边的指定字符
	//TrimLeft去除左边字符
	//TrimRight去除右边字符
	println(strings.HasPrefix("ssawe1", "ss")) //判断是否以指定字符开头返回true和false

	println(strings.HasSuffix("ssawe1", "ss")) //判断是否以指定字符结尾返回true和false
}
