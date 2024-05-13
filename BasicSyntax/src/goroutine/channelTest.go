package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var chans chan interface{}
	chans = make(chan interface{}, 10)
	chans <- 10
	chans <- 20
	cat := Cat{
		Name: "tom",
		Age:  10,
	}
	chans <- cat
	<-chans
	<-chans
	cats := <-chans
	fmt.Println(cats)
}
