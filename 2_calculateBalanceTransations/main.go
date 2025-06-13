package main

import (
	"fmt"
)

func main() {
	transactions := []float64{}

	for {
		transaction := getTransaction()

		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	var userBalance float64 = calculateBalance(transactions)

	fmt.Printf("Your balance: %.2f \n", userBalance)
}

func getTransaction() (transaction float64) {
	fmt.Println("Enter your transaction or 0 for exit: ")
	fmt.Scan(&transaction)
	return
}

func calculateBalance(transactions []float64) (summa float64) {
	for _, value := range transactions {
		summa += value
	}

	return
}
