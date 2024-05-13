package main

import (
	"fmt"
	"time"
)

func main() {
	times := time.Now() //获取当前时间
	fmt.Println(times)
	fmt.Println("年：", time.Now().Year())
	fmt.Println("月：", int(time.Now().Month()))
	fmt.Println("日：", time.Now().Day())

	//格式化当前日期的方式
	fmt.Printf("当前年月日：%d-%d-%d %d:%d:%d\n", time.Now().Year(), time.Now().Month(),
		time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	time1 := fmt.Sprintf("当前年月日：%d-%d-%d %d:%d:%d\n", time.Now().Year(), time.Now().Month(),
		time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Println(time1)

	fmt.Println(time.Now().Unix())     //unix时间戳
	fmt.Println(time.Now().UnixNano()) //unix纳秒时间戳

	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100) //0.1秒=毫秒*100
		if i == 100 {
			break
		}
	}

}
