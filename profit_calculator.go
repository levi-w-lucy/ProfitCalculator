package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	retreiveValFromUser("Enter your revenue: ", &revenue)
	retreiveValFromUser("Enter your total expenses: ", &expenses)
	retreiveValFromUser("Enter your tax rate: ", &taxRate)

	earningsBeforeTax, profit, ebtProfitRatio := getFinancialData(revenue, expenses, taxRate)
	fmt.Println("Your earnings before tax is $" + fmt.Sprintf("%.2f", earningsBeforeTax))
	fmt.Println("Your earnings after tax is $" + fmt.Sprintf("%.2f", profit))
	fmt.Println("Your EBT to Profit Ratio is " + fmt.Sprintf("%.2f", ebtProfitRatio))
}

func retreiveValFromUser(message string, variablePtr *float64) {
	fmt.Print(message)
	fmt.Scan(variablePtr)
}

func getFinancialData(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ebtProfitRatio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate/100)
	ebtProfitRatio = earningsBeforeTax / profit
	return
}
