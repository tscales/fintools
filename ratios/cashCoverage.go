package ratios

// CashCoverageRatio is similar to the TimesInterestEarned ratio, but takes
// depreciation and amortization in to account
func CashCoverageRatio(ebit, depreciation, amortization, interest float32) (float32, error) {
	if interest == 0 {
		return 0.0, ErrDivideByZero
	}
	return (ebit + depreciation + amortization) / interest, nil
}
