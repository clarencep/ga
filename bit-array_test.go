package ga

import (
	"fmt"
	"testing"
)

func TestBitArray1(t *testing.T) {
	x := NewBitArray(100)
	x.FillRandBits()
	t.Logf("%s \n", x.String())

	y := NewBitArray(100)
	y.FillRandBits()
	t.Logf("%s \n", y.String())

	z := x.CrossAt(40, y)
	t.Logf("%s \n", z.String())

	toBits := func(n int) string {
		bits := NewBitArray(32)
		bits.SetBulk(0, n)
		return fmt.Sprintf("%s -- (0x%x)", bits.String(), n)
	}

	t.Logf("%s\n", toBits(z.GetInt(0, -1)))
	t.Logf("%s\n", toBits(z.GetInt(32, -1)))
	for i := 0; i <= 40; i++ {
		t.Logf("%s ----[40:%d]\n", toBits(z.GetInt(40, i)), i)
	}

	t.Fail()
}
