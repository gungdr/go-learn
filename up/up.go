package up

import (
	"fmt"
)

func up(n int, u int, d int) {
	p := 0
	i := 0
	for p < n {
		p = p + u
		i++
		if p >= n {
			fmt.Println(i)
			break
		}
		p = p - d
		i++
	}

}

func Generate() {
	up(7, 2, 1)
	fmt.Println()
	up(10, 5, 2)
}
