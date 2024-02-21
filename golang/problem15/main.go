package main

import (
	"fmt"
	"math/big"
)

func factorial(x big.Int) *big.Int {
	n := new(big.Int).Set(&x)

	one := big.NewInt(1)

	if n.Cmp(one) <= 0 {
		return big.NewInt(1)
	}

	result := new(big.Int).Set(n)

	n.Sub(n, one)

	return result.Mul(result, factorial(*n))
}

func count_lattice_paths(x big.Int) *big.Int {
	xfactorial := factorial(x)
	x2 := new(big.Int).Mul(&x, big.NewInt(2))
	numerator := factorial(*x2)
	result := new(big.Int).Div(numerator, xfactorial)
	return result.Div(result, xfactorial)
}

func main() {
	x := big.NewInt(20)
	fmt.Printf("number of lattice paths for %v = %v\n", x.String(), count_lattice_paths(*x).String())
}
