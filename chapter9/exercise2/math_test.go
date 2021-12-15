package math

import "testing"

func TestAverage(t *testing.T) {
	input := make([]float64, 0)
	expected := 0.0
	output := Average(input)
	if expected != output {
		t.Errorf("expected %f to equal %f", output, expected)
	}
}

func TestMin(t *testing.T) {
	t.Run("a nonempty slice of numbers", func(t *testing.T) {
		input := []float64{-9, 33, -12, -39.5, 101.87}
		expected := -39.5
		output := Min(input)
		testNumbersEquality(t, "Min", input, output, expected)
	})

	t.Run("an empty slice of numbers", func(t *testing.T) {
		input := []float64{}
		expected := 0.0
		output := Min(input)
		testNumbersEquality(t, "Min", input, output, expected)
	})
}

func TestMax(t *testing.T) {
	t.Run("a nonempty slice of numbers", func(t *testing.T) {
		input := []float64{-9, 33, -12, -39.5, 101.87}
		expected := 101.87
		output := Max(input)
		testNumbersEquality(t, "Max", input, output, expected)
	})

	t.Run("an empty slice of numbers", func(t *testing.T) {
		input := []float64{}
		expected := 0.0
		output := Max(input)
		testNumbersEquality(t, "Max", input, output, expected)
	})
}

func testNumbersEquality(t testing.TB, funcName string, input interface{}, output, expected float64) {
	t.Helper()
	if expected != output {
		t.Errorf("%s(%v) = %f, expected %f", funcName, input, output, expected)
	}
}
