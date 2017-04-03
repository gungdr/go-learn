package palindrome

import (
	"fmt"
)

func Generate() {
	kodok := "kodok"
	kodoklen := len(kodok)
	for i := 0; i < (kodoklen / 2); i++ {
		if string(kodok[(kodoklen-1)-i]) != string(kodok[i]) {
			fmt.Printf("Not a palindrome")
			return
		}
	}
	fmt.Printf("Is a palindrome")

}
