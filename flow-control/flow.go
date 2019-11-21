package main

import "fmt"

import "math"

// for i := 0; i < 10; i++
// for {} is an infinite loop

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	var sum int = 1
	for i:=0 ; i < 100; i++ {
		if i % 2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	main2()
}