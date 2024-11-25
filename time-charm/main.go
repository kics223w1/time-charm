package main

import (
	"fmt"

	"github.com/kics223w1/time-charm/timecharm"
)

func main() {
	options := timecharm.Options{
		MillisecondsDecimalDigits: (4),
	}
	fmt.Printf("huy %s\n", timecharm.PrettyMilliseconds(float64(33.333), options))
}
