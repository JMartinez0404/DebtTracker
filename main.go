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

func (d debt) getName() {
	fmt.Println("Name of Debt: ", d.name)
}
func (d debt) getAmountOwed() {
	fmt.Println("Amount Owed: ", d.amount)
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
	fmt.Println()
	fmt.Println(debtSlice)
}

func addDebts(nums ...float64) float64 {
	totalDebt := 0.0

	for _, x := range nums {
		totalDebt += x
	}

	return totalDebt
}

///////////////////////////////////////////////////////////////////////////////
func main() {
	printMenu()

	//nums := []float64{1200, 3000.75, 3500}
	//fmt.Println(addDebts(nums...))

	os.Exit(1)
}
