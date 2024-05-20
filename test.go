package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	finalCost = bulkMessageCost
	if isPremiumUser {
		finalCost = bulkMessageCost - (bulkMessageCost * discountRate)
	}
	if accountBalance < finalCost {
		fmt.Println(insufficientFundMessage)
	} else {
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	}
	fmt.Println("Account balance:", accountBalance)
}
