Make sure both the main.go and main_test.go file are in the same folder

**calculateSalesTax(state string, itemType string) float32:** Determining the applicable sales tax rate depends on both the state and type of item, 
with certain exemptions applying based on specific states and item types.

**Checkout(state string, purchasedItems []Purchases) (float32, error):** Calculate the final cost, including taxes, for a selection of purchased items within a specified state. 
Utilize the calculateSalesTax function to apply the relevant taxes depending on the state and type of item.

**To run the tests**The repository contains unit tests (main_test.go) designed to verify the accuracy of the Checkout function across diverse scenarios, 
such as different states, item types, and tax rates. These tests confirm that the function operates as intended and manages various scenarios accurately.
