package main

import (
	"fmt"

	"github.com/kics223w1/time-charm/timecharm"
)

func main() {
	options := timecharm.Options{
		MillisecondsDecimalDigits: 3,
	}
	fmt.Printf("33.333:  %s\n", timecharm.PrettyMilliseconds(float64(33.333), options))
}
