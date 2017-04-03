package binary

import (
	"fmt"
)

func Generate() {
	var input int
	_, err := fmt.Scanf("%d", &input)
	var arr []int
	if err == nil {
		//generate
		for input > 0 {
			rem := input % 2
			input = (input - rem) / 2
			arr = append(arr, rem)
		}
		//print
		for i := len(arr) - 1; i >= 0; i-- {
			fmt.Print(arr[i])
		}
		fmt.Println()
	}

}
