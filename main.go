package main

import (
	"bufio"
	"fmt"
	"os"
)

///////////////////////////////////////////////////////////////////////////////
type debt struct {
	name   string
	amount float64
}

///////////////////////////////////////////////////////////////////////////////
func printMenu() {
	var numDebts int
	var debtAmt float64
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
		d := debt{debtName, debtAmt}
		debtSlice = append(debtSlice, d)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("---------------Debt Statistics---------------")
	for x := range debtSlice {
		fmt.Printf("%s", debtSlice[x].name)
		fmt.Printf("  $%.2f\n", debtSlice[x].amount)
	}
	fmt.Println("--------------Total Amount Owed--------------")
	fmt.Printf("  $%.2f\n", addDebts(debtSlice))
}

func addDebts(s []debt) float64 {
	var totalDebt float64

	for x := range s {
		totalDebt += s[x].amount
	}

	return totalDebt
}

///////////////////////////////////////////////////////////////////////////////
func main() {
	printMenu()

	os.Exit(1)
}
