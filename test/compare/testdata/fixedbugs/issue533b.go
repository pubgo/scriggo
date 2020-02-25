// run

package main

import (
	"fmt"
	"math"
)

func Lgamma(x float64) (float64, int) {
	return math.Lgamma(x)
}

func main() {
	lgamma, sign := Lgamma(20)
	fmt.Println(lgamma)
	fmt.Println(sign)
}
