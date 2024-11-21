package timecharm

import (
	"errors"
	"fmt"
	"math"
)

type TimeComponents struct {
	Days         int64
	Hours        int64
	Minutes      int64
	Seconds      int64
	Milliseconds int64
	Microseconds int64
	Nanoseconds  int64
}

func toZeroIfInfinity(value float64) float64 {
	if math.IsInf(value, 0) {
		return 0
	}
	return value
}

func parseNumber(milliseconds float64) TimeComponents {
	return TimeComponents{
		Days:         int64(milliseconds / 86_400_000),
		Hours:        int64(math.Mod(milliseconds/3_600_000, 24)),
		Minutes:      int64(math.Mod(milliseconds/60_000, 60)),
		Seconds:      int64(math.Mod(milliseconds/1000, 60)),
		Milliseconds: int64(math.Mod(milliseconds, 1000)),
		Microseconds: int64(math.Mod(toZeroIfInfinity(milliseconds*1000), 1000)),
		Nanoseconds:  int64(math.Mod(toZeroIfInfinity(milliseconds*1e6), 1000)),
	}
}

func parseBigint(milliseconds int64) TimeComponents {
	return TimeComponents{
		Days:         milliseconds / 86_400_000,
		Hours:        (milliseconds / 3_600_000) % 24,
		Minutes:      (milliseconds / 60_000) % 60,
		Seconds:      (milliseconds / 1000) % 60,
		Milliseconds: milliseconds % 1000,
		Microseconds: 0,
		Nanoseconds:  0,
	}
}

func ParseMilliseconds(input interface{}) (TimeComponents, error) {
	// Check if the input is NaN
	if value, ok := input.(float64); ok && math.IsNaN(value) {
		return TimeComponents{}, fmt.Errorf("input is NaN")
	}

	switch v := input.(type) {
	case float64:
		if !math.IsInf(v, 0) {
			components := parseNumber(math.Abs(v))
			if v < 0 {
				components.Days = -components.Days
				components.Hours = -components.Hours
				components.Minutes = -components.Minutes
				components.Seconds = -components.Seconds
				components.Milliseconds = -components.Milliseconds
			}
			return components, nil
		}
	case int64:
		return parseBigint(v), nil
	default:
		return TimeComponents{}, errors.New("expected a number of type int64 or float64")
	}
	return TimeComponents{}, errors.New("expected a number of type int64 or float64")
}
