package main

import "fmt"

var memo []uint64

func main() {
	var n uint64 = 30
	memo = make([]uint64, n+1)
	fmt.Println(fibo(n))
}

func fibo(n uint64) uint64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if memo[n] != 0 {
			return memo[n]
		}
		memo[n] = fibo(n-2) + fibo(n-1)
		return memo[n]
	}
}
