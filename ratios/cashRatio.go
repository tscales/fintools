package ratios

//CashRatio is the company's cash assets divided by its current liabilities.
func CashRatio(cash, liabilities float32) (float32, error) {
	if liabilities == 0 {
		return 0.0, ErrDivideByZero
	}
	return cash / liabilities, nil
}
