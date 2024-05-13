package main

func main() {
label1:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				break label1
			}
			println("j=", j)
		}
	}

}
