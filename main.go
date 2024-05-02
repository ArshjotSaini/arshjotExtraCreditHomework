//Arshjot Saini (00465708)

package main

import (
	"fmt"
)

// Define a struct to represent purchased items
type Purchases struct {
	Name  string
	Type  string
	Price float32
}

// Function to calculate sales tax based on state and item type
func calculateSalesTax(state string, itemType string) float32 {
	// Check specific tax exemptions based on state and item type
	if state == "MT" {
		return 0.0
	} else if state == "ID" && itemType == "Software" {
		return 0.0
	} else if itemType == "WIC Eligible food" {
		return 0.0
	}

	// Apply default tax rates based on the state
	switch state {
	case "ID":
		return 0.06 // Idaho sales tax rate
	case "WA":
		return 0.0938 // Washington sales tax rate
	default:
		return 0.0 // Default: no sales tax
	}
}

// Function to calculate total price including tax for purchased items in a state
func Checkout(state string, purchasedItems []Purchases) (float32, error) {
	var totalTax float32
	var totalPrice float32

	// Check if the state is supported for checkout
	switch state {
	case "MT", "ID", "WA":
		// Calculate tax and total price for each item
		for _, item := range purchasedItems {
			unitTax := calculateSalesTax(state, item.Type)
			unitPrice := item.Price

			// Accumulate total tax and price
			totalTax += unitTax * unitPrice
			totalPrice += unitPrice
		}

		// Add total tax to the final price
		totalPrice += totalTax
		return totalPrice, nil
	default:
		// Return an error if the state is not supported
		return 0.0, fmt.Errorf("we don't do business in %v", state)
	}
}

func main() {
	// List of purchased items
	items := []Purchases{
		{"Apple", "WIC Eligible food", 1.0},
		{"Laptop", "Software", 1000.0},
		{"Chair", "Everything else", 50.0},
	}

	state := "ID"
	// Calculate total amount to charge for purchased items in a state
	total, err := Checkout(state, items)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Total Amount to Charge in %s: $%.2f\n", state, total)
	}
}
