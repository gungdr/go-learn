package reverse

import (
	"fmt"
	"strings"
)

func Generate() {

	input := "Agung Dwi Riansah"
	split := strings.Split(input, " ")

	for i := 0; i < len(split); i++ {
		l := len(split[i])
		if i > 0 {
			fmt.Print(" ")
		}
		for j := 0; j < len(split[i]); j++ {
			if j == 0 {
				fmt.Print(strings.ToUpper(string(split[i][l-1-j])))
			} else if j == len(split[i])-1 {
				fmt.Print(strings.ToLower(string(split[i][l-1-j])))
			} else {
				fmt.Print(string(split[i][l-1-j]))
			}

		}

	}
	fmt.Println()
}
