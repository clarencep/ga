package main

import (
	"clarencep/ga"
	"fmt"
	"math"
)

func main() {
	min := 0
	max := 9
	length := 17
	count := 300
	evolves := 100
	retainRate := 0.2
	randomSelectRate := 0.5
	mutationRate := 0.01

	fitness := func(v interface{}) float64 {
		x, _ := v.(float64)
		return x + 10*math.Sin(5*x) + 7*math.Cos(4*x)
	}

	decode := func(bits *ga.BitArray) interface{} {
		x := bits.GetInt(0, length)
		return (float64(min) + float64(x)*float64(max-min)) / float64((int(1)<<uint(length))-1)
	}

	// fmt.Printf("initializing ga...\n")
	g := ga.NewBasicBitsGa(fitness, decode, length, count)
	for i := 0; i < evolves; i++ {
		// fmt.Printf("[%d]:\t processing...\n", i)

		g.Evolve(retainRate, randomSelectRate, mutationRate)

		fmt.Printf("[%d]:\t result: %s\n", i, g.ResultAsString())
	}

	fmt.Printf("==== final result ====:\n")
	result := g.Result()
	fmt.Printf("Fitness: %v\n", result.Fitness)
	fmt.Printf("Gene bits: %v\n", result.Bits)
	fmt.Printf("Decoded gene: %v\n", result.Decoded)

}

// ============ OUTPUT: ==============
// ...
// [89]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [90]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [91]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [92]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [93]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [94]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [95]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [96]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [97]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [98]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// [99]:    result: 24.855363 --  X.X.XXXX .XXXXX.X X -- 7.856726507007652
// ==== final result ====:
// Fitness: 24.85536281266667
// Gene bits:  X.X.XXXX .XXXXX.X X
// Decoded gene: 7.856726507007652
