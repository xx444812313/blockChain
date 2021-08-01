package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	fmt.Println(math.MaxInt64 )
	bigInt := big.NewInt(math.MaxInt64)
	fmt.Println(bigInt.Mul(bigInt, bigInt).String())

}
