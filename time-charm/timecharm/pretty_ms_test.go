package timecharm

import (
	"testing"
)

func runTests(t *testing.T, title string, defaultOptions Options, cases [][]interface{}) {
	t.Run(title, func(t *testing.T) {
		for _, testCase := range cases {
			var milliseconds int64
			var options Options
			var expected string

			if len(testCase) == 3 {
				milliseconds, options, expected = testCase[0].(int64), testCase[1].(Options), testCase[2].(string)
			} else {
				milliseconds, expected = testCase[0].(int64), testCase[1].(string)
			}

			result := PrettyMilliseconds(milliseconds, options)
			if result != expected {
				t.Errorf("Number(%v): expected %s, got %s", milliseconds, expected, result)
			}
		}
	})
}

func TestPrettifyMilliseconds(t *testing.T) {
	runTests(t, "prettify milliseconds", Options{}, [][]interface{}{
		{int64(0), "0ms"},
		{int64(1), "1ms"},
		{int64(999), "999ms"},
		{int64(1000), "1s"},
		{int64(1000 + 400), "1.4s"},
		{int64((1000 * 2) + 400), "2.4s"},
		{int64(1000 * 55), "55s"},
		{int64(1000 * 67), "1m 7s"},
		{int64(1000 * 60 * 5), "5m"},
		{int64(1000 * 60 * 67), "1h 7m"},
		{int64(1000 * 60 * 60 * 12), "12h"},
		{int64(1000 * 60 * 60 * 40), "1d 16h"},
		{int64(1000 * 60 * 60 * 999), "41d 15h"},
		{int64(1000 * 60 * 60 * 24 * 465), "1y 100d"},
		{int64(1000 * 60 * 67 * 24 * 465), "1y 154d 6h"},
	})
}

