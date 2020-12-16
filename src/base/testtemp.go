package main

func myfunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	myfunction(66, 77)
}
