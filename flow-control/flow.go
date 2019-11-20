package main

import "fmt"

// for i := 0; i < 10; i++
// for {} is an infinite loop

func main() {
	var sum int = 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)
}