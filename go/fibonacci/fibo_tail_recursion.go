package main

import "fmt"

func main() {
	var n uint64 = 30
	fmt.Println(fibo(n))
}

func fibo(n uint64) uint64 {
	return fiboTail(n, 0, 1)
}

func fiboTail(n uint64, first uint64, second uint64) uint64 {
	switch n {
	case 0:
		return first
	case 1:
		return second
	default:
		return fiboTail(n-1, second, first+second)
	}
}
