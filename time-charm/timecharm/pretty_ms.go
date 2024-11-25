package timecharm

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const SECOND_ROUNDING_EPSILON = 0.000_000_1
type Options struct {
	ColonNotation              bool
	Compact                    bool
	FormatSubMilliseconds      bool
	SeparateMilliseconds       bool
	Verbose                    bool
	UnitCount                  int
	SecondsDecimalDigits       *int
	MillisecondsDecimalDigits  int
	HideYearAndDays            bool
	HideYear                   bool
	HideSeconds                bool
	KeepDecimalsOnWholeSeconds bool
}

func PrettyMilliseconds(milliseconds interface{}, options Options) string {
	// Convert milliseconds to float64 or int64 based on the input type
	var msFloat64 float64
	var msInt64 int64
	var isFloat bool

	switch v := milliseconds.(type) {
	case int64:
		msInt64 = v
	case float64:
		msFloat64 = v
		isFloat = true
	default:
		// Handle the error case, e.g., return an empty string or an error message
		return ""
	}

	sign := ""
	if isFloat && msFloat64 < 0 {
		sign = "-"
		msFloat64 = -msFloat64
	} else if !isFloat && msInt64 < 0 {
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
		options.SecondsDecimalDigits = nil
		options.MillisecondsDecimalDigits = 0
	}

	result := []string{}
	var parsed TimeComponents
	var err error
	if isFloat {
		parsed, err = ParseMilliseconds(msFloat64)
	} else {
		parsed, err = ParseMilliseconds(msInt64)
	}
	if err != nil {
		return ""
	}
	days := parsed.Days

	fmt.Printf("parsed: %v\n", parsed)

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


	fmt.Printf("parsed.Seconds: %d\n", parsed)

	if !options.HideSeconds {

		if options.SeparateMilliseconds || options.FormatSubMilliseconds || (!options.ColonNotation && (isFloat && msFloat64 < 1000 || !isFloat && msInt64 < 1000)) {

			add(parsed.Seconds, "second", "s", nil, &result, options)

			
			if options.FormatSubMilliseconds {
				add(parsed.Milliseconds, "millisecond", "ms", nil, &result, options)
				add(parsed.Microseconds, "microsecond", "µs", nil, &result, options)
				add(parsed.Nanoseconds, "nanosecond", "ns", nil, &result, options)

			} else {
				millisecondsAndBelow := float64(parsed.Milliseconds) + float64(parsed.Microseconds) / 1000 + float64(parsed.Nanoseconds) / 1e6
				
				fmt.Printf("millisecondsAndBelow: %f\n", millisecondsAndBelow)

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



				// Check if millisecondsString contains a decimal point
				if strings.Contains(millisecondsString, ".") {
					// Parse as float64
					millisecondsValue, err := strconv.ParseFloat(millisecondsString, 64)
					if err != nil {
						// handle error, e.g., log or return an empty string
						return ""
					}
					// Convert to int64 if needed
					millisecondsValueInt64 := int64(millisecondsValue)
					add(millisecondsValueInt64, "millisecond", "ms", &millisecondsString, &result, options)
				} else {
					// Parse as int64
					millisecondsValue, err := strconv.ParseInt(millisecondsString, 10, 64)
					if err != nil {
						// handle error, e.g., log or return an empty string
						return ""
					}
					add(millisecondsValue, "millisecond", "ms", &millisecondsString, &result, options)
				}
			}
			
		}  else {

			// Calculate seconds
			var seconds float64
			if isFloat {
				seconds = math.Mod(msFloat64/1000, 60)
			} else {
				seconds = math.Mod(float64(msInt64)/1000, 60)
			}

			fmt.Printf("huy vao else seconds: %f\n", seconds)


			// Determine seconds decimal digits
			secondsDecimalDigits := 1
			if options.SecondsDecimalDigits != nil { // Check if it's explicitly set
				secondsDecimalDigits = *options.SecondsDecimalDigits
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

			fmt.Printf("huy vao else seconds 2: %f\n", seconds)


			// Add formatted seconds to result
			secondsValue, err := strconv.ParseFloat(secondsString, 64)
			if err != nil {
				// handle error, e.g., log or return an empty string
				return ""
			}
			add(secondsValue, "second", "s", &secondsString, &result, options)
		}
	}

	fmt.Printf("result: %v\n", result)

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

	fmt.Printf("result: %v\n", result)

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


func pluralize(word string, count interface{}) string {

	fmt.Printf("word: %s count: %v\n", word, count)

	// Check if the count is exactly 1 for both int64 and float64
	if count == int64(1) || count == float64(1) {
		return word
	}
	return word + "s"
}

func add(value interface{}, long, short string, valueString *string, result *[]string, options Options) {
	var valueFloat64 float64
	var valueInt64 int64
	var isFloat bool

	switch v := value.(type) {
	case int64:
		valueInt64 = v
	case float64:
		valueFloat64 = v
		isFloat = true
	default:
		return
	}

	if (len(*result) == 0 || !options.ColonNotation) && ((isFloat && valueFloat64 == 0) || (!isFloat && valueInt64 == 0)) && !(options.ColonNotation && short == "m") {
		return
	}

	if valueString == nil {
		if isFloat {
			v := fmt.Sprintf("%.0f", valueFloat64)
			valueString = &v
		} else {
			v := fmt.Sprintf("%d", valueInt64)
			valueString = &v
		}
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
			fmt.Printf("valueFloat64: %f\n", valueFloat64)
			fmt.Printf("valueInt64: %d\n", valueInt64)

			if isFloat {
				*valueString += " " + pluralize(long, valueFloat64)
			} else {
				*valueString += " " + pluralize(long, valueInt64)
			}
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

