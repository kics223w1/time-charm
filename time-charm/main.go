package main

import (
	"fmt"

	"math/big"

	"github.com/kics223w1/time-charm/timecharm"
)

func main() {
	maxValue := new(big.Int)

	options := timecharm.Options{
		MillisecondsDecimalDigits: (4),
	}

	fmt.Printf("huy %s\n", timecharm.PrettyMilliseconds((maxValue), options))
}
