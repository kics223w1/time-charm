package timecharm

import (
	"fmt"
	"math"
	"strconv"
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
	HideYearAndDays           bool
	HideYear                 bool
	HideSeconds              bool
	KeepDecimalsOnWholeSeconds bool
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
	parsed, err := ParseMilliseconds(milliseconds)
	if err != nil {
		return ""
	}
	days := parsed.Days

	if options.HideYearAndDays && days > 0 {
		hours := (days * 24) + parsed.Hours
		add(hours, "hour", "h", nil, &result, options)
	} else {
		if options.HideYear {
			years := days / 365
			remainingDays := days % 365
			add(years, "year", "y", nil, &result, options)
			add(remainingDays, "day", "d", nil, &result, options)
		} else {
			add(parsed.Days, "year", "y", nil, &result, options)
		}

		add(parsed.Hours, "hour", "h", nil, &result, options)
	} 

	add(parsed.Minutes, "minute", "m", nil, &result, options)

	if !options.HideSeconds {

		if options.SeparateMilliseconds || options.FormatSubMilliseconds || (options.ColonNotation && milliseconds < 1000) {
			add(parsed.Seconds, "second", "s", nil, &result, options)
			
			if options.FormatSubMilliseconds {
				add(parsed.Milliseconds, "millisecond", "ms", nil, &result, options)
				add(parsed.Microseconds, "microsecond", "µs", nil, &result, options)
				add(parsed.Nanoseconds, "nanosecond", "ns", nil, &result, options)
			} else {
				millisecondsAndBelow := parsed.Seconds + parsed.Microseconds / 1000 + parsed.Nanoseconds / 1e6
				millisecondsDecimalDigits := 0
				if options.MillisecondsDecimalDigits != 0 {
					millisecondsDecimalDigits = options.MillisecondsDecimalDigits
				}

				var roundedMilliseconds float64
				if millisecondsAndBelow >= 1 {
					roundedMilliseconds = math.Round(float64(millisecondsAndBelow))
				} else {
					roundedMilliseconds = math.Ceil(float64(millisecondsAndBelow))
				}

				var millisecondsString string
				if millisecondsDecimalDigits > 0 {
					millisecondsString = fmt.Sprintf("%.*f", millisecondsDecimalDigits, millisecondsAndBelow)
				} else {
					millisecondsString = fmt.Sprintf("%.0f", roundedMilliseconds)
				}

				millisecondsValue, err := strconv.ParseInt(millisecondsString, 10, 64)
				if err != nil {
					// handle error, e.g., log or return an empty string
					return ""
				}
				add(millisecondsValue, "millisecond", "ms", &millisecondsString, &result, options)
			}
			
		}  else {
			// Calculate seconds
			seconds := (milliseconds / 1000) % 60

			// Determine seconds decimal digits
			secondsDecimalDigits := 1
			if options.SecondsDecimalDigits != 0 {
				secondsDecimalDigits = options.SecondsDecimalDigits
			}

			// Format seconds with specified decimal digits
			secondsFixed := floorDecimals(float64(seconds), secondsDecimalDigits)

			// Remove trailing zeros if not keeping decimals on whole seconds
			secondsString := secondsFixed
			if !options.KeepDecimalsOnWholeSeconds {
				secondsString = strings.TrimRight(strings.TrimRight(secondsFixed, "0"), ".")
			}

			// Add formatted seconds to result
			add(int64(seconds), "second", "s", &secondsString, &result, options)
		}
	}

	if len(result) == 0 {
		if options.Verbose {
			return sign + "0 milliseconds"
		} else {
			return sign + "0ms"
		}
	}

	separator := ":"
	if options.ColonNotation {
		separator = ""
	}

	if options.UnitCount > 0 {
		result = result[:max(options.UnitCount, 1)]
	}

	return sign + strings.Join(result, separator)
}	

func floorDecimals(value float64, decimalDigits int) string {
	factor := math.Pow(10, float64(decimalDigits))
	flooredInterimValue := math.Floor((value * factor) + SECOND_ROUNDING_EPSILON)
	flooredValue := math.Round(flooredInterimValue) / factor
	return fmt.Sprintf("%.*f", decimalDigits, flooredValue)
}


func pluralize(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}

func add(value int64, long, short string, valueString *string, result *[]string, options Options) {
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
			*valueString += " " + pluralize(long, int(value))
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

