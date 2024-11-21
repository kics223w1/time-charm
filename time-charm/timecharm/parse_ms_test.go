package timecharm

import (
	"math"
	"testing"
)

func TestParseMilliseconds(t *testing.T) {
	tests := []struct {
		milliseconds int64
		expected     TimeComponents
	}{
		{1000 + 400, TimeComponents{0, 0, 0, 1, 400, 0, 0}},
		{1000 * 55, TimeComponents{0, 0, 0, 55, 0, 0, 0}},
		{1000 * 67, TimeComponents{0, 0, 1, 7, 0, 0, 0}},
		{1000 * 60 * 5, TimeComponents{0, 0, 5, 0, 0, 0, 0}},
		{1000 * 60 * 67, TimeComponents{0, 1, 7, 0, 0, 0, 0}},
		{1000 * 60 * 60 * 12, TimeComponents{0, 12, 0, 0, 0, 0, 0}},
		{1000 * 60 * 60 * 40, TimeComponents{1, 16, 0, 0, 0, 0, 0}},
		{1000 * 60 * 60 * 999, TimeComponents{41, 15, 0, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, err := ParseMilliseconds(tt.milliseconds)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}

	t.Run("float value part 1", func(t *testing.T) {
		result, err := ParseMilliseconds(1000 * 60 + 500 + 0.345678)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := TimeComponents{0, 0, 1, 0, 500, 345, 678}
		if result != expected {
			t.Errorf("got %v, want %v", result, expected)
		}
	})

	t.Run("float value part 2", func(t *testing.T) {
		result, err := ParseMilliseconds(0.000543)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := TimeComponents{0, 0, 0, 0, 0, 0, 543}
		if result != expected {
			t.Errorf("got %v, want %v", result, expected)
		}
	})

	t.Run("max value", func(t *testing.T) {
		result, err := ParseMilliseconds(math.MaxFloat64)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := TimeComponents{
			Days:         math.MaxInt64,
			Hours:        8,
			Minutes:      8,
			Seconds:      48,
			Milliseconds: 368,
			Microseconds: 0,
			Nanoseconds:  0,
		}
		if result != expected {
			t.Errorf("got %v, want %v", result, expected)
		}
	})

	t.Run("min value", func(t *testing.T) {
		result, err := ParseMilliseconds(math.SmallestNonzeroFloat64)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := TimeComponents{0, 0, 0, 0, 0, 0, 0}
		if result != expected {
			t.Errorf("got %v, want %v", result, expected)
		}
	})
}



func TestHandleNegativeMilliseconds(t *testing.T) {
	times := []int64{
		// 0.0005,
		// 0.3,
		100 + 400,
		1000 * 55,
		1000 * 67,
		1000 * 60 * 5,
		1000 * 60 * 67,
		1000 * 60 * 60 * 12,
		1000 * 60 * 60 * 40,
		1000 * 60 * 60 * 999,
	}

	for _, milliseconds := range times {
		t.Run("", func(t *testing.T) {
			positive, err := ParseMilliseconds(milliseconds)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			negative, err := ParseMilliseconds(-milliseconds)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if positive.Days != -negative.Days ||
				positive.Hours != -negative.Hours ||
				positive.Minutes != -negative.Minutes ||
				positive.Seconds != -negative.Seconds ||
				positive.Milliseconds != -negative.Milliseconds {
				t.Errorf("negative values do not match positive values: got %v, want %v", negative, positive)
			}
		})
	}
}


func TestHandleNegativeFloatMilliseconds(t *testing.T) {
	times := []float64{
		0.0005,
		0.3,
	}

	for _, milliseconds := range times {
		t.Run("", func(t *testing.T) {
			positive, err := ParseMilliseconds(milliseconds)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			negative, err := ParseMilliseconds(-milliseconds)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if positive.Days != -negative.Days ||
				positive.Hours != -negative.Hours ||
				positive.Minutes != -negative.Minutes ||
				positive.Seconds != -negative.Seconds ||
				positive.Milliseconds != -negative.Milliseconds {
				t.Errorf("negative values do not match positive values: got %v, want %v", negative, positive)
			}
		})
	}
}
