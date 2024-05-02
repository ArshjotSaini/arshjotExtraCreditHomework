//Arshjot Saini (00465708)

package main

import (
	"testing"
)

var tests = []struct {
	name  string
	state string
	input []Purchases
	want  float32 // Expected total price
}{
	{"Washington", "WA", bucket1, 135.82}, // Total price including tax for WA state
	{"Idaho", "ID", bucket1, 212.0},       // Total price including tax for ID state
	{"MT", "MT", bucket2, 135.0},          // Total price (no tax) for MT state
}

var bucket1 = []Purchases{
	{"Mango", "Fruit", 15},
	{"Orange", "Fruit", 20},
	{"MS Office", "Software", 100},
	{"Soap", "Detergent", 50},
}

var bucket2 = []Purchases{
	{"Potato", "Fruit", 15},
	{"Banana", "Fruit", 20},
	{"Windows OS", "Software", 100},
}

func TestCheckout(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalPrice, err := Checkout(tt.state, tt.input)

			// Check if error is expected
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Check if total price matches expected
			if totalPrice != tt.want {
				t.Errorf("got total price %.2f, want %.2f", totalPrice, tt.want)
			}
		})
	}
}
