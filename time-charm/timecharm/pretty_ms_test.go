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

func TestHaveASeparateMillisecondsOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1100)), interface{}(Options{SeparateMilliseconds: false}), interface{}("1.1s")},
		{interface{}(int64(1100)), interface{}(Options{SeparateMilliseconds: true}), interface{}("1s 100ms")},
	})
} 


func TestHaveAFormatSubMillisecondsOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(float64(0.4)), interface{}(Options{FormatSubMilliseconds: true}), interface{}("400µs")},
		{interface{}(float64(0.123_571)), interface{}(Options{FormatSubMilliseconds: true}), interface{}("123µs 571ns")},
		{interface{}(float64(0.123_456_789)), interface{}(Options{FormatSubMilliseconds: true}), interface{}("123µs 456ns")},
		{interface{}(float64((60 * 60 * 1000) + (23 * 1000) + 433 + 0.123_456)), interface{}(Options{FormatSubMilliseconds: true}), interface{}("1h 23s 433ms 123µs 456ns")},
	})
} 

func TestVerboseAndCompactOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000)), interface{}(Options{Compact: true, Verbose: true}), interface{}("1 second")},
		// {interface{}(int64(1000 + 400)), interface{}(Options{Compact: true, Verbose: true}), interface{}("1 second")},
		// {interface{}(int64((1000 * 2) + 400)), interface{}(Options{Compact: true, Verbose: true}), interface{}("2 seconds")},
		{interface{}(int64(1000 * 5)), interface{}(Options{Compact: true, Verbose: true}), interface{}("5 seconds")},
		{interface{}(int64(1000 * 55)), interface{}(Options{Compact: true, Verbose: true}), interface{}("55 seconds")},
		{interface{}(int64(1000 * 67)), interface{}(Options{Compact: true, Verbose: true}), interface{}("1 minute")},
		{interface{}(int64(1000 * 60 * 5)), interface{}(Options{Compact: true, Verbose: true}), interface{}("5 minutes")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{Compact: true, Verbose: true}), interface{}("1 hour")},
		{interface{}(int64(1000 * 60 * 60 * 12)), interface{}(Options{Compact: true, Verbose: true}), interface{}("12 hours")},
		{interface{}(int64(1000 * 60 * 60 * 40)), interface{}(Options{Compact: true, Verbose: true}), interface{}("1 day")},
		{interface{}(int64(1000 * 60 * 60 * 999)), interface{}(Options{Compact: true, Verbose: true}), interface{}("41 days")},
		{interface{}(int64(1000 * 60 * 60 * 24 * 465)), interface{}(Options{Compact: true, Verbose: true}), interface{}("1 year")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 750)), interface{}(Options{Compact: true, Verbose: true}), interface{}("2 years")},
	})
}

func TestVerboseAndUnitCountOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000 * 60)), interface{}(Options{Verbose: true, UnitCount: 1}), interface{}("1 minute")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{Verbose: true, UnitCount: 1}), interface{}("1 hour")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{Verbose: true, UnitCount: 2}), interface{}("1 hour 7 minutes")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true, UnitCount: 1}), interface{}("1 year")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true, UnitCount: 2}), interface{}("1 year 154 days")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true, UnitCount: 3}), interface{}("1 year 154 days 6 hours")},
	})
}

func TestVerboseAndSecondsDecimalDigitsOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("1 second")},
		{interface{}(int64(1000 + 400)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("1.4000 seconds")},
		{interface{}(int64((1000 * 2) + 400)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("2.4000 seconds")},
		{interface{}(int64((1000 * 5) + 254)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("5.2540 seconds")},
		{interface{}(int64(33_333)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("33.3330 seconds")},
	})
}

func TestVerboseAndMillisecondsDecimalDigitsOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(float64(1)), interface{}(Options{Verbose: true, MillisecondsDecimalDigits: 4}), interface{}("1.0000 millisecond")},
		// {interface{}(float64(1 + 0.4)), interface{}(Options{Verbose: true, MillisecondsDecimalDigits: 4}), interface{}("1.4000 milliseconds")},
		{interface{}(float64((1 * 2) + 0.4)), interface{}(Options{Verbose: true, MillisecondsDecimalDigits: 4}), interface{}("2.4000 milliseconds")},
		{interface{}(float64((1 * 5) + 0.254)), interface{}(Options{Verbose: true, MillisecondsDecimalDigits: 4}), interface{}("5.2540 milliseconds")},
		{interface{}(float64(33.333)), interface{}(Options{Verbose: true, MillisecondsDecimalDigits: 4}), interface{}("33.3330 milliseconds")},
	})
}

func TestVerboseAndFormatSubMillisecondsOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(float64(0.4)), interface{}(Options{FormatSubMilliseconds: true, Verbose: true}), interface{}("400 microseconds")},
		{interface{}(float64(0.123_571)), interface{}(Options{FormatSubMilliseconds: true, Verbose: true}), interface{}("123 microseconds 571 nanoseconds")},
		{interface{}(float64(0.123_456_789)), interface{}(Options{FormatSubMilliseconds: true, Verbose: true}), interface{}("123 microseconds 456 nanoseconds")},
		{interface{}(float64(0.001)), interface{}(Options{FormatSubMilliseconds: true, Verbose: true}), interface{}("1 microsecond")},
	})
}

func TestCompactOverridesUnitCountOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true, Compact: true, UnitCount: 1}), interface{}("1 year")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true, Compact: true, UnitCount: 2}), interface{}("1 year")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{Verbose: true, Compact: true, UnitCount: 3}), interface{}("1 year")},
	})
}

func TestSeparateMillisecondsAndFormatSubMillisecondsOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(float64(1010.340_067)), interface{}(Options{SeparateMilliseconds: true, FormatSubMilliseconds: true}), interface{}("1s 10ms 340µs 67ns")},
		{interface{}(float64(60*1000 + 34 + 0.000_005)), interface{}(Options{SeparateMilliseconds: true, FormatSubMilliseconds: true}), interface{}("1m 34ms 5ns")},
	})
}

func TestProperlyRoundsMillisecondsWithSecondsDecimalDigits(t *testing.T) {
	runTests(t, [][]interface{}{
		// {interface{}(int64(3 * 60 * 1000)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("3 minutes")},
		{interface{}(int64((3 * 60 * 1000) - 1)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("2 minutes 59 seconds")},
		// {interface{}(int64(365 * 24 * 3600 * 1000)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("1 year")},
		{interface{}(int64((365 * 24 * 3600 * 1000) - 1)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("364 days 23 hours 59 minutes 59 seconds")},
		// {interface{}(int64(24 * 3600 * 1000)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("1 day")},
		{interface{}(int64((24 * 3600 * 1000) - 1)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("23 hours 59 minutes 59 seconds")},
		// {interface{}(int64(3600 * 1000)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("1 hour")},
		{interface{}(int64((3600 * 1000) - 1)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("59 minutes 59 seconds")},
		// {interface{}(int64(2 * 3600 * 1000)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("2 hours")},
		{interface{}(int64((2 * 3600 * 1000) - 1)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(0)}), interface{}("1 hour 59 minutes 59 seconds")},
	})
}

func TestNegativeMillisecondsWithOptions(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(-0)), interface{}(Options{}), interface{}("0ms")},
		{interface{}(float64(-0.1)), interface{}(Options{}), interface{}("-1ms")},
		{interface{}(int64(-1)), interface{}(Options{}), interface{}("-1ms")},
		{interface{}(int64(-999)), interface{}(Options{}), interface{}("-999ms")},
		{interface{}(int64(-1000)), interface{}(Options{}), interface{}("-1s")},
		{interface{}(int64(-1000 + 400)), interface{}(Options{}), interface{}("-600ms")},
		{interface{}(int64((-1000 * 2) + 400)), interface{}(Options{}), interface{}("-1.6s")},
		{interface{}(int64(-13_370)), interface{}(Options{}), interface{}("-13.3s")},
		{interface{}(int64(-9007199254740991)), interface{}(Options{}), interface{}("-285616y 151d 8h 59m 0.9s")},
		// With compact option
		{interface{}(int64(-1000 * 60 * 60 * 999)), interface{}(Options{Compact: true}), interface{}("-41d")},
		{interface{}(int64(-1000 * 60 * 60 * 24 * 465)), interface{}(Options{Compact: true}), interface{}("-1y")},
		// With unit-count
		{interface{}(int64(-1000 * 60 * 67)), interface{}(Options{UnitCount: 2}), interface{}("-1h 7m")},
		{interface{}(int64(-1000 * 60 * 67 * 24 * 465)), interface{}(Options{UnitCount: 1}), interface{}("-1y")},
		{interface{}(int64(-1000 * 60 * 67 * 24 * 465)), interface{}(Options{UnitCount: 2}), interface{}("-1y 154d")},
		// With verbose and secondsDecimalDigits
		{interface{}(int64((-1000 * 5) - 254)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("-5.2540 seconds")},
		{interface{}(int64(-33_333)), interface{}(Options{Verbose: true, SecondsDecimalDigits: ptr(4)}), interface{}("-33.3330 seconds")},
		// With verbose and compact
		{interface{}(int64(-1000 * 60 * 5)), interface{}(Options{Verbose: true, Compact: true}), interface{}("-5 minutes")},
		{interface{}(int64(-1000 * 60 * 67)), interface{}(Options{Verbose: true, Compact: true}), interface{}("-1 hour")},
		{interface{}(int64(-1000 * 60 * 60 * 12)), interface{}(Options{Verbose: true, Compact: true}), interface{}("-12 hours")},
		// With separateMilliseconds option
		{interface{}(int64(-1001)), interface{}(Options{SeparateMilliseconds: true}), interface{}("-1s 1ms")},
		{interface{}(int64(-1234)), interface{}(Options{SeparateMilliseconds: true}), interface{}("-1s 234ms")},
		// With formatSubMilliseconds option
		{interface{}(float64(-1.234_567)), interface{}(Options{FormatSubMilliseconds: true}), interface{}("-1ms 234µs 567ns")},
		{interface{}(float64(-1234.567)), interface{}(Options{FormatSubMilliseconds: true}), interface{}("-1s 234ms 567µs")},
	})
}

