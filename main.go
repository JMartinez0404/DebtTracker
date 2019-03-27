package main

import (
	"bufio"
	"fmt"
	"os"
)

///////////////////////////////////////////////////////////////////////////////
type debt struct {
	name         string
	amount       float64
	interestRate float64
	life         int
}

///////////////////////////////////////////////////////////////////////////////
func printMenu() {
	var numDebts int
	var debtAmt float64
	var debtRate float64
	var debtLife int
	debtSlice := make([]debt, 0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--------WELCOME TO YOUR DEBT TRACKER---------")
	fmt.Println("---------------------------------------------")
	fmt.Println()
	fmt.Println("How many debts do you have? ")
	fmt.Scanln(&numDebts)
	for x := 1; x <= numDebts; x++ {
		fmt.Printf("What would you like to name debt #%d?\n", x)
		debtName, _ := reader.ReadString('\n')
		fmt.Printf("How much do you owe for debt #%d?\n", x)
		fmt.Scanln(&debtAmt)
		fmt.Printf("What is the interest rate of #%d?\n", x)
		fmt.Scanln(&debtRate)
		fmt.Printf("How old is debt #%d in months?\n", x)
		fmt.Scanln(&debtLife)
		d := debt{debtName, debtAmt, debtRate, debtLife}
		debtSlice = append(debtSlice, d)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("---------------Debt Statistics---------------")
	for x := range debtSlice {
		fmt.Printf("NAME: %s", debtSlice[x].name)
		fmt.Printf("  AMT: $%.2f\n", debtSlice[x].amount)
		fmt.Printf("  RATE: %.2f\n", debtSlice[x].interestRate)
		fmt.Printf("  LENGTH: %d\n", debtSlice[x].life)
	}
	fmt.Println("--------------Total Amount Owed--------------")
	fmt.Printf("  $%.2f\n", addDebts(debtSlice))
}

func addDebts(s []debt) float64 {
	var totalDebt float64

	for x := range s {
		totalDebt += s[x].amount + ((s[x].amount * (s[x].interestRate / 100)) * float64(s[x].life))
	}

	return totalDebt
}

///////////////////////////////////////////////////////////////////////////////
func main() {
	printMenu()

	os.Exit(1)
}
