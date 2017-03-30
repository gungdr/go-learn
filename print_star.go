package main

import (
	"fmt"
)
func test_a(n int){
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j >= i {
				fmt.Print(j)
			} else {
				fmt.Print("*")
			}

		}
		fmt.Println("")
	}
}
func test_b(n int,j int) {
	if n > 0 {		
		test_b(n-1,j)
		if(n<j){
			fmt.Print("*")
		}else{
			fmt.Print(n)
		}
	}
}
func main() {
	n := 5
	for i:=1;i<=n;i++{
		test_b(n,i)		
		fmt.Println("")
	}
	
	test_a(n)
}
