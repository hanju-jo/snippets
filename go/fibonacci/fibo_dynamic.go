package main

import "fmt"

func main() {
	var n uint64 = 30
	arr := make([]uint64, n+1)
	fmt.Println(fibo(n, arr))
}

func fibo(n uint64, arr []uint64) uint64 {
	arr[0] = 0
	arr[1] = 1
	
	var i uint64
	for i = 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
