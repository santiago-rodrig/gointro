package math

import "testing"

func TestAverage(t *testing.T) {
	input := make([]float64, 0)
	expected := 0.0
	output := Average(input)
	if expected != output {
		t.Errorf("Expected %f to equal %f", output, expected)
	}
}
