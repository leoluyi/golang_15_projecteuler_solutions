package main

import "fmt"

func main() {
	sum := 0
	count := 1_000
	for i := 0; i < count; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
