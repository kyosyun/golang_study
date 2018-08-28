package test

import "testing"

func TestSum(t *testing.T) {
	for _, test := range []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
	} {
		result := sum(test.a, test.b)
		if result != test.expected {
			t.Errorf("expected: %d but result: %d", test.expected, result)
		}
	}
}
