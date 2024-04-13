package main

func test3() {
	var j int
	for i := 2; i <= 100; i++ {
		for j = 2; j <= 100; j++ {
			if i%j == 0 {
				break
			}
		}
		if j >= i {
			println(i)
		}
	}
}

func main() {
	test3()
}
