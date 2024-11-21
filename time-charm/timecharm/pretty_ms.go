package timecharm

import (
	"fmt"
	"math"
	"strings"
)

const SECOND_ROUNDING_EPSILON = 0.000_000_1


type Options struct {
	ColonNotation bool
	Compact       bool
	FormatSubMilliseconds bool
	SeparateMilliseconds  bool
	Verbose               bool
	UnitCount             int
	SecondsDecimalDigits  int
	MillisecondsDecimalDigits int
}

func PrettyMilliseconds(milliseconds int64, options Options) string {

	sign := ""
	if milliseconds < 0 {
		sign = "-"
		milliseconds = -milliseconds
	}

	if options.ColonNotation {
		options.Compact = false
		options.SeparateMilliseconds = false
		options.FormatSubMilliseconds = false
		options.Verbose = false
	}

	if options.Compact {
		options.UnitCount = 1
		options.SeparateMilliseconds = false
		options.FormatSubMilliseconds = false
	}

	result := []string{}


}

func floorDecimals(value float64, decimalDigits int) string {
	factor := math.Pow(10, float64(decimalDigits))
	flooredInterimValue := math.Floor((value * factor) + SECOND_ROUNDING_EPSILON)
	flooredValue := math.Round(flooredInterimValue) / factor
	return fmt.Sprintf("%.*f", decimalDigits, flooredValue)
}

func add(value int, long, short string, valueString *string, result *[]string, options Options) {
	if (len(*result) == 0 || !options.ColonNotation) && value == 0 && !(options.ColonNotation && short == "m") {
		return
	}

	if valueString == nil {
		v := fmt.Sprintf("%d", value)
		valueString = &v
	}

	if options.ColonNotation {
		wholeDigits := len(*valueString)
		if strings.Contains(*valueString, ".") {
			wholeDigits = len(strings.Split(*valueString, ".")[0])
		}
		minLength := 1
		if len(*result) > 0 {
			minLength = 2
		}
		*valueString = strings.Repeat("0", max(0, minLength-wholeDigits)) + *valueString
	} else {
		if options.Verbose {
			*valueString += " " + pluralize(long, value)
		} else {
			*valueString += short
		}
	}

	*result = append(*result, *valueString)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pluralize(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}