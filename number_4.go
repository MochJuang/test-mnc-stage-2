package main

import (
	"fmt"
	"math"
	"time"
)

// Function to determine if an employee can take personal leave
func canTakePersonalLeave(joinDate time.Time, plannedLeaveDate time.Time, duration int, totalPublicHolidays int) (bool, string) {
	const annualLeave = 14
	const probationDays = 180
	const maxConsecutiveLeave = 3

	if duration > maxConsecutiveLeave {
		return false, "Personal leave duration cannot exceed 3 consecutive days."
	}
	fmt.Printf("Join Date: %s\n", joinDate.Format("02-01-2006"))

	probationEndDate := joinDate.AddDate(0, 0, probationDays)
	fmt.Printf("Probation End Date: %s\n", probationEndDate.Format("02-01-2006"))
	if plannedLeaveDate.Before(probationEndDate) {
		return false, "New employees are not entitled to take personal leave during the first 180 days."
	}

	fmt.Printf("Duration: %d\n", duration)

	// ex : 14 - 7 (public holidays) = 7
	personalLeave := annualLeave - totalPublicHolidays

	// if still probation period
	endOfYear := time.Date(joinDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)

	// ex : 2024-12-31 - 2024-08-28 = ex: 125 days, if next year reset to 365 days (not new employee anymore)
	daysRemaining := endOfYear.Sub(probationEndDate).Hours() / 24

	if joinDate.Year() != plannedLeaveDate.Year() {
		//  have passed probation period
		endOfYear = time.Date(plannedLeaveDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
		daysRemaining = 365
	}

	// Calculate available personal leave quota (rounded down), 125 / 365 * 7 = 1.9178 = 1
	// or if next year, 365 / 365 * 7 = 7
	availableLeave := int(math.Floor(daysRemaining / 365 * float64(personalLeave)))

	// Check if the planned leave date is within the range after probation
	fmt.Printf("Planned Leave Date: %s\n", plannedLeaveDate.Format("02-01-2006"))
	fmt.Printf("Available Leave: %d\n", availableLeave)
	fmt.Printf("Days Remaining in Year: %.0f\n", daysRemaining)

	if duration > availableLeave {
		return false, fmt.Sprintf("The employee can only take %d days of personal leave in the first year.", availableLeave)
	}

	return true, "The employee is allowed to take personal leave."
}

func testCanTakePersonalLeave() {
	// Test cases
	tests := []struct {
		joinDate            time.Time
		plannedLeaveDate    time.Time
		duration            int
		totalPublicHolidays int
		expectedResult      bool
		expectedReason      string
	}{
		// Test Case 1
		{
			joinDate:            time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC),
			plannedLeaveDate:    time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
			duration:            1,
			totalPublicHolidays: 7,
			expectedResult:      false,
			expectedReason:      "New employees are not entitled to take personal leave during the first 180 days.",
		},
		{
			joinDate:            time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC),
			plannedLeaveDate:    time.Date(2021, 11, 5, 0, 0, 0, 0, time.UTC),
			duration:            3,
			totalPublicHolidays: 7,
			expectedResult:      false,
			expectedReason:      "The employee can only take 1 days of personal leave in the first year.",
		},
		{
			joinDate:            time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
			plannedLeaveDate:    time.Date(2021, 12, 18, 0, 0, 0, 0, time.UTC),
			duration:            1,
			totalPublicHolidays: 7,
			expectedResult:      true,
			expectedReason:      "The employee is allowed to take personal leave.",
		},
		{
			joinDate:            time.Date(2021, 1, 5, 0, 0, 0, 0, time.UTC),
			plannedLeaveDate:    time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC),
			duration:            3,
			totalPublicHolidays: 7,
			expectedResult:      true,
			expectedReason:      "The employee is allowed to take personal leave.",
		}}

	// Run test cases
	for i, test := range tests {
		result, reason := canTakePersonalLeave(test.joinDate, test.plannedLeaveDate, test.duration, test.totalPublicHolidays)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Expected Result: %v, Expected Reason: %s\n", test.expectedResult, test.expectedReason)
		fmt.Printf("Actual Result: %v, Actual Reason: %s\n\n", result, reason)
		fmt.Println()
	}

}
