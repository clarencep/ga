package ga

import (
	"math"
	"testing"
)

func TestBBGa1(t *testing.T) {
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

	decode := func(bits *BitArray) interface{} {
		x := bits.GetInt(0, length)
		return (float64(min) + float64(x)*float64(max-min)) / float64((int(1)<<uint(length))-1)
	}

	ga := NewBasicBitsGa(fitness, decode, length, count)
	for i := 0; i < evolves; i++ {
		ga.Evolve(retainRate, randomSelectRate, mutationRate)

		t.Logf("[%d]:\t %s\n", i, ga.ResultAsString())
	}

	t.Fail()
}
