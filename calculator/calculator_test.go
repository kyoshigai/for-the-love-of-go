package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type TestCase struct {
		a, b float64
		want float64
	}
	testCases := []TestCase{
		{a: 2, b: 1, want: 1},
		{a: 2, b: -2, want: 4},
		{a: 0, b: 0, want: 0},
		{a: 0, b: 3, want: -3},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b float64
		want float64
	}{
		{a: 6, b: 2, want: 12},
		{a: 3, b: -2, want: -6},
		{a: 0, b: 4, want: 0},
		{a: 0, b: -6, want: 0},
		{a: 0, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.3333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valied input, got %v", tc.a, tc.b, err)
		}
		if !closeEnough(tc.want, got, 0.0001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	var a, b float64 = 1, 0
	_, err := calculator.Divide(a, b)
	if err == nil {
		t.Errorf("Divide(%f, %f): want error for invalid input, got: %v", a, b, err)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
}
