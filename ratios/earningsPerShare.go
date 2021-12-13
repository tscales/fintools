package ratios

func EarningsPerShare(netIncome, sharesOutstanding float32) (float32, error) {
	if sharesOutstanding == 0 {
		return 0.0, ErrDivideByZero
	}
	return netIncome / sharesOutstanding, nil
}
