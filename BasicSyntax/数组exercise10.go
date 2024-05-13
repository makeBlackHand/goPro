package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成评委打分的数组
	var arr [8]int
	rand.Seed(time.Now().Unix())
	l := len(arr)
	for i := 0; i < l; i++ {
		arr[i] = rand.Intn(10)
	}
	fmt.Println("8位评分评分:", arr)

	//找出最高分的评委和最低分的评委
	var max, min, m, n int = arr[0], arr[0], 0, 0
	//这里用map更方便,我选择用切片解决
	var sliceMax []int = make([]int, 0) //最高分可能有多个,存储最高分的下标
	var sliceMin []int = make([]int, 0) //同最高分
	for i := 1; i < l; i++ {
		if arr[i] > max {
			max = arr[i]
			m = i
		}
		if arr[i] < min {
			min = arr[i]
			n = i
		}
	}
	for i := 0; i < l; i++ {
		if arr[i] == max {
			sliceMax = append(sliceMax, i) //如果评委的序号从1开始算,就 i+1
		}
		if arr[i] == min {
			sliceMin = append(sliceMin, i)
		}
	}
	fmt.Printf("打最高分的评委是:%v 号评委,评分是:%v\n", sliceMax, max)
	fmt.Printf("打最低分的评委是:%v 号评委,评分是:%v\n", sliceMin, min)

	//求平均值,获取得分
	var sum float64
	for i := 0; i < l; i++ {
		if i == n || i == m {
			continue
		}
		sum += float64(arr[i])
	}
	sum /= (float64(l - 2))
	fmt.Printf("得分:%.1f\n", sum) //四舍五入保留一位小数

	//找出最佳评委和最差评委
	//先找出最佳n1和最差m1的分数
	var m1, n1 float64
	var m2, n2, m3, n3 int = arr[0], arr[0], arr[0], arr[0]
	if float64(arr[0]) > sum {
		m1 = float64(arr[0]) - sum
		n1 = float64(arr[0]) - sum
	} else {
		m1 = sum - float64(arr[0])
		n1 = sum - float64(arr[0])
	}
	for i := 1; i < l; i++ {
		if float64(arr[i]) > sum {
			if float64(arr[i])-sum >= m1 {
				m1 = float64(arr[i]) - sum //获取最差
				m2 = arr[i]                //最差的数
			}
			if float64(arr[i])-sum <= n1 {
				n1 = float64(arr[i]) - sum //获取最佳
				n2 = arr[i]                //最佳的值
			}
		} else {
			if sum-float64(arr[i]) >= m1 {
				m1 = sum - float64(arr[i])
				m3 = arr[i] //减去的元素和被减去的元素不相等,但差可以相等,都算最优或最差
			}
			if sum-float64(arr[i]) <= n1 {
				n1 = sum - float64(arr[i])
				n3 = arr[i]
			}
		}
	}
	sliceMax1 := make([]int, 0) //获取下标的切片
	sliceMin1 := make([]int, 0)
	sliceMax2 := make([]int, 0) //获取评分的切片
	sliceMin2 := make([]int, 0)
	//获取可能评分不相等但差值相等的评分
	for i := 0; i < l; i++ {
		if arr[i] == m2 && m1 == float64(arr[i])-sum {
			sliceMax1 = append(sliceMax1, i)
			sliceMax2 = append(sliceMax2, arr[i])
		} else if arr[i] == m3 && m1 == sum-float64(arr[i]) {
			sliceMax1 = append(sliceMax1, i)
			sliceMax2 = append(sliceMax2, arr[i])
		}
		if arr[i] == n2 && n1 == float64(arr[i])-sum {
			sliceMin1 = append(sliceMin1, i)
			sliceMin2 = append(sliceMin2, arr[i])
		} else if arr[i] == n3 && n1 == sum-float64(arr[i]) {
			sliceMin1 = append(sliceMin1, i)
			sliceMin2 = append(sliceMin2, arr[i])
		}
	}
	fmt.Printf("最佳评委是 %v 号评委,评分:%v\n", sliceMin1, sliceMin2)
	fmt.Printf("最差评委是 %v 号评委,评分:%v\n", sliceMax1, sliceMax2)
}
