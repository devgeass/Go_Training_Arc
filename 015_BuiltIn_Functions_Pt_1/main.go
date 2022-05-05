package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Abs(-4) = ", math.Abs(-4))
	fmt.Println("Ceil(2.25) = ", math.Ceil(2.25))
	fmt.Println("Floor(2.85) = ", math.Floor(2.85))
	fmt.Println("Log(100) = ", math.Log(100))
	fmt.Println("Log10(100) = ", math.Log10(100))
	fmt.Println("Pow(4,2) = ", math.Pow(4, 2))
	fmt.Println("Pow10(2) = ", math.Pow10(2))
	fmt.Println("Sqrt(25) = ", math.Sqrt(25))
	fmt.Println("Max(25,100) = ", math.Max(25, 100))
	fmt.Println("Min(25,100) = ", math.Min(25, 100))
	fmt.Println("Sin(30) = ", math.Sin(30))
	fmt.Println("Cos(30) = ", math.Cos(30))
	fmt.Println("Tan(30) = ", math.Tan(30))
	fmt.Println("Mod(10,3) = ", math.Mod(10, 3))
	fmt.Println("Remainder(10,3) = ", math.Remainder(10, 3))
	fmt.Println("Exp(2) = ", math.Exp(2))
	fmt.Println("Inf(1) = ", math.Inf(1))
	fmt.Println("Inf(-1) = ", math.Inf(-1))
	fmt.Println("PI = ", math.Pi)
	fmt.Println("PHI = ", math.Phi)
	fmt.Println("e = ", math.E)

	fmt.Println("Rand.Int() = ", rand.Int())
	fmt.Println("Rand.Intn(73) = ", rand.Intn(73))
	fmt.Println("Rand.Float32() = ", rand.Float32())
	fmt.Println("Rand.Float64() = ", rand.Float64())
	fmt.Println("Rand.Perm(10) = ", rand.Perm(10))

}
