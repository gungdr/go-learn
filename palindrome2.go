package main

import "fmt"

func main() {

	number := 10001
	num := number
	rev := 0
	for number > 0{
		rem := number % 10
		rev = rev * 10 + rem
		number = number/10
	}

	if num == rev {
		fmt.Println("is a palindrome")
	}else{
		fmt.Println("not a palindrome")
	}
	fmt.Println(num)
	fmt.Println(rev)


	
}