func TestColonNotationOption(t *testing.T) {
	runTests(t, [][]interface{}{
		// Default formats
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true}), interface{}("0:01")},
		{interface{}(int64(1543)), interface{}(Options{ColonNotation: true}), interface{}("0:01.5")},
		{interface{}(int64(1000 * 60)), interface{}(Options{ColonNotation: true}), interface{}("1:00")},
		{interface{}(int64(1000 * 90)), interface{}(Options{ColonNotation: true}), interface{}("1:30")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true}), interface{}("1:35.5")},
		{interface{}(int64((1000 * 60 * 10) + 543)), interface{}(Options{ColonNotation: true}), interface{}("10:00.5")},
		{interface{}(int64((1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true}), interface{}("59:59.5")},
		{interface{}(int64((1000 * 60 * 60 * 15) + (1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true}), interface{}("15:59:59.5")},

		// Together with `secondsDecimalDigits`
		// {interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0)}), interface{}("0:00")},
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1)}), interface{}("0:00.9")},
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(2)}), interface{}("0:00.99")},
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("0:00.999")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0)}), interface{}("0:01")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1)}), interface{}("0:01")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(2)}), interface{}("0:01")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("0:01")},
		{interface{}(int64(1001)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0)}), interface{}("0:01")},
		{interface{}(int64(1001)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1)}), interface{}("0:01")},
		{interface{}(int64(1001)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(2)}), interface{}("0:01")},
		{interface{}(int64(1001)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("0:01.001")},
		{interface{}(int64(1543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0)}), interface{}("0:01")},
		{interface{}(int64(1543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1)}), interface{}("0:01.5")},
		{interface{}(int64(1543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(2)}), interface{}("0:01.54")},
		{interface{}(int64(1543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("0:01.543")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0)}), interface{}("1:35")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1)}), interface{}("1:35.5")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(2)}), interface{}("1:35.54")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("1:35.543")},
		{interface{}(int64((1000 * 60 * 10) + 543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("10:00.543")},
		{interface{}(int64((1000 * 60 * 60 * 15) + (1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3)}), interface{}("15:59:59.543")},

		// Together with `keepDecimalsOnWholeSeconds`
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0), KeepDecimalsOnWholeSeconds: true}), interface{}("0:00")},
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1), KeepDecimalsOnWholeSeconds: true}), interface{}("0:00.9")},
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(2), KeepDecimalsOnWholeSeconds: true}), interface{}("0:00.99")},
		{interface{}(int64(999)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3), KeepDecimalsOnWholeSeconds: true}), interface{}("0:00.999")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, KeepDecimalsOnWholeSeconds: true}), interface{}("0:01.0")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0), KeepDecimalsOnWholeSeconds: true}), interface{}("0:01")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1), KeepDecimalsOnWholeSeconds: true}), interface{}("0:01.0")},
		{interface{}(int64(1000)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3), KeepDecimalsOnWholeSeconds: true}), interface{}("0:01.000")},
		{interface{}(int64(1000 * 90)), interface{}(Options{ColonNotation: true, KeepDecimalsOnWholeSeconds: true}), interface{}("1:30.0")},
		{interface{}(int64(1000 * 90)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3), KeepDecimalsOnWholeSeconds: true}), interface{}("1:30.000")},
		{interface{}(int64(1000 * 60 * 10)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(3), KeepDecimalsOnWholeSeconds: true}), interface{}("10:00.000")},

		// Together with `unitCount`
		{interface{}(int64(1000 * 90)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0), UnitCount: 1}), interface{}("1")},
		// {interface{}(int64(1000 * 90)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0), UnitCount: 2}), interface{}("1:30")},
		// {interface{}(int64(1000 * 60 * 90)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(0), UnitCount: 3}), interface{}("1:30:00")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1), UnitCount: 1}), interface{}("1")},
		{interface{}(int64(95_543)), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1), UnitCount: 2}), interface{}("1:35.5")},
		{interface{}(int64(95_543 + (1000 * 60 * 60))), interface{}(Options{ColonNotation: true, SecondsDecimalDigits: ptr(1), UnitCount: 3}), interface{}("1:01:35.5")},

		// Make sure incompatible options fall back to `colonNotation`
		{interface{}(int64((1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true, FormatSubMilliseconds: true}), interface{}("59:59.5")},
		{interface{}(int64((1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true, SeparateMilliseconds: true}), interface{}("59:59.5")},
		{interface{}(int64((1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true, Verbose: true}), interface{}("59:59.5")},
		{interface{}(int64((1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{ColonNotation: true, Compact: true}), interface{}("59:59.5")},

		// Big numbers
		{interface{}(int64(9007199254740991)), interface{}(Options{ColonNotation: true}), interface{}("285616:151:08:59:00.9")},
	})
}

func TestHideYearOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64(1000 * 60)), interface{}(Options{HideYear: true}), interface{}("1m")},
		{interface{}(int64(1000 * 60)), interface{}(Options{HideYear: false}), interface{}("1m")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{HideYear: true}), interface{}("1h 7m")},
		{interface{}(int64(1000 * 60 * 67)), interface{}(Options{HideYear: false}), interface{}("1h 7m")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{HideYear: false}), interface{}("1y 154d 6h")},
		{interface{}(int64(1000 * 60 * 67 * 24 * 465)), interface{}(Options{HideYear: true}), interface{}("519d 6h")},
		{interface{}(int64((1000 * 60 * 67 * 24 * 465) + (1000 * 60) + 6500)), interface{}(Options{HideYear: false}), interface{}("1y 154d 6h 1m 6.5s")},
		{interface{}(int64((1000 * 60 * 67 * 24 * 465) + (1000 * 60) + 6500)), interface{}(Options{HideYear: true}), interface{}("519d 6h 1m 6.5s")},
	})
}

