package main

import "fmt"

///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
type debt struct {
	name   string
	amount float32
}

func (d debt) getName() {
	fmt.Println("Name of Debt: ", d.name)
}
func (d debt) getAmountOwed() {
	fmt.Println("Amount Owed: ", d.amount)
}

///////////////////////////////////////////////////////////////////////////////
func main() {
	d := debt{"School Loans", 5000}
	d.getName()
	d.getAmountOwed()
}