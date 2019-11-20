/*
function to calcuate square root by newtons method
*/
package main

import (
	"fmt"
)

/* 
Sqrt function approximates the square root of a input
*/
func Sqrt(x float64) float64 {
	var z float64 = x / 2
	count := 0
	oldDifference := 0
	difference := 0
	for ; count < 10; {
		var newz float64
		newz = z - (z * z - x) / (2 * z)
		difference = int(z - newz)
		if (oldDifference == difference ) {
			count++
		} else {
			count = 0
		}
		z = newz
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(-1),
	)
}