package main

import "fmt"

type Good struct {
	name  string
	price float64
}
type Brand struct {
	name    string
	address string
}
type Tv struct {
	Good
	Brand
}

func main() {
	tv := Tv{Good{
		name:  "电视",
		price: 1999.2,
	},
		Brand{
			name:    "小米",
			address: "徐州",
		}}
	fmt.Println(tv)
}
