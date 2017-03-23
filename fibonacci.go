package main

import (
	"fmt"
)

func main() {
	n:=35
	for i:=0;i<=n;i++{
		x := fib(i)
		fmt.Print(x)
		fmt.Print(" ")
	}
	
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	f := fib(n-1) + fib(n-2)
	return f

}
