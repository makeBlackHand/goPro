package main

import "fmt"

func main() {
	var sum, toatl float64
	var arr [2][2]float64
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%d班的第%d个学生的成绩\n", i+1, j+1)
			fmt.Scanln(&arr[i][j])
		}
	}
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}
		toatl += sum
		fmt.Println("每班班总分为：", sum, "平均分", float64(sum)/2.0)
	}
	fmt.Println("全部", toatl, "平均分", float64(toatl)/4.0)
}
