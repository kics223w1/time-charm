package timecharm

import (
	"testing"
)

func ptr(i int) *int {
	return &i
}

func runTests(t *testing.T, cases [][]interface{}) {
	t.Run("TestPrettyMilliseconds", func(t *testing.T) {
		for _, testCase := range cases {
			var milliseconds interface{}
			var options Options
			var expected string

			milliseconds, options, expected = testCase[0], testCase[1].(Options), testCase[2].(string)


			result := PrettyMilliseconds(milliseconds, options)
			if result != expected {
				t.Errorf("Number(%v): expected %s, got %s", milliseconds, expected, result)
			}
		}
	})
}

func TestPrettifyMilliseconds(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(0)), interface{}(Options{}), interface{}("0ms")},
		{interface{}(int64(1)), interface{}(Options{}), interface{}("1ms")},
		{interface{}(int64(999)), interface{}(Options{}), interface{}("999ms")},
		{interface{}(int64(1000)), interface{}(Options{}), interface{}("1s")},
		{interface{}(int64(1000 + 400)), interface{}(Options{}), interface{}("1.4s")},
		{interface{}(int64((1000 * 2) + 400)), interface{}(Options{}), interface{}("2.4s")},
		{interface{}(int64(1000 * 55)), interface{}(Options{}), interface{}("55s")},
		{interface{}(int64(1000 * 67)), interface{}(Options{}), interface{}("1m 7s")},
		{interface{}(int64(1000 * 60 * 5)), interface{}(Options{}), interface{}("5m")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{}), interface{}("1h 7m")},
		{interface{}(int64(1000 * 60 * 60 * 12)), interface{}(Options{}), interface{}("12h")},
		{interface{}(int64(1000 * 60 * 60 * 40)), interface{}(Options{}), interface{}("1d 16h")},
		{interface{}(int64(1000 * 60 * 60 * 999)), interface{}(Options{}), interface{}("41d 15h")},
		{interface{}(int64(1000 * 60 * 60 * 24 * 465)), interface{}(Options{}), interface{}("1y 100d")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{}), interface{}("1y 154d 6h")},
	})
}


func TestHaveACompactOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000 + 4)), interface{}(Options{Compact: true}), interface{}("1s")},
		{interface{}(int64(1000 * 60 * 60 * 999)), interface{}(Options{Compact: true}), interface{}("41d")},
		{interface{}(int64(1000 * 60 * 60 * 24 * 465)), interface{}(Options{Compact: true}), interface{}("1y")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Compact: true}), interface{}("1y")},
	})
}

func TestHaveAUnitCountOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000 * 60)), interface{}(Options{UnitCount: 0}), interface{}("1m")},
		{interface{}(int64(1000 * 60)), interface{}(Options{UnitCount: 1}), interface{}("1m")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{UnitCount: 1}), interface{}("1h")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{UnitCount: 2}), interface{}("1h 7m")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{UnitCount: 1}), interface{}("1y")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{UnitCount: 2}), interface{}("1y 154d")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{UnitCount: 3}), interface{}("1y 154d 6h")},
	})
}

func TestHaveASecondsDecimalDigitsOption(t *testing.T) {
	runTests(t,  [][]interface{}{
		{interface{}(int64(10_000)), interface{}(Options{}), interface{}("10s")},
		{interface{}(int64(33_333)), interface{}(Options{}), interface{}("33.3s")},
		{interface{}(int64(999)), interface{}(Options{SecondsDecimalDigits: ptr(0)}), interface{}("999ms")},
		{interface{}(int64(1000)), interface{}(Options{SecondsDecimalDigits: ptr(0)}), interface{}("1s")},
		{interface{}(int64(1999)), interface{}(Options{SecondsDecimalDigits: ptr(0)}), interface{}("1s")},
		{interface{}(int64(2000)), interface{}(Options{SecondsDecimalDigits: ptr(0)}), interface{}("2s")},
		{interface{}(int64(33_333)), interface{}(Options{SecondsDecimalDigits: ptr(0)}), interface{}("33s")},
		{interface{}(int64(33_333)), interface{}(Options{SecondsDecimalDigits: ptr(4)}), interface{}("33.3330s")},
	})
}

func TestHaveAMillisecondsDecimalDigitsOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(float64(33.333)), interface{}(Options{}), interface{}("33ms")},
		{interface{}(float64(33.333)), interface{}(Options{MillisecondsDecimalDigits: 0}), interface{}("33ms")},
		{interface{}(float64(33.333)), interface{}(Options{MillisecondsDecimalDigits: 4}), interface{}("33.3330ms")},
	})
}

func TestHaveAKeepDecimalsOnWholeSecondsOption(t *testing.T) {
	runTests(t,  [][]interface{}{
		{interface{}(int64(1000 * 33)), interface{}(Options{SecondsDecimalDigits: ptr(2), KeepDecimalsOnWholeSeconds: true}), interface{}("33.00s")},
		{interface{}(float64(1000 * 33.000_04)), interface{}(Options{SecondsDecimalDigits: ptr(2), KeepDecimalsOnWholeSeconds: true}), interface{}("33.00s")},
	})
}

func TestHaveAVerboseOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(0)), interface{}(Options{Verbose: true}), interface{}("0 milliseconds")},
		{interface{}(float64(0.1)), interface{}(Options{Verbose: true}), interface{}("1 millisecond")},
		{interface{}(int64(1)), interface{}(Options{Verbose: true}), interface{}("1 millisecond")},
		{interface{}(int64(1000)), interface{}(Options{Verbose: true}), interface{}("1 second")},
		{interface{}(int64(1000 + 400)), interface{}(Options{Verbose: true}), interface{}("1.4 seconds")},
		{interface{}(int64(1000 * 2 + 400)), interface{}(Options{Verbose: true}), interface{}("2.4 seconds")},
		{interface{}(int64(1000 * 5)), interface{}(Options{Verbose: true}), interface{}("5 seconds")},
		{interface{}(int64(1000 * 55)), interface{}(Options{Verbose: true}), interface{}("55 seconds")},
		{interface{}(int64(1000 * 67)), interface{}(Options{Verbose: true}), interface{}("1 minute 7 seconds")},
		{interface{}(int64(1000 * 60 * 5)), interface{}(Options{Verbose: true}), interface{}("5 minutes")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{Verbose: true}), interface{}("1 hour 7 minutes")},
		{interface{}(int64(1000 * 60 * 60 * 12)), interface{}(Options{Verbose: true}), interface{}("12 hours")},
		{interface{}(int64(1000 * 60 * 60 * 40)), interface{}(Options{Verbose: true}), interface{}("1 day 16 hours")},
		{interface{}(int64(1000 * 60 * 60 * 999)), interface{}(Options{Verbose: true}), interface{}("41 days 15 hours")},
		{interface{}(int64(1000 * 60 * 60 * 24 * 465)), interface{}(Options{Verbose: true}), interface{}("1 year 100 days")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true}), interface{}("1 year 154 days 6 hours")},
	})
}



