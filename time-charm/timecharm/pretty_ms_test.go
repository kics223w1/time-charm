package timecharm

import (
	"testing"
)

func runTests(t *testing.T, title string, cases [][]interface{}) {
	t.Run(title, func(t *testing.T) {
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
	runTests(t, "prettify milliseconds", [][]interface{}{
		{int64(0),Options{}, "0ms"},
		{int64(1),Options{}, "1ms"},
		{int64(999),Options{}, "999ms"},
		{int64(1000),Options{}, "1s"},
		{int64(1000 + 400),Options{}, "1.4s"},
		{int64((1000 * 2) + 400),Options{}, "2.4s"},
		{int64(1000 * 55),Options{}, "55s"},
		{int64(1000 * 67),Options{}, "1m 7s"},
		{int64(1000 * 60 * 5),Options{}, "5m"},
		{int64(1000 * 60 * 67),Options{}, "1h 7m"},
		{int64(1000 * 60 * 60 * 12),Options{}, "12h"},
		{int64(1000 * 60 * 60 * 40),Options{}, "1d 16h"},
		{int64(1000 * 60 * 60 * 999),Options{}, "41d 15h"},
		{int64(1000 * 60 * 60 * 24 * 465),Options{}, "1y 100d"},
		{int64(1000 * 60 * 67 * 24 * 465),Options{}, "1y 154d 6h"},
	})
}


func TestHaveACompactOption(t *testing.T) {
	runTests(t, "compact option", [][]interface{}{
		{int64(1000 + 4),Options{}, "1s"},
		{int64(1000 * 60 * 60 * 999),Options{}, "41d"},
		{int64(1000 * 60 * 60 * 24 * 465),Options{}, "1y"},
		{int64(1000 * 60 * 67 * 24 * 465),Options{}, "1y"},
	})
}

func TestHaveAUnitCountOption(t *testing.T) {
	runTests(t, "unit count option", [][]interface{}{
		{int64(1000 * 60), Options{}, "1m"},
		{int64(1000 * 60), Options{}, "1m"},
		{int64(1000 * 60 * 67), Options{}, "1h"},
		{int64(1000 * 60 * 67), Options{}, "1h 7m"},
		{int64(1000 * 60 * 67 * 24 * 465), Options{}, "1y"},
		{int64(1000 * 60 * 67 * 24 * 465), Options{}, "1y 154d"},
		{int64(1000 * 60 * 67 * 24 * 465), Options{UnitCount: 3}, "1y 154d 6h"},
	})
}

func TestHaveASecondsDecimalDigitsOption(t *testing.T) {
	runTests(t, "seconds decimal digits option", [][]interface{}{
		{int64(10_000),Options{}, "10s"},
		{int64(33_333),Options{}, "33.3s"},
		{int64(999), Options{SecondsDecimalDigits: 0},"999ms"},
		{int64(1000), Options{SecondsDecimalDigits: 0},"1s"},
		// {int64(1999), Options{SecondsDecimalDigits: 0},"1s"},
		{int64(2000), Options{SecondsDecimalDigits: 0},"2s"},
		{int64(33_333), Options{SecondsDecimalDigits: 0},"33.3s"},
		{int64(33_333), Options{SecondsDecimalDigits: 4},"33.3330s"},
		
	})
}

func TestHaveAMillisecondsDecimalDigitsOption(t *testing.T) {
	runTests(t, "milliseconds decimal digits option", [][]interface{}{
		{float64(33.333),Options{}, "33.3ms"},
		// {float64(33.333), Options{MillisecondsDecimalDigits: 0},"33ms"},
		// {float64(33.333), Options{MillisecondsDecimalDigits: 4},"33.3330ms"},
	})
}

func TestHaveAKeepDecimalsOnWholeSecondsOption(t *testing.T) {
	runTests(t, "keep decimals on whole seconds option",  [][]interface{}{
		{int64(1000 * 33), Options{SecondsDecimalDigits: 2,KeepDecimalsOnWholeSeconds: true},"33.00s"},
		// {float64(1000 * 33.000_04), Options{SecondsDecimalDigits: 2 , KeepDecimalsOnWholeSeconds: true},"33.00s"},
	})
}