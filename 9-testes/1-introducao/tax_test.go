package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f, got %f", item.expected, result)
		}
	}
}

// go test -bench=.
// go help test
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-2.5, -2.0, -1.0, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0.0 && result != 0.0 {
			t.Errorf("Received %v, expected 0", result)
		}
		if amount > 20000.0 && result != 20.0 {
			t.Errorf("Received %v, expected 20", result)
		}
	})
}
