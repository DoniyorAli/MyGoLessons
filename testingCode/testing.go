package main

import (
	"fmt"
	"math/big"
)

// Fibonacci using math/big for arbitrary large n
func fibBig(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := 2; i <= n; i++ {
		// tmp = a + b
		tmp := new(big.Int).Add(a, b)
		// shift: a = b, b = tmp
		a.Set(b)
		b.Set(tmp)
	}
	return b
}

func main() {
	n := 500
	fmt.Printf("Fib(%d) = %s (big.Int)\n", n, fibBig(n).String())
}