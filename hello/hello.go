//Package hello is the package - go program is made up of packages
package hello

import ( 
	"fmt"
	"math/rand"
	"math"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swapStrings(s1, s2 string) (string, string) {
	return s2, s1
}

// naked return & also named returns
func split(sum int) (x, y int) {
	x = sum * 4 /9
	y = sum - x
	return
}

var i, j = 1, 2

func main() {
	var i int = 32
	var j float32 = float32(i)
	const k int = 30
	fmt.Println("hello, world")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(math.MaxUint32))
	// fmt.Println(math.Pi)
	// fmt.Println(add(42, 13))
	// fmt.Println(swapStrings("World", "Hello"))
	fmt.Println(split(17))
	fmt.Println(i, j, k)
}

// create functions in go
