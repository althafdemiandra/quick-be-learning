package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name     string
		num1     float64
		num2     float64
		operator string
		want     float64
		wantErr  bool
	}{
		{"addition", 2, 3, "+", 5, false},
		{"subtraction", 5, 2, "-", 3, false},
		{"multiplication", 4, 2, "*", 8, false},
		{"division", 10, 2, "/", 5, false},
		{"division by zero", 10, 0, "/", 0, true},
		{"invalid operator", 2, 3, "^", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculate(test.num1, test.num2, test.operator)

			// Check error cases
			if test.wantErr {
				if err == nil {
					t.Errorf("calculate(%f, %f, %s) expected an error but got none", test.num1, test.num2, test.operator)
				}
				return
			}

			// Check non-error cases
			if err != nil {
				t.Errorf("calculate(%f, %f, %s) unexpected error: %v", test.num1, test.num2, test.operator, err)
				return
			}

			if result != test.want {
				t.Errorf("calculate(%f, %f, %s) = %f, want %f", test.num1, test.num2, test.operator, result, test.want)
			}
		})
	}
}
