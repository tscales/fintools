package ratios

func DividenPayoutRatio(cashDividends float64, netIncome float64) (float64, error) {
	if netIncome == 0 {
		return 0.0, ErrDivideByZero
	}
	return cashDividends / netIncome, nil
}
