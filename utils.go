package ga

import (
	"fmt"
	"math/rand"
	"os"
)

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func bitMask(num int) uint32 {
	if num >= 32 {
		return 0xffffffff
	}

	return (1 << uint(num)) - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func logf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}
