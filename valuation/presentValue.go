package valuation

// PresentValue takes a value cashFLow that represents future cash
// and represents what the value would be today if it was it were invested
// at with an interest rate rateOfReturn ( a percentage represented as a decimal)
// example: 25,000 dollars in cashFlow made from a rate of return of 3 percent
// 25,000 / 1 + .03 = 24271.84
func PresentValue(cashFlow, rateOfReturn float64) float64 {
	return cashFlow / (1 + rateOfReturn)
}
