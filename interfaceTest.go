package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}
type HeroSlice []Hero

func (h HeroSlice) Len() int {
	return len(h)
}
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}
func (h HeroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	var intSlice = []int{1, 56, 32, 12, -6}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heroslice HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("怪兽-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroslice = append(heroslice, hero)
	}
	for _, i := range heroslice {
		fmt.Println(i)
	}
	sort.Sort(heroslice)
	fmt.Println("排序后:")
	fmt.Println(heroslice)
}
