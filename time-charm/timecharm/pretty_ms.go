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

func PrettyMilliseconds(milliseconds interface{}, options Options) string {
	// Convert milliseconds to float64
	var ms float64
	switch v := milliseconds.(type) {
	case int64:
		ms = float64(v)
	case float64:
		ms = v
	default:
		// Handle the error case, e.g., return an empty string or an error message
		return ""
	}
	// Convert to int64 for further processing
	msInt64 := int64(ms)

	sign := ""
	if msInt64 < 0 {
		sign = "-"
		msInt64 = -msInt64
	}

	if options.ColonNotation {
		options.Compact = false
		options.SeparateMilliseconds = false
		options.FormatSubMilliseconds = false
		options.Verbose = false
	}

	if options.Compact {
		options.UnitCount = 1
		options.SecondsDecimalDigits = 0
		options.MillisecondsDecimalDigits = 0
	}

	result := []string{}
	parsed, err := ParseMilliseconds(msInt64)
	if err != nil {
		return ""
	}
	days := parsed.Days

	if options.HideYearAndDays && days > 0 {
		hours := (days * 24) + parsed.Hours
		add(hours, "hour", "h", nil, &result, options)
	} else {
		if options.HideYear {
			add(parsed.Days, "day", "d", nil, &result, options)
		} else {
			years := days / 365
			remainingDays := days % 365
			add(years, "year", "y", nil, &result, options)
			add(remainingDays, "day", "d", nil, &result, options)
		}

		add(parsed.Hours, "hour", "h", nil, &result, options)
	} 

	add(parsed.Minutes, "minute", "m", nil, &result, options)

	if !options.HideSeconds {

		if options.SeparateMilliseconds || options.FormatSubMilliseconds || (!options.ColonNotation && msInt64 < 1000) {
			add(parsed.Seconds, "second", "s", nil, &result, options)

			if options.FormatSubMilliseconds {
				add(parsed.Milliseconds, "millisecond", "ms", nil, &result, options)
				add(parsed.Microseconds, "microsecond", "µs", nil, &result, options)
				add(parsed.Nanoseconds, "nanosecond", "ns", nil, &result, options)

			} else {
				millisecondsAndBelow := float64(parsed.Milliseconds) + float64(parsed.Microseconds) / 1000 + float64(parsed.Nanoseconds) / 1e6
				
				millisecondsDecimalDigits := 0
				if options.MillisecondsDecimalDigits != 0 {
					millisecondsDecimalDigits = options.MillisecondsDecimalDigits
				}



				var roundedMilliseconds float64
				if millisecondsAndBelow >= 1 {
					roundedMilliseconds = math.Round(millisecondsAndBelow)
				} else {
					roundedMilliseconds = math.Ceil(millisecondsAndBelow)
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
			seconds := math.Mod(float64(msInt64)/1000, 60)

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

				// After removing trailing zeros, we have to make sure the number
				// of decimal digits equals the specified number of decimal digits
				decimalPart := strings.Split(secondsString, ".")
				if len(decimalPart) > 1 {
					secondsString = decimalPart[0] + "." + decimalPart[1] + strings.Repeat("0", secondsDecimalDigits-len(decimalPart[1]))
				}
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

	separator := " "
	if options.ColonNotation {
		separator = ""
	}

	if options.UnitCount > 0 {
		result = result[:max(options.UnitCount, 1)]
	}

	return sign + strings.Join(result, separator)
}	

func floorDecimals(value float64, decimalDigits int) string {

	fmt.Printf("value: %f decimalDigits: %d\n", value, decimalDigits)

	factor := math.Pow(10, float64(decimalDigits))
	flooredValue := math.Floor(value*factor+SECOND_ROUNDING_EPSILON) / factor

	rs:= fmt.Sprintf("%.*f", decimalDigits, flooredValue)

	fmt.Printf("rs: %s\n", rs)

	return rs
}


func pluralize(word string, count int) string {

	fmt.Printf("word: %s count: %d\n", word, count)
	if count == 1 {
		return word
	}
	return word + "s"
}

func add(value interface{}, long, short string, valueString *string, result *[]string, options Options) {
	// Convert value to int64 for comparison
	valueInt64, ok := value.(int64)
	if !ok {
		return
	}

	if (len(*result) == 0 || !options.ColonNotation) && valueInt64 == 0 && !(options.ColonNotation && short == "m") {
		return
	}

	if valueString == nil {
		v := fmt.Sprintf("%d", valueInt64)
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

		fmt.Printf("valueString: %d\n", valueInt64)

		if options.Verbose {
			*valueString += " " + pluralize(long, int(valueInt64))
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

