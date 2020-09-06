package calculator_test

import (
	"calculator"
	"testing"
)

func TestArithmetics(t *testing.T) {
	t.Run("TestAdd", func(t *testing.T) {
		t.Parallel()
		var want float64 = 4
		got := calculator.Add(2, 2)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	})

	t.Run("TestSubtract", func(t *testing.T) {
		t.Parallel()
		var want float64 = 2
		got := calculator.Subtract(4, 2)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	})
}
