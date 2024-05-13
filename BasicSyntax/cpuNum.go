package main

import "runtime"

func main() {
	num := runtime.NumCPU()
	println(num)
	runtime.GOMAXPROCS(num - 1)
	println(num)
}
