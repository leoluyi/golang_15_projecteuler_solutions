package main

import "fmt"

func main() {
	a1 := 0
	a2 := 1
	sum := 0
	const max = 4_000_000

	for a2 < max {
		// fmt.Println(a2)
		if a2%2 == 0 {
			sum += a2
		}

		oldFab := a1
		a1, a2 = a2, a2+oldFab
	}
	fmt.Printf("Sum: %d\n", sum)
}
