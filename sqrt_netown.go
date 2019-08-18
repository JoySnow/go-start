package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// For the calculation:
	//  https://blog.csdn.net/WitsMakeMen/article/details/81061217
	z := 1.0
	dit := 1.0
	// 1e-1 requires       50 times loop;
	// 1e-2 requires     5000 times loop;
	// 1e-3 requires   500000 times loop;
	// 1e-4 requires 50000000 times loop.
	min_diff := 1e-2
	i := 0
	for ; math.Abs(dit) > min_diff; i++ {
	//for ; i<10; i++ {
		dit = (z*z - x) / 2 * z
		z -= dit
		//fmt.Println(z, dit, math.Abs(dit))
	}
	fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