func TestHideSecondsOption(t *testing.T) {
	runTests(t, [][]interface{}{
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: false}), interface{}("1m 6.5s")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true}), interface{}("1m")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true, SecondsDecimalDigits: ptr(3)}), interface{}("1m")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true, KeepDecimalsOnWholeSeconds: true}), interface{}("1m")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true, FormatSubMilliseconds: true}), interface{}("1m")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true, SeparateMilliseconds: true}), interface{}("1m")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true, Verbose: true}), interface{}("1 minute")},
		{interface{}(int64((1000 * 60) + 6500)), interface{}(Options{HideSeconds: true, Compact: true}), interface{}("1m")},
	})
}

func TestHideYearAndDaysHideSecondsAndColonNotationOptions(t *testing.T) {
	// bigIntValue, _ := big.NewInt(0).SetString("49935920412842103004035395481028987999464046534956943499699299111988127994452371877941544064657466158761238598198439573398422590802628939657907651862093754718347197382375356132290413913997035817798852363459759428417939788028673041157169044258923152298554951723373534213538382550255361078125112229495590", 10)
	runTests(t, [][]interface{}{
		{interface{}(int64((1000 * 60 * 60 * 15) + (1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{HideSeconds: true, HideYearAndDays: true, ColonNotation: true}), interface{}("15:59")},
		{interface{}(int64((1000 * 60 * 67 * 24 * 465) + (1000 * 60 * 60 * 15) + (1000 * 60 * 59) + (1000 * 59) + 543)), interface{}(Options{HideSeconds: true, HideYearAndDays: true, ColonNotation: true}), interface{}("12477:59")},
		// For BigInt, Go doesn't have a direct equivalent, so we use a large integer representation
		// This is a placeholder; you might need to adjust it based on your actual implementation
		// {interface{}(bigIntValue), interface{}(Options{HideSeconds: true, HideYearAndDays: true, ColonNotation: true}), interface{}("49935920412842103004035395481028987999464046534956943499699299111988127994452371877941544064657466158761238598198439573398422590802628939657907651862093754718347197382375356132290413913997035817798852363459759428417939788028673041157169044258923152298554951723373534213538382550255361078125112229495590:14")},
	})
}










