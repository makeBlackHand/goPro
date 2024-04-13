package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Students struct {
	Name  string
	Age   int
	Score float64
}
type StudentSlice []Students

func (s StudentSlice) Len() int {
	return len(s)
}
func (s StudentSlice) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}
func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var s2 StudentSlice
	for i := 0; i < 5; i++ {
		s1 := Students{
			Name:  fmt.Sprintf("学生-%d", rand.Intn(10)),
			Age:   rand.Intn(20),
			Score: float64(rand.Intn(100)) + rand.Float64(),
		}
		s2 = append(s2, s1)
	}

	sort.Sort(s2)
	fmt.Println(s2)

}
