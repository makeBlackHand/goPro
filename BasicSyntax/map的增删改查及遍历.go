package main

import "fmt"

func main() {
	city := make(map[string]string)
	city["no1"] = "上海滩"
	city["no2"] = "迪士尼"
	city["no3"] = "东方明珠"
	fmt.Println(city)
	delete(city, "no1") //删除map中的key及其值，即使不存在也不报错
	fmt.Println(city)
	val, or := city["no2"] //val是找到了key的值返回给val，or是判断是否有key返回true或者false
	if or {
		fmt.Println("找到了", val)
	} else {
		fmt.Println("没找到")
	}
	city = make(map[string]string) //清除map所有东西
	//map清除所有新方法：clear(map)
	fmt.Println(city)

	name := make(map[string]string)
	name["no1"] = "张三"
	name["no2"] = "李四"
	name["no3"] = "王五"
	for s, s2 := range name { //map的遍历必须用for-range,因为map循环的内容不一定是数字
		fmt.Printf("s=%v,s2=%v\n", s, s2)
	}

	studentmap := make(map[string]map[string]string)
	studentmap["no1"] = make(map[string]string, 3)
	studentmap["no1"]["sex"] = "男"
	studentmap["no1"]["name"] = "张三"
	studentmap["no1"]["address"] = "松花江"
	studentmap["no2"] = make(map[string]string, 2)
	studentmap["no2"]["sex"] = "女"
	studentmap["no2"]["name"] = "孙尚香"
	fmt.Println(studentmap)
	fmt.Println("studentmap的value有", len(studentmap))               //计算studentmap里的value个数
	fmt.Println("studentmap[“no1”]的value有", len(studentmap["no1"])) //计算studentmap["no1"]里的value个数
	for s, m := range studentmap {                                  //s取的是map里的key，m取的是key后的值
		fmt.Println("s=", s)
		for s2, s3 := range m { //s2取的是学号里的key，m取的是key后的值
			fmt.Printf("s2=%v,s3=%v\n", s2, s3)
		}
	}
}
