package main

import (
	"fmt"
	"os"
)

const financialDataFile = "financialData.txt"

func main() {
	var revenue, expenses, taxRate float64

	revenue = retreiveValFromUser("Enter your revenue: ")
	expenses = retreiveValFromUser("Enter your total expenses: ")
	taxRate = retreiveValFromUser("Enter your tax rate: ")

	earningsBeforeTax, profit, ebtProfitRatio := getFinancialData(revenue, expenses, taxRate)

	ebtString := "Your earnings before tax is $" + fmt.Sprintf("%.2f", earningsBeforeTax)
	earningsString := "Your earnings after tax is $" + fmt.Sprintf("%.2f", profit)
	ebtRatio := "Your EBT to Profit Ratio is " + fmt.Sprintf("%.2f", ebtProfitRatio)

	fmt.Println(ebtString)
	fmt.Println(earningsString)
	fmt.Println(ebtRatio)

	writeFinancialData(fmt.Sprintf("%s\n%s\n%s", ebtString, earningsString, ebtRatio))
}

func retreiveValFromUser(message string) float64 {
	var inputVal float64

	fmt.Print(message)
	fmt.Scan(&inputVal)
	validateNumberInput(inputVal)

	return inputVal
}

func validateNumberInput(input float64) {
	if input <= 0 {
		panic("Input value was invalid. Value should be greater than 0")
	}
}

func getFinancialData(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ebtProfitRatio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate/100)
	ebtProfitRatio = earningsBeforeTax / profit
	return
}

func writeFinancialData(financialData string) {
	os.WriteFile(financialDataFile, []byte(financialData), 0644)
}
