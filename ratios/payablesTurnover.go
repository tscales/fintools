package ratios

// PayablesTurnover states in days how long
// it takes for a company to pay their bills.
func PayablesTurnover(costOfGoodsSold, accountsPayable float32) (float32, error) {
	if accountsPayable == 0 {
		return 0.0, ErrDivideByZero
	}
	return costOfGoodsSold / accountsPayable, nil
}
