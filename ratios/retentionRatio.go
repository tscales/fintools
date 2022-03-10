package ratios

func RetentionRatio(retainedEarnings float64, netIncome float64) (float64, error) {
	if netIncome == 0 {
		return 0.0, ErrDivideByZero
	}
	return retainedEarnings / netIncome, nil
}
