package ratios

func DebtEquityRatio(debt, equity float32) (float32, error) {
	if equity == 0 {
		return 0.0, ErrDivideByZero
	}
	return debt / equity, nil
}
