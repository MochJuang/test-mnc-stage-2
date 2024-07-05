package main

import "fmt"

func calculateChange(totalPayment, amountPaid int) {
	var change int
	var fractions = make(map[int]int)
	if amountPaid < totalPayment {
		fmt.Println("False, Amount Paid is less than Total Payment.")
		return
	}

	change = amountPaid - totalPayment
	roundedChange := change - (change % 100)

	availableDenominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	remainingChange := roundedChange
	for _, fraction := range availableDenominations {
		if remainingChange >= fraction {
			count := remainingChange / fraction
			fractions[fraction] = count
			remainingChange -= fraction * count
		}
	}

	fmt.Printf("Change that must be given by the cashier: %d\n", change)
	fmt.Println("Money fraction:")
	for fraction, count := range fractions {
		unit := "coin"
		if fraction >= 1000 {
			unit = "sheet"
		}
		fmt.Printf("%d %s(s) %d\n", count, unit, fraction)
	}

}

func testCalculateChange() {
	totalPayment := 700649
	amountPaid := 750000
	calculateChange(totalPayment, amountPaid)

	fmt.Println()
	totalPayment = 65900
	amountPaid = 100000
	calculateChange(totalPayment, amountPaid)

	fmt.Println()
	totalPayment = 80000
	amountPaid = 50000
	calculateChange(totalPayment, amountPaid)
}
