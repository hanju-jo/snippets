package main

import "fmt"

func main() {
	var n uint64 = 30
	fmt.Println(fibo(n))
}

func fibo(n uint64) uint64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibo(n-1) + fibo(n-2)

	}
}
