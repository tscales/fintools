package ratios

// TimesInterestEarned calculates how well
// a company can cover its debt based on its income
func TimesInterestEarned(ebit, interest float32) (float32, error) {
	if interest == 0 {
		return 0.0, ErrDivideByZero
	}
	return ebit / interest, nil
}
