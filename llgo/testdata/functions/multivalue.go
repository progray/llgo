package main

func swap(a, b int) (int, int) {
	return b, a
}

func sub(a, b int) int {
	return a - b
}

func printint(a int, extra ...int) {
	println(a)
	for _, b := range extra {
		println("extra:", b)
	}
}

func main() {
	println(sub(swap(1, 2)))
	printint(swap(10, 20))
}